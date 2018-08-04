
-- 导出表 lzkp_bi.sys_region 结构
CREATE TABLE IF NOT EXISTS `sys_region` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '标识',
  `code` varchar(32) NOT NULL COMMENT '编号',
  `name` varchar(256) NOT NULL COMMENT '名称',
  `shorthand` varchar(64) NOT NULL COMMENT '简码',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父级id',
  `parent_ids` varchar(255) NOT NULL DEFAULT '0,' COMMENT '父级ids',
  `depth` tinyint(2) NOT NULL DEFAULT 1 COMMENT '深度',
  `is_end` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否未级',
  `is_disabled` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
  `disabled_date` datetime NOT NULL DEFAULT '2199-01-01' COMMENT '禁用日期',
  `api_url` varchar(255) NOT NULL DEFAULT '' COMMENT '接口地址',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `gmt_create` datetime(6) NOT NULL DEFAULT current_timestamp(6) COMMENT '创建日期',
  `gmt_modified` datetime(6) NOT NULL DEFAULT '1900-01-01 00:00:00.000000' COMMENT '更新日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`) USING BTREE,
  KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_parent_id` (`parent_id`) USING BTREE,
  KEY `idx_parent_ids` (`parent_ids`) USING BTREE,
  KEY `idx_is_disabled` (`is_disabled`) USING BTREE,
  KEY `idx_is_end` (`is_end`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='区划表';

-- 正在导出表  lzkp_bi.sys_region 的数据
INSERT IGNORE INTO `sys_region` (`id`, `code`, `name`, `shorthand`, `parent_id`, `parent_ids`, `depth`, `is_end`, `is_disabled`, `disabled_date`, `api_url`, `memo`, `gmt_create`, `gmt_modified`) VALUES
	(1, '3713', '临沂市', 'lys', 0, '0,', 1, 0, 0, NULL, '', '', '2018-05-23 10:15:08.610152', '1900-01-01 00:00:00.000000'),
	(2, '371300', '市直', 'sz', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-23 10:17:52.543405', '1900-01-01 00:00:00.000000'),
	(3, '371302', '兰山区', 'lsq', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-23 10:21:19.140449', '1900-01-01 00:00:00.000000'),
	(4, '371311', '罗庄区', 'lzp', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-26 09:19:53.384811', '1900-01-01 00:00:00.000000'),
	(5, '371312', '河东区', 'hdz', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:05:45.533203', '1900-01-01 00:00:00.000000'),
	(6, '371321', '沂南县', 'ynx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:06:16.989413', '1900-01-01 00:00:00.000000'),
	(7, '371322', '郯城县', 'tcx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:06:33.263325', '1900-01-01 00:00:00.000000'),
	(8, '371323', '沂水县', 'ysx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:06:49.411412', '1900-01-01 00:00:00.000000'),
	(9, '371324', '兰陵县', 'llx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:07:30.968390', '1900-01-01 00:00:00.000000'),
	(10, '371325', '费县', 'fx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:07:48.174359', '1900-01-01 00:00:00.000000'),
	(11, '371326', '平邑县', 'pyx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:08:10.783680', '1900-01-01 00:00:00.000000'),
	(12, '371327', '莒南县', 'jnx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:08:25.039009', '1900-01-01 00:00:00.000000'),
	(13, '371328', '蒙阴县', 'myx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:08:48.993049', '1900-01-01 00:00:00.000000'),
	(14, '371329', '临沭县', 'lsx', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:09:08.546779', '1900-01-01 00:00:00.000000'),
	(15, '371313', '高新区', 'gxq', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:12:21.561583', '1900-01-01 00:00:00.000000'),
	(16, '371314', '经开区', 'jkq', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:12:37.434487', '1900-01-01 00:00:00.000000'),
	(17, '371315', '临港区', 'lgq', 1, '0,1,', 2, 1, 0, NULL, '', '', '2018-05-31 17:12:53.343556', '1900-01-01 00:00:00.000000');

-- 
-- 审阅率汇总
-- v_date date:要处理的期间,v_succ：是否成功执行了过程
-- 
DROP PROCEDURE IF EXISTS pgene_review_rate;
delimiter $$
CREATE PROCEDURE pgene_review_rate(v_date date,out v_succ tinyint)
BEGIN
DECLARE v_firstday,v_lastday date;
DECLARE v_period mediumint UNSIGNED;
DECLARE v_start_time datetime;
DECLARE EXIT HANDLER FOR SQLEXCEPTION
BEGIN
  ROLLBACK;
  SET v_succ=0;
END;
SET v_succ=0;
SET v_start_time=NOW();
-- 计算这个月的起止日期
SELECT STR_TO_DATE(DATE_FORMAT(v_date,'%Y-%m-01'),'%Y-%m-%d') INTO v_firstday;
SELECT LAST_DAY(v_date) INTO v_lastday;
SET v_period=CONVERT(DATE_FORMAT(v_date,'%Y%m'),SIGNED);

TRUNCATE TABLE rep_review_rate;
INSERT INTO rep_review_rate (period,organ_id,depart_id,user_id,target_num,actual_num)
WITH cte_review AS
(SELECT DATE_FORMAT(wl.report_date,'%Y%m') period,wus.organ_id,wus.depart_id,wr.audit_user_id,
			 (CASE WHEN wl.id IS NULL OR DATEDIFF(wr.gmt_create,wl.report_date) > 12 THEN 0 ELSE 1 END) fact_number
	FROM work_review wr
			 INNER JOIN work_log wl ON wl.id = wr.work_log_id
			 INNER JOIN arc_organ ao ON ao.id = wr.organ_id
			 INNER JOIN sys_user su ON su.official_type_id>=1 AND su.id = wr.audit_user_id
			 INNER JOIN work_user_state wus ON wr.audit_user_id=wus.user_id AND CONVERT(wl.report_date,DATE)=wus.in_date 
  WHERE wl.report_date >= ao.entry_date
 UNION ALL
	SELECT DATE_FORMAT(wrq.gmt_create,'%Y%m') period,wus.organ_id,wus.depart_id,wrq.audit_user_id,        
			 0 fact_number
	FROM work_review_queue wrq
			 INNER JOIN arc_organ ao ON ao.id = wrq.organ_id
			 INNER JOIN sys_user su ON su.official_type_id>=1 AND su.id = wrq.audit_user_id
			 INNER JOIN work_user_state wus ON wrq.audit_user_id=wus.user_id AND CONVERT(wrq.gmt_create,DATE)=wus.in_date					 
  WHERE wrq.gmt_create >= ao.entry_date)
SELECT period,organ_id,depart_id,audit_user_id,count(*) AS target_number,SUM(fact_number) AS fact_number
FROM cte_review
GROUP BY period,organ_id,depart_id,audit_user_id
ORDER BY period,organ_id,depart_id,audit_user_id;
-- 写入执行日志
IF NOT EXISTS(SELECT id FROM sys_dispatch_log WHERE di_type=5 AND di_date=v_date) THEN
  INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo) VALUES(5,v_date,v_start_time,NOW(),'ok');
ELSE
  UPDATE sys_dispatch_log SET begin_time=v_start_time,end_time=NOW() WHERE di_type=5 AND di_date=v_date;
END IF;
SET v_succ=1;
END;$$

delimiter ;	

call pgene_review_rate ('2018-07-12',@succ);