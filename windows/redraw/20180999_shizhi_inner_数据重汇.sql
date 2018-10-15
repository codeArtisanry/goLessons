SET @b_date='2018-07-01';
SET @e_date='2018-10-06';
-- 删除已经汇总过去的队列
DELETE a,b FROM work_log_queue a
  INNER JOIN work_log_list_queue b ON a.id=b.queue_id
  WHERE a.log_date<=@e_date;
-- 根据日志表保存的所在科室更新用户状态表，解决生成不对的情况
UPDATE work_user_state a 
  INNER JOIN work_log b ON a.in_date=b.log_date AND a.user_id=b.user_id
  SET a.organ_id=b.organ_id,a.depart_id=b.depart_id,a.official_depart_id=b.official_depart_id
  WHERE a.in_date BETWEEN @b_date AND @e_date;
-- 应用更新单位假期的情况
UPDATE work_user_state a
  INNER JOIN arc_organ_holiday b ON a.organ_id=b.organ_id AND a.in_date BETWEEN b.begin_date AND b.end_date
  SET a.is_workday=0,a.is_must=0,a.dictionary_id=1 
  WHERE a.in_date BETWEEN @b_date AND @e_date AND a.is_workday=1;
-- 更新日志实际填报状态
UPDATE work_user_state a
  LEFT JOIN work_log b on a.in_date=b.log_date and a.user_id=b.user_id
  SET a.is_report=(CASE WHEN ifnull(b.state,0)>=1 THEN 1 ELSE 0 END)
  WHERE a.in_date BETWEEN @b_date AND @e_date;
 
-- 根据用户状态记录写入标准工作天数和时长
TRUNCATE TABLE rep_assess_user;
INSERT INTO rep_assess_user
  (period,organ_id,depart_id,user_id,standard_days,standard_hours)
SELECT CONVERT(date_format(in_date,'%Y%m'),UNSIGNED) AS period,
  organ_id,official_depart_id,user_id,SUM(is_workday) AS standard_days,
  SUM(is_workday)*7 AS standard_hours
  FROM work_user_state
  WHERE in_date BETWEEN @b_date AND @e_date AND is_assess=1 AND is_workday=1
  GROUP BY period,organ_id,official_depart_id,user_id
  ORDER BY period,organ_id,official_depart_id,user_id;
  
-- 根据日志填报情况汇总写入个人实际工作天数和时长
-- 根据每篇日志id进行首次分组汇总,归属科室取用户状态表
DROP TEMPORARY TABLE IF EXISTS temp_work_queue;
CREATE TEMPORARY TABLE temp_work_queue
SELECT a.id,MAX(a.user_id) AS user_id,MAX(a.log_date) AS log_date,
  MAX(c.organ_id) AS organ_id,MAX(c.official_depart_id) AS depart_id,
  MAX(b.work_log_id) AS work_log_id,0 AS is_revoke,
  1 AS actual_days,
  COUNT(*) AS actual_rows,
  SUM(b.cost_time) AS actual_hours,
  IF(MAX(c.is_workday)=1,1,0) AS workday_days,
  IF(MAX(c.is_workday)=1,COUNT(*),0) AS workday_rows,
  IF(MAX(c.is_workday)=1,SUM(b.cost_time),0) AS workday_hours,
  IF(SUM(CASE WHEN b.time_type_id=2 THEN 1 ELSE 0 END)>=1,1,0) AS overtime_days,
  SUM(CASE WHEN b.time_type_id=2 THEN 1 ELSE 0 END) AS overtime_rows,
  SUM(CASE WHEN b.time_type_id=2 THEN b.cost_time ELSE 0 END) AS overtime_hours,
  IF(SUM(CASE WHEN b.space_type_id=2 THEN 1 ELSE 0 END)>=1,1,0) AS travel_days,
  SUM(CASE WHEN b.space_type_id=2 THEN 1 ELSE 0 END) AS travel_rows,
  SUM(CASE WHEN b.space_type_id=2 THEN b.cost_time ELSE 0 END) AS travel_hours,
  (CASE WHEN MAX(b.duty_id)>0 THEN 1 ELSE 0 END) AS pt_duty,
  MAX(c.is_workday) AS is_workday,MAX(c.is_must) AS is_must,MAX(c.is_report) AS is_report,
  MAX(c.dictionary_id) AS dictionary_id,MAX(c.holiday_time) AS holiday_time 
  FROM work_log a
  INNER JOIN work_log_list b ON a.id=b.work_log_id AND (b.class_id!=48 or b.holiday_type_id IN (1,6,7,8,9))
  INNER JOIN work_user_state c ON a.log_date=c.in_date AND a.user_id=c.user_id AND c.is_assess=1
  WHERE a.log_date BETWEEN @b_date AND @e_date AND a.state>=1
  GROUP BY a.id;
-- 3.1汇总并更新用户实际发生情况
DROP TEMPORARY TABLE IF EXISTS temp_assess_user;
CREATE TEMPORARY TABLE temp_assess_user
SELECT CONVERT(DATE_FORMAT(log_date,'%Y%m'),SIGNED) AS period,
  organ_id,depart_id,user_id,
  SUM(actual_days) AS actual_days,SUM(actual_rows) AS actual_rows,SUM(actual_hours) AS actual_hours,
  SUM(workday_days) AS workday_days,SUM(workday_rows) AS workday_rows,SUM(workday_hours) AS workday_hours,
  SUM(overtime_days) AS overtime_days,SUM(overtime_rows) AS overtime_rows,SUM(overtime_hours) AS overtime_hours,
  SUM(travel_days) AS travel_days,SUM(travel_rows) AS travel_rows,SUM(travel_hours) AS travel_hours,
  SUM(pt_duty) AS pt_duty,0 AS up_flag   
  FROM temp_work_queue
  GROUP BY period,organ_id,depart_id,user_id
  ORDER BY period,organ_id,depart_id,user_id;
-- 更新用户指标表实际发生情况
UPDATE rep_assess_user a 
  INNER JOIN temp_assess_user b 
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id 
  SET b.up_flag=1,
  a.actual_days=a.actual_days+b.actual_days,a.actual_rows=a.actual_rows+b.actual_rows,a.actual_hours=a.actual_hours+b.actual_hours,
  a.workday_days=a.workday_days+b.workday_days,a.workday_rows=a.workday_rows+b.workday_rows,a.workday_hours=a.workday_hours+b.workday_hours,
  a.overtime_days=a.overtime_days+b.overtime_days,a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours,
  a.travel_days=a.travel_days+b.travel_days,a.travel_rows=a.travel_rows+b.travel_rows,a.travel_hours=a.travel_hours+b.travel_hours,
  a.pt_duty=a.pt_duty+b.pt_duty;
-- 插入新的用户记录
INSERT INTO rep_assess_user(period,organ_id,depart_id,user_id,standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty)
  SELECT period,organ_id,depart_id,user_id,0 AS standard_days,0 AS standard_hours,
    actual_days,actual_rows,actual_hours,
    workday_days,workday_rows,workday_hours,
    overtime_days,overtime_rows,overtime_hours,
    travel_days,travel_rows,travel_hours,
    pt_duty
  FROM temp_assess_user WHERE up_flag=0;

-- rep_assess_depart
-- 根据日志填报情况汇总写入科室实际工作天数和时长
TRUNCATE TABLE rep_assess_depart;
INSERT INTO rep_assess_depart(period,organ_id,depart_id,
  standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty)
  SELECT period,organ_id,depart_id,
    SUM(standard_days),SUM(standard_hours),
    SUM(actual_days),SUM(actual_rows),SUM(actual_hours),
    SUM(workday_days),SUM(workday_rows),SUM(workday_hours),
    SUM(overtime_days),SUM(overtime_rows),SUM(overtime_hours),
    SUM(travel_days),SUM(travel_rows),SUM(travel_hours),
    SUM(pt_duty)
  FROM rep_assess_user
  GROUP BY period,organ_id,depart_id
  ORDER BY period,organ_id,depart_id;
  
-- rep_assess_organ
-- 根据日志填报情况汇总写入单位实际工作天数和时长
TRUNCATE TABLE rep_assess_organ;
INSERT INTO rep_assess_organ(period,organ_id,
  standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty)
  SELECT period,organ_id,
    SUM(standard_days),SUM(standard_hours),
    SUM(actual_days),SUM(actual_rows),SUM(actual_hours),
    SUM(workday_days),SUM(workday_rows),SUM(workday_hours),
    SUM(overtime_days),SUM(overtime_rows),SUM(overtime_hours),
    SUM(travel_days),SUM(travel_rows),SUM(travel_hours),
    SUM(pt_duty)
  FROM rep_assess_depart
  GROUP BY period,organ_id
  ORDER BY period,organ_id;

-- 履职执行情况数据汇总准备,抽调、主要领导不考核统计,归属科室取日志表
DROP TEMPORARY TABLE IF EXISTS temp_query1;
CREATE TEMPORARY TABLE temp_query1
SELECT b.work_log_id,a.log_date,a.organ_id,a.official_depart_id,a.user_id,
  b.cost_time,b.time_type_id,b.class_id,b.duty_id
  FROM work_log a  
  INNER JOIN work_log_list b ON a.state>=1 AND a.id=b.work_log_id and (b.class_id!=48 or b.holiday_type_id IN (1,6,7,8,9))
  INNER JOIN work_user_state c ON a.log_date=c.in_date AND a.user_id=c.user_id AND c.is_assess=1
  WHERE a.log_date BETWEEN @b_date AND @e_date;
  
-- rep_perform_duty 岗位职责
TRUNCATE TABLE rep_perform_duty;
INSERT INTO rep_perform_duty(period,organ_id,depart_id,user_id,duty_id,
work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT CONVERT(DATE_FORMAT(log_date,'%Y%m'),UNSIGNED) AS period,
  organ_id,official_depart_id,user_id,duty_id,
  COUNT(*) AS ts,SUM(cost_time) AS sc,
  SUM(CASE WHEN time_type_id=2 THEN 1 ELSE 0 END) AS jbts,
  SUM(CASE WHEN time_type_id=2 THEN cost_time ELSE 0 END) AS jbsc 
  FROM temp_query1
  WHERE duty_id>=1
  GROUP BY period,organ_id,official_depart_id,user_id,duty_id
  ORDER BY period,organ_id,official_depart_id,user_id,duty_id;
 
-- rep_perform_class 工作分类职责分类
TRUNCATE TABLE rep_perform_class;
INSERT INTO rep_perform_class(period,organ_id,depart_id,user_id,
level2_id,work_rows,work_hours,overtime_rows,overtime_hours,level1_id)
SELECT a1.*,b1.level1_id as level1_id 
FROM  
  (SELECT CONVERT(DATE_FORMAT(log_date,'%Y%m'),UNSIGNED) AS period,
  organ_id,official_depart_id,user_id,class_id,    
  COUNT(*) AS ts,SUM(cost_time) AS sc,
  SUM(CASE WHEN time_type_id=2 THEN 1 ELSE 0 END) AS jbts,
  SUM(CASE WHEN time_type_id=2 THEN cost_time ELSE 0 END) AS jbsc 
  FROM temp_query1
  GROUP BY period,organ_id,official_depart_id,user_id,class_id) a1
  INNER JOIN 
    (SELECT id,IF(depth=1,id,CONVERT(SUBSTRING(SUBSTRING_INDEX(parent_ids,',',2),3),UNSIGNED)) AS level1_id 
     FROM arc_work_class) b1 ON a1.class_id=b1.id
ORDER BY a1.period,a1.organ_id,a1.official_depart_id,a1.user_id,a1.class_id;

-- rep_perform_class_pt 工作分类执行人次
DROP TEMPORARY TABLE IF EXISTS temp_pt;
CREATE TEMPORARY TABLE temp_pt 
  SELECT a.work_log_id,b.level1_id as class_id
  FROM temp_query1 a  
  INNER JOIN 
    (SELECT id,IF(depth=1,id,CONVERT(SUBSTRING(SUBSTRING_INDEX(parent_ids,',',2),3),UNSIGNED)) AS level1_id 
     FROM arc_work_class) b ON a.class_id=b.id
	GROUP BY a.work_log_id,level1_id;

TRUNCATE TABLE rep_perform_class_pt;
INSERT INTO rep_perform_class_pt(period,organ_id,depart_id,user_id,class_id,pt_num)
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,
  a.organ_id,a.official_depart_id,a.user_id,b.class_id,COUNT(*) AS pt_num    
  FROM work_log a 
  INNER JOIN temp_pt b ON a.id=b.work_log_id
  GROUP BY period,organ_id,official_depart_id,user_id,class_id
  ORDER BY period,organ_id,official_depart_id,user_id,class_id;
  
-- rep_perform_class_ptl 工作分类执行人次_二级
DROP TEMPORARY TABLE IF EXISTS temp_pt;
CREATE TEMPORARY TABLE temp_pt 
  SELECT a.work_log_id,a.class_id
  FROM temp_query1 a  
  INNER JOIN arc_work_class b ON a.class_id=b.id
	GROUP BY a.work_log_id,a.class_id;

TRUNCATE TABLE rep_perform_class_ptl;
INSERT INTO rep_perform_class_ptl(period,organ_id,depart_id,user_id,class_id,pt_num)
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,
  a.organ_id,a.official_depart_id,a.user_id,b.class_id,COUNT(*) AS pt_num    
  FROM work_log a 
  INNER JOIN temp_pt b ON a.id=b.work_log_id
  GROUP BY period,organ_id,official_depart_id,user_id,class_id
  ORDER BY period,organ_id,official_depart_id,user_id,class_id;

-- 填报率汇总,归属科室取用户状态表
TRUNCATE TABLE rep_rate_user;
INSERT INTO rep_rate_user(period,organ_id,official_depart_id,user_id,target_number,actual_number)
SELECT CONVERT(DATE_FORMAT(in_date,'%Y%m'),UNSIGNED) AS period,
 organ_id,official_depart_id,user_id,SUM(is_must) AS target_number,SUM(is_report) AS actual_number
FROM work_user_state
WHERE in_date BETWEEN @b_date AND @e_date AND is_must=1
GROUP BY period,organ_id,official_depart_id,user_id
ORDER BY period,organ_id,official_depart_id,user_id;

TRUNCATE TABLE rep_rate_depart;
INSERT INTO rep_rate_depart(period,organ_id,official_depart_id,target_number,actual_number)
SELECT period,organ_id,official_depart_id,SUM(target_number) AS target_number,SUM(actual_number) AS actual_number
FROM rep_rate_user
GROUP BY period,organ_id,official_depart_id
ORDER BY period,organ_id,official_depart_id;

TRUNCATE TABLE rep_rate_organ;
INSERT INTO rep_rate_organ(period,organ_id,target_number,actual_number)
SELECT period,organ_id,SUM(target_number) AS target_number,SUM(actual_number) AS actual_number
FROM rep_rate_depart
GROUP BY period,organ_id
ORDER BY period,organ_id;

-- 审阅率重汇
SET @e_date=ADDDATE(@e_date,1);
TRUNCATE TABLE rep_review_rate;
INSERT INTO rep_review_rate (period,organ_id,depart_id,user_id,target_num,actual_num)
WITH cte_review AS
(SELECT DATE_FORMAT(wl.report_date,'%Y%m') period,wr.organ_id,ad.official_depart_id AS depart_id,wr.audit_user_id,
			 (CASE WHEN wl.id IS NULL OR DATEDIFF(wr.gmt_create,wl.report_date) > 12 THEN 0 ELSE 1 END) fact_number
	FROM work_review wr
			 INNER JOIN work_log wl ON wl.id = wr.work_log_id
			 INNER JOIN arc_organ ao ON ao.id = wr.organ_id
			 INNER JOIN sys_user su ON su.official_type_id>=1 AND su.id = wr.audit_user_id
			 INNER JOIN arc_department ad ON wr.depart_id=ad.id
			 INNER JOIN work_user_state k ON wr.log_date=k.in_date AND wr.user_id=k.user_id AND k.is_workday=1 
  WHERE wl.report_date>=@b_date AND wl.report_date<@e_date
 UNION ALL
	SELECT DATE_FORMAT(wrq.gmt_create,'%Y%m') period,wrq.organ_id,ad.official_depart_id AS depart_id,wrq.audit_user_id,        
			 0 fact_number
	FROM work_review_queue wrq
			 INNER JOIN arc_organ ao ON ao.id = wrq.organ_id
			 INNER JOIN sys_user su ON su.official_type_id>=1 AND su.id = wrq.audit_user_id
			 INNER JOIN arc_department ad ON wrq.depart_id=ad.id
			 INNER JOIN work_user_state k ON wrq.log_date=k.in_date AND wrq.user_id=k.user_id AND k.is_workday=1 
  WHERE wrq.gmt_create>=@b_date AND wrq.gmt_create<@e_date)
SELECT period,organ_id,depart_id,audit_user_id,count(*) AS target_number,SUM(fact_number) AS fact_number
FROM cte_review
GROUP BY period,organ_id,depart_id,audit_user_id
ORDER BY period,organ_id,depart_id,audit_user_id;

-- 对16张汇总表进行201806以前的历史数据备份并删除
delete from rep_assess_depart where period<=201806;
delete from rep_assess_organ where period<=201806;
delete from rep_assess_user where period<=201806;
delete from rep_job_rate where period<=201806;	
delete from rep_keep_official where period<=201806;
delete from rep_perform_approval where period<=201806;
delete from rep_perform_class where period<=201806;
delete from rep_perform_class_pt where period<=201806;		
delete from rep_perform_class_ptl where period<=201806;
delete from rep_perform_duty where period<=201806;	
delete from rep_perform_supervision where period<=201806;	
delete from rep_perform_yearly_plan where period<=201806;
delete from rep_rate_depart where period<=201806;
delete from rep_rate_organ where period<=201806;
delete from rep_rate_user where period<=201806;	
delete from rep_review_rate where period<=201806;
