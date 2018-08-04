
DROP TABLE IF EXISTS `临沂市荣军医院人员信息表`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `临沂市荣军医院人员信息表` (
  `序号` varchar(255) DEFAULT NULL,
  `姓名` varchar(255) DEFAULT NULL,
  `岗位(身份)标识` varchar(255) DEFAULT NULL,
  `编制所在科室` varchar(255) DEFAULT NULL,
  `调岗情况` varchar(255) DEFAULT NULL,
  `修改意见` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `临沂市荣军医院人员信息表`
--

LOCK TABLES `临沂市荣军医院人员信息表` WRITE;
/*!40000 ALTER TABLE `临沂市荣军医院人员信息表` DISABLE KEYS */;
INSERT INTO `临沂市荣军医院人员信息表` VALUES ('1','吕锋','公务员','吕院长',NULL,'事业'),('2','公茂忠','事业','公书记',NULL,NULL),('3','荣梅','公务员','办公室','借调','事业'),('4','高树乾','公务员','办公室','借调','事业'),('5','鲁守明','事业','办公室',NULL,NULL),('6','刘学生','事业','办公室',NULL,NULL),('7','吴英萍','公务员','办公室','借调','事业'),('8','刘明磊','事业','办公室',NULL,'编外'),('9','张立征','事业','办公室',NULL,'编外'),('10','刘元帅','事业','办公室',NULL,NULL),('11','王瑶','事业','办公室',NULL,NULL),('12','李继韬','公务员','办公室','借调','事业'),('13','黄大伟','事业','财务科',NULL,NULL),('14','张玉霞','事业','财务科',NULL,NULL),('15','房书妮','事业','财务科',NULL,'编外'),('16','葛立芝','事业','财务科',NULL,'编外'),('17','徐颖','事业','财务科','混岗','编外'),('18','杨云莲','事业','财务科',NULL,'编外'),('19','孙华飞','事业','财务科',NULL,'编外'),('20','王洛玲','事业','财务科',NULL,'编外'),('21','段效林','事业','财务科',NULL,NULL),('22','高华','事业','高院长',NULL,NULL),('23','王继委','事业','医教科',NULL,NULL),('24','范玉霞','事业','医保办',NULL,NULL),('25','魏伟','事业','医保办',NULL,'编外'),('26','葛安国','事业','防治科',NULL,NULL),('27','高林平','事业','防治科',NULL,NULL),('28','朱海林','事业','防治科',NULL,'编外'),('29','侯安涛','事业','药房',NULL,NULL),('30','刘忠菊','事业','药房',NULL,NULL),('31','梁艳','事业','药房',NULL,'编外'),('32','张庆','事业','药房',NULL,NULL),('33','李春','事业','药房',NULL,NULL),('34','王东昕','事业','老年科',NULL,NULL),('35','张辉','事业','老年科',NULL,'编外'),('36','李恋','事业','老年科',NULL,'编外'),('37','孙运堂','事业','老年科',NULL,'编外'),('38','张裕裕','事业','老年科',NULL,'编外'),('39','朱燕','事业','老年科',NULL,NULL),('40','张琨','事业','老年科',NULL,NULL),('41','吴晓兰','事业','老年科',NULL,'编外'),('42','宋刚','事业','老年科',NULL,NULL),('43','张栋','事业','老年科',NULL,NULL),('44','刘建','事业','老年科','混岗',NULL),('45','吕文健','事业','康复科',NULL,'编外'),('46','左津如','事业','康复科',NULL,'编外'),('47','宋维本','事业','康复科',NULL,'编外'),('48','逯芳','事业','康复科',NULL,'编外'),('49','阚京梅','事业','康复科',NULL,NULL),('50','张蕾','事业','康复科',NULL,'编外'),('51','高强','事业','康复科',NULL,'编外'),('52','梁立升','事业','康复科',NULL,'编外'),('53','韩士杰','事业','康复科',NULL,'编外'),('54','王静','事业','康复科',NULL,'编外'),('55','王玉琨','事业','一区（医疗）',NULL,'编外'),('56','高丹萍','事业','一区（医疗）',NULL,'编外'),('57','孙运亮','事业','一区（医疗）',NULL,NULL),('58','付成杨','事业','一区（医疗）',NULL,NULL),('59','王媛媛','事业','一区（医疗）','混岗',NULL),('60','杨慧','事业','二区（医疗）',NULL,'编外'),('61','刁俊荣','事业','二区（医疗）',NULL,NULL),('62','王关龙','事业','二区（医疗）',NULL,NULL),('63','李洪全','事业','三区（医疗）',NULL,NULL),('64','崔文敏','事业','三区（医疗）',NULL,NULL),('65','韩帮海','事业','三区（医疗）',NULL,NULL),('66','魏效峰','事业','三区（医疗）',NULL,NULL),('67','赵连红','事业','三区（医疗）',NULL,NULL),('68','李振','事业','六区（医疗）',NULL,NULL),('69','朱现玉','事业','六区（医疗）',NULL,NULL),('70','刘刚','事业','六区（医疗）',NULL,NULL),('71','孙文正','事业','六区（医疗）',NULL,NULL),('72','李号号','事业','七区（医疗）',NULL,NULL),('73','曾敏','事业','七区（医疗）',NULL,'编外'),('74','黄萌娜','事业','七区（医疗）',NULL,'编外'),('75','张海东','事业','七区（医疗）',NULL,NULL),('76','李婷','事业','七区（医疗）',NULL,NULL),('77','张鹏兰','事业','心理康复',NULL,'编外'),('78','张艳','事业','心理康复',NULL,'编外'),('79','张露','事业','心理康复',NULL,NULL),('80','徐伟','事业','心理康复',NULL,'编外'),('81','刘晓利','事业','心理康复',NULL,NULL),('82','刘莹','事业','心理康复',NULL,'编外'),('83','刘思含','事业','心理康复',NULL,'编外'),('84','宋晓丽','事业','心理康复',NULL,NULL),('85','王炳霖','事业','心理康复',NULL,'编外'),('86','文欣','事业','病案室',NULL,'编外'),('87','信先凤','事业','病案室',NULL,NULL),('88','朱炜','事业','网络部',NULL,NULL),('89','郑永梅','事业','护理部',NULL,NULL),('90','陈少虎','事业','控感办',NULL,NULL),('91','赵国强','事业','摆药站',NULL,NULL),('92','史静','事业','摆药站',NULL,NULL),('93','田刚','事业','摆药站',NULL,NULL),('94','张清平','事业','摆药站','混岗','编外'),('95','刘欣','事业','摆药站','混岗','编外'),('96','柳娟','事业','一区',NULL,'编外'),('97','余飞','事业','一区',NULL,NULL),('98','杨晓云','事业','一区',NULL,NULL),('99','袁堂燕','事业','一区',NULL,'编外'),('100','吕霜','事业','一区',NULL,'编外'),('101','杨彩霞','事业','一区',NULL,'编外'),('102','许传才','事业','一区',NULL,'编外'),('103','薛宪伟','事业','一区',NULL,'编外'),('104','梁伟','事业','一区',NULL,'编外'),('105','张发达','事业','一区',NULL,'编外'),('106','陈明慧','事业','一区',NULL,NULL),('107','姜杨','事业','一区','混岗','编外'),('108','宋克恩','事业','二区',NULL,'编外'),('109','聂宗成','事业','二区','混岗',NULL),('110','李善晓','事业','二区',NULL,'编外'),('111','李晓云','事业','二区',NULL,'编外'),('112','顾宗艳','事业','二区',NULL,'编外'),('113','张华','事业','二区',NULL,'编外'),('114','张婷婷','事业','二区',NULL,'编外'),('115','刘利媛','事业','二区',NULL,'编外'),('116','王雪','事业','二区',NULL,'编外'),('117','刘春丽','事业','二区',NULL,'编外'),('118','姜良利','事业','二区',NULL,'编外'),('119','杨四妹','事业','二区','混岗','编外'),('120','孙红艳','事业','三区','混岗','编外'),('121','张发扬','事业','三区','混岗','编外'),('122','沈中伟','事业','三区',NULL,'编外'),('123','付传欣','事业','三区',NULL,'编外'),('124','王君伟','事业','三区',NULL,'编外'),('125','孟庆凤','事业','三区',NULL,'编外'),('126','张永伟','事业','三区',NULL,'编外'),('127','徐艺文','事业','三区',NULL,'编外'),('128','葛晓婷','事业','三区',NULL,'编外'),('129','徐桂珍','事业','三区',NULL,'编外'),('130','张宏艳','事业','三区',NULL,'编外'),('131','张永义','事业','三区',NULL,'编外'),('132','杨璐璐','事业','三区',NULL,'编外'),('133','黄艳','事业','三区','混岗','编外'),('134','杜桂凤','事业','六区',NULL,''),('135','董凡萌','事业','六区',NULL,'编外'),('136','鲁帆','事业','六区',NULL,'编外'),('137','刘莹莹','事业','六区',NULL,'编外'),('138','曹广胜','事业','六区',NULL,'编外'),('139','王善笑','事业','六区',NULL,'编外'),('140','朱国栋','事业','六区',NULL,'编外'),('141','刘俊辰','事业','六区',NULL,'编外'),('142','赵玉艳','事业','六区',NULL,'编外'),('143','陈晶晶','事业','六区',NULL,'编外'),('144','徐艳','事业','六区',NULL,'编外'),('145','李涛','事业','六区',NULL,'编外'),('146','王树艳','事业','六区',NULL,'编外'),('147','王阳阳','事业','六区',NULL,'编外'),('148','史杜娟','事业','六区',NULL,'编外'),('149','宋百丽','事业','六区','混岗','编外'),('150','周岩娟','事业','七区','混岗','编外'),('151','王娜','事业','七区','混岗','编外'),('152','徐丛梅','事业','七区',NULL,NULL),('153','朱蒙蒙','事业','七区',NULL,'编外'),('154','于晓宁','事业','七区',NULL,'编外'),('155','张冠军','事业','七区',NULL,'编外'),('156','公培培','事业','七区',NULL,'编外'),('157','李红满','事业','七区',NULL,'编外'),('158','管媛媛','事业','七区',NULL,'编外'),('159','林小雁','事业','七区',NULL,'编外'),('160','侯典坤','事业','七区',NULL,'编外'),('161','梁纪伟','事业','老年一科','混岗','编外'),('162','来守红','事业','老年一科',NULL,'编外'),('163','张媛媛','事业','老年一科',NULL,NULL),('164','张丽丽','事业','老年一科',NULL,'编外'),('165','张虹','事业','老年一科',NULL,'编外'),('166','王辉','事业','老年一科',NULL,'编外'),('167','林善善','事业','老年一科',NULL,'编外'),('168','梁培存','事业','老年一科',NULL,'编外'),('169','钮薛花','事业','老年一科',NULL,'编外'),('170','吕军珍','事业','老年一科',NULL,'编外'),('171','朱元庆','事业','老年一科',NULL,'编外'),('172','赵帅帅','事业','老年一科',NULL,'编外'),('173','李娟','事业','老年一科',NULL,NULL),('174','景敏','事业','老年一科',NULL,'编外'),('175','王建雷','事业','老年一科',NULL,'编外'),('176','刘建美','事业','老年一科',NULL,'编外'),('177','朱琳','事业','老年一科',NULL,'编外'),('178','王玉珠','事业','老年一科',NULL,'编外'),('179','薛祥荣','事业','老年一科',NULL,'编外'),('180','孙崇艳','事业','老年一科',NULL,'编外'),('181','倪瑞清','事业','老年一科','混岗','编外'),('182','张洪英','事业','老年一科',NULL,'编外'),('183','韩笑','事业','老年一科',NULL,'编外'),('184','王萌萌','事业','老年一科','混岗','编外'),('185','满纪荣','事业','老年二科',NULL,NULL),('186','郝淑敏','事业','老年二科',NULL,'编外'),('187','徐莹莹','事业','老年二科',NULL,'编外'),('188','王丽荣','事业','老年二科',NULL,'编外'),('189','宋娜','事业','老年二科',NULL,'编外'),('190','柏伟','事业','老年二科',NULL,'编外'),('191','刘宁宁','事业','老年二科',NULL,'编外'),('192','李静','事业','老年二科',NULL,'编外'),('193','侯伟利','事业','老年二科',NULL,'编外'),('194','刘晓杰','事业','老年二科',NULL,'编外'),('195','雷婷婷','事业','老年二科',NULL,'编外'),('196','周同','事业','老年二科',NULL,'编外'),('197','李佳佳','事业','老年二科',NULL,'编外'),('198','王敏梅','事业','老年二科','混岗','编外'),('199','杜敏','事业','老年二科','混岗','编外'),('200','方圆','事业','老年二科','混岗','编外'),('201','闫桂慧','事业','老年三科',NULL,'编外'),('202','于春枫','事业','老年三科',NULL,NULL),('203','张丽','事业','老年三科','混岗','编外'),('204','张美艳','事业','老年三科',NULL,'编外'),('205','刘春苗','事业','老年三科',NULL,'编外'),('206','张夫凤','事业','老年三科',NULL,'编外'),('207','王慧','事业','老年三科',NULL,NULL),('208','刘慧芳','事业','老年三科',NULL,'编外'),('209','曹飞','事业','老年三科',NULL,'编外'),('210','刘霞','事业','老年三科',NULL,'编外'),('211','鲁娜岚','事业','老年四科','混岗',NULL),('212','高金秋','事业','老年四科','混岗','编外'),('213','吴媛媛','事业','老年四科','混岗','编外'),('214','董学丽','事业','老年四科',NULL,NULL),('215','马丽','事业','老年四科',NULL,'编外'),('216','柏建英','事业','老年四科',NULL,'编外'),('217','胡金玲','事业','老年四科',NULL,'编外'),('218','毛传芹','事业','老年四科',NULL,'编外'),('219','张培才','事业','老年四科',NULL,'编外'),('220','李玲玲','事业','老年四科',NULL,'编外'),('221','徐国强','事业','老年四科',NULL,'编外'),('222','徐仰慧','事业','老年四科',NULL,'编外'),('223','陈继红','事业','老年四科',NULL,'编外'),('224','张蕊','事业','老年四科',NULL,'编外'),('225','孙丽','事业','老年四科',NULL,'编外'),('226','刘君','事业','老年四科',NULL,'编外'),('227','杜增才','事业','老年四科',NULL,'编外'),('228','王伟','事业','老年四科',NULL,'编外'),('229','刘茼芝','事业','老年四科',NULL,'编外'),('230','赵清华','事业','老年四科',NULL,'编外'),('231','文红彬','事业','老年四科',NULL,'编外'),('232','张珊珊','事业','老年四科',NULL,'编外'),('233','庄云云','事业','老年五科','混岗','编外'),('234','张新华','事业','老年五科','混岗','编外'),('235','汤有源','事业','老年五科',NULL,'编外'),('236','朱银超','事业','老年五科',NULL,'编外'),('237','王静','事业','老年五科','混岗','编外'),('238','吕玉梅','事业','老年五科',NULL,NULL),('239','陈晓磊','事业','老年五科',NULL,'编外'),('240','徐敏','事业','老年五科',NULL,'编外'),('241','周云霞','事业','老年五科',NULL,'编外'),('242','袁蒙蒙','事业','老年五科',NULL,'编外'),('243','朱秋秋','事业','老年五科',NULL,'编外'),('244','李盈盈','事业','老年五科',NULL,'编外'),('245','冉德敏','事业','老年五科',NULL,'编外'),('246','李文娜','事业','老年五科',NULL,'编外'),('247','陈璐','事业','老年五科',NULL,'编外'),('248','李云','事业','老年五科',NULL,'编外'),('249','吴艳芹','事业','老年五科',NULL,'编外'),('250','薛园园','事业','老年五科',NULL,'编外'),('251','朱孔香','事业','门诊部',NULL,NULL),('252','孙玉梅','事业','心理测验',NULL,NULL),('253','公萃栾','事业','心理测验',NULL,'编外'),('254','王静','事业','心理测验',NULL,NULL),('255','王乐娜','事业','心理测验',NULL,NULL),('256','李超','事业','影像科',NULL,NULL),('257','文庆松','事业','影像科',NULL,NULL),('258','程仕亮','事业','影像科',NULL,'编外'),('259','晁家伟','事业','影像科',NULL,NULL),('260','李先华','事业','检验科',NULL,NULL),('261','沈慧娟','事业','检验科',NULL,'编外'),('262','任洪娟','事业','检验科',NULL,NULL),('263','张磊','事业','电生理',NULL,NULL),('264','李作珍','事业','电生理',NULL,NULL),('265','王雪峰','事业','口腔科',NULL,'编外'),('266','刘春霞','事业','中医科',NULL,'编外'),('267','杨颖','事业','杨院长',NULL,NULL),('268','王萍','事业','质控科',NULL,NULL),('269','唐震','事业','总务科',NULL,NULL),('270','崔强','事业','总务科',NULL,NULL),('271','咸丰年','事业','总务科',NULL,NULL),('272','盛左超','事业','总务科',NULL,NULL),('273','李明景','事业','总务科',NULL,NULL),('274','王德平','事业','总务科','混岗',NULL),('275','周西良','事业','膳食科',NULL,NULL),('276','李秀泉','事业','李院长',NULL,NULL),('277','王军艳','事业','信息科',NULL,NULL),('278','周海燕','事业','信息科',NULL,NULL);
/*!40000 ALTER TABLE `临沂市荣军医院人员信息表` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'shizhi'
--

--
-- Dumping routines for database 'shizhi'
--
/*!50003 DROP PROCEDURE IF EXISTS `pgene_job_rate` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE PROCEDURE `pgene_job_rate`(v_date date,out v_succ tinyint)
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
-- 创建单位人员临时表条件：实体、日报、加入和撤销日期统计期间有效单位，在编且有编制单位、考核日期在统计期间的人员。
DROP TEMPORARY TABLE IF EXISTS temp_user_info;
CREATE TEMPORARY TABLE temp_user_info
SELECT a.id,a.official_organ_id,CONVERT(0,INT) AS is_report,
  CONVERT(0,INT) AS is_sendout,CONVERT(0,INT) AS is_unjob 
  FROM sys_user a
  INNER JOIN arc_organ b ON a.official_organ_id=b.id AND b.organ_type>0 AND b.report_type IN (1,3) 
  AND b.entry_date <=v_lastday AND IF(b.revoke_date='1900-01-01','2099-12-31',b.revoke_date)>=v_firstday
  WHERE a.official_type_id>=1 AND a.official_organ_id>=1 AND a.entry_date<=v_lastday AND a.retire_date>=v_firstday;
-- 更新用户当天的调岗状态,标记调出和不在岗人员
UPDATE temp_user_info a 
  INNER JOIN work_user_state b ON a.id=b.user_id AND b.in_date=v_date
  LEFT JOIN arc_transfer_type c ON b.transfer_type_id=c.id
  SET a.is_sendout=(CASE WHEN IFNULL(c.dictionary_id,0) IN (2,3) THEN 1 ELSE 0 END),
      a.is_unjob=(CASE WHEN IFNULL(c.dictionary_id,0)=5 THEN 1 ELSE 0 END);
-- 更新日志填报情况,判断逻辑：需要填报日志(不包括跨单位调整、外派、内退和长期病假),没有填报的视为空岗
UPDATE temp_user_info a1 
  INNER JOIN (
  SELECT a.user_id,
    (CASE WHEN SUM(a.is_must)>=1 AND SUM(a.is_report)<=0 THEN 0 ELSE 1 END) AS is_report  
    FROM work_user_state a
    INNER JOIN temp_user_info b ON a.user_id=b.id AND a.organ_id=b.official_organ_id AND b.is_sendout=0 AND b.is_unjob=0
    WHERE a.in_date BETWEEN v_firstday AND v_lastday
    GROUP BY a.user_id) b1 ON a1.id=b1.user_id
  SET a1.is_report=b1.is_report
  WHERE a1.is_sendout=0 AND a1.is_unjob=0;
-- 汇总不在岗和填报情况
DROP TEMPORARY TABLE IF EXISTS temp_job_info;
CREATE TEMPORARY TABLE temp_job_info
SELECT v_period AS period,official_organ_id AS organ_id,
  SUM(is_unjob) AS un_job_num,SUM(is_sendout) AS send_out_num,SUM(is_report) AS report_user_num,0 AS up_flag
  FROM temp_user_info 
  GROUP BY official_organ_id;
-- 更新到汇总表中                                                                                                                                      
START TRANSACTION;
  UPDATE rep_job_rate a 
    INNER JOIN temp_job_info b ON a.period=b.period AND a.organ_id=b.organ_id
    SET b.up_flag=1,a.un_job_num=b.un_job_num,a.send_out_num=b.send_out_num,a.report_user_num=b.report_user_num;
  INSERT INTO rep_job_rate(period,organ_id,official_actual_num,un_job_num,report_user_num,send_out_num)
    SELECT period,organ_id,0,un_job_num,report_user_num,send_out_num 
    FROM temp_job_info WHERE up_flag=0;
  -- 更新实有编制人数,当月取机构表 以往月份取机构编制表
  IF v_period>=CONVERT(DATE_FORMAT(NOW(),'%Y%m'),SIGNED) THEN
    UPDATE rep_job_rate a
      INNER JOIN arc_organ b ON a.period=v_period AND a.organ_id=b.id
      SET official_actual_num=b.actual_num;
  ELSE
    UPDATE rep_job_rate a
      INNER JOIN rep_keep_official b ON a.period=v_period AND a.period=b.period AND a.organ_id=b.organ_id AND b.depart_id=0
      SET official_actual_num=b.actual_num;
  END IF;
  -- 写入执行日志
  IF NOT EXISTS(SELECT id FROM sys_dispatch_log WHERE di_type=6 AND di_date=v_date) THEN
    INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo) VALUES(6,v_date,v_start_time,NOW(),'ok');
  ELSE
    UPDATE sys_dispatch_log SET begin_time=v_start_time,end_time=NOW() WHERE di_type=6 AND di_date=v_date;
  END IF;
COMMIT;

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

SET v_succ=1;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `pgene_review_rate` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE PROCEDURE `pgene_review_rate`(v_date date,out v_succ tinyint)
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
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `pgene_user_holiday` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_VALUE_ON_ZERO' */ ;
DELIMITER ;;
CREATE PROCEDURE `pgene_user_holiday`(v_begin_date date,v_date date,OUT v_succ tinyint)
BEGIN
DECLARE v_counter tinyint UNSIGNED;
DECLARE v_err_count int UNSIGNED;
DECLARE v_suc_count int UNSIGNED;
DECLARE v_start_time datetime;
DECLARE v_done INT DEFAULT 0;
DECLARE v_error INT DEFAULT 0;
DECLARE v_startday,v_endday date;
DECLARE v_in_time tinyint;
DECLARE v_holiday_classid bigint UNSIGNED;
DECLARE v_day_hour decimal(6,1);
DECLARE v_cost_time decimal(6,1);
DECLARE v_new_logid bigint UNSIGNED;
DECLARE v_new_queueid bigint UNSIGNED;
DECLARE v_public_holiday tinyint UNSIGNED;

DECLARE cv_id,cv_user_id,cv_dictionary_id bigint UNSIGNED;
DECLARE cv_reason varchar(50);
DECLARE cv_begin_date,cv_end_date date;
DECLARE cv_begin_time,cv_end_time tinyint;
DECLARE cv_days decimal(8,1) UNSIGNED;
DECLARE cv_used_count int;
DECLARE cv_last_date date;
DECLARE cv_un_audit_ids varchar(100);
DECLARE cv_organ_id,cv_depart_id,cv_official_depart_id bigint UNSIGNED;
DECLARE cur_holiday CURSOR FOR
  SELECT a.id,a.user_id,a.dictionary_id,a.reason,a.begin_date,a.begin_time,a.end_date,a.end_time,a.days,a.used_count,a.last_date,
    c.organ_id,c.depart_id,d.official_depart_id,IFNULL(e.current_auditor,'') AS un_audit_ids
    FROM arc_user_holiday a
    INNER JOIN sys_user c ON a.user_id=c.id AND c.is_job=1
    INNER JOIN arc_department d ON c.depart_id=d.id
	LEFT JOIN arc_user_extend e ON a.user_id=e.id
    WHERE a.end_date>=IFNULL(v_begin_date,v_Date) AND a.begin_date<=v_Date AND a.last_date<a.end_date;
DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done=1;
DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET v_error=1;
DECLARE EXIT HANDLER FOR SQLSTATE '45001' SET v_succ=0;
SET v_succ=1;
SET v_err_count=0;
SET v_suc_count=0;
SET v_start_time=NOW();

IF EXISTS(SELECT id FROM sys_dispatch_log WHERE di_date=v_Date AND di_type=2) THEN
    SIGNAL SQLSTATE '45001' SET MYSQL_ERRNO=2000,MESSAGE_TEXT='当前日期已执行过用户假期调度!';
END IF;
IF ISNULL(v_begin_date) THEN
  SET v_begin_date=v_date;
END IF;

SET v_day_hour=(SELECT CONVERT(parameter_value,DECIMAL) FROM sys_parameter WHERE code='0001');
IF ISNULL(v_day_hour) THEN
  SET v_day_hour=7;
END IF;

SET v_holiday_classid=(SELECT CONVERT(parameter_value,UNSIGNED) FROM sys_parameter WHERE code='0004');
IF ISNULL(v_holiday_classid) THEN
  SET v_holiday_classid=48;
END IF;

DROP TEMPORARY TABLE IF EXISTS temp_log_review;
CREATE TEMPORARY TABLE temp_log_review
SELECT a.id AS holiday_id,d.id AS user_id
    FROM arc_user_holiday a
    INNER JOIN sys_user b ON a.user_id=b.id AND b.is_job=1
    INNER JOIN arc_user_extend c ON a.user_id=c.id
    INNER JOIN sys_user d ON FIND_IN_SET(d.id,c.current_auditor)
    WHERE a.end_date>=IFNULL(v_begin_date,v_Date) AND a.begin_date<=v_Date AND a.last_date<a.end_date
	ORDER BY holiday_id;
ALTER TABLE temp_log_review ADD INDEX idx_temp_log_review(holiday_id);

OPEN cur_holiday;
REPEAT
  FETCH cur_holiday INTO cv_id,cv_user_id,cv_dictionary_id,cv_reason,cv_begin_date,cv_begin_time,cv_end_date,cv_end_time,cv_days,cv_used_count,cv_last_date,
    cv_organ_id,cv_depart_id,cv_official_depart_id,cv_un_audit_ids;
  IF NOT v_done THEN
    SET v_counter=0;


     SET v_startday=IFNULL(cv_begin_date,v_date);
     IF IFNULL(cv_last_date,'1900-01-01')>=v_startday THEN
       SET v_startday=ADDDATE(cv_last_date,INTERVAL 1 DAY);
     END IF;
     IF v_begin_date>v_startday THEN
       SET v_startday=v_begin_date;
     END IF;
     SET v_endday=cv_end_date;
     IF v_endday>v_date THEN
       SET v_endday=v_date;
     END IF;
     label_newday:
     WHILE (v_startday<=v_endday) AND (v_counter<=31) DO
       SET v_counter=v_counter+1;

       block1:BEGIN
         DECLARE v_no_record tinyint DEFAULT 0;
         DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_no_record=1;
         SELECT dictionary_id INTO v_public_holiday FROM work_user_state
           WHERE in_date=v_startday AND user_id=cv_user_id LIMIT 1;
         IF (v_no_record=1) OR (IFNULL(v_public_holiday,0) IN (1,2,3,4)) THEN
              UPDATE arc_user_holiday SET last_date=v_startday WHERE id=cv_id;
              SET v_startday=ADDDATE(v_startday,INTERVAL 1 DAY);
           ITERATE label_newday;
         END IF;
       END block1;
       SET v_in_time=(CASE WHEN v_startday=cv_end_date AND cv_end_time=0 THEN 0 WHEN v_startday=cv_begin_date AND cv_begin_time=1 THEN 1 ELSE 2 END);
       SET v_cost_time=(CASE WHEN v_in_time=2 THEN v_day_hour ELSE ROUND(v_day_hour/2,1) END);

       IF NOT EXISTS(SELECT * FROM work_log WHERE user_id=cv_user_id AND log_date=v_startday) THEN

         START TRANSACTION;
         SET v_error=0;

          INSERT INTO work_log(user_id,log_date,state,report_date,organ_id,depart_id,official_depart_id,cost_time,
            un_audit_ids,audited_ids,plan_ids,return_memo,term_type)
          VALUES(cv_user_id,v_startday,1,NOW(),cv_organ_id,cv_depart_id,cv_official_depart_id,v_cost_time,cv_un_audit_ids,'',CONCAT(cv_id,','),'',0);
          SET v_new_logid=@@identity;
          INSERT INTO work_log_list(work_log_id,sort,content,ampm,cost_time,time_type_id,space_type_id,
            attachment_ids,class_id,duty_id,plan_id,supervision_id,approval_id)
          VALUES(v_new_logid,1,cv_reason,ROUND(v_in_time+1,0),v_cost_time,1,0,'',v_holiday_classid,0,0,0,0);

          INSERT INTO work_review_queue(audit_user_id,log_date,work_log_id,organ_id,depart_id,user_id)
            SELECT user_id,v_startday,v_new_logid,cv_organ_id,cv_depart_id,cv_user_id
            FROM temp_log_review WHERE holiday_id=cv_id;

          INSERT INTO work_log_queue(user_id,log_date,organ_id,official_depart_id,cost_time,work_log_id,is_revoke,flag)
            VALUES(cv_user_id,v_startday,cv_organ_id,cv_official_depart_id,v_cost_time,v_new_logid,0,0);
          SET v_new_queueid=@@identity;

          INSERT INTO work_log_list_queue(queue_id,cost_time,time_type_id,space_type_id,
            class_id,duty_id,plan_id,supervision_id,approval_id)
          VALUES(v_new_queueid,v_cost_time,1,0,v_holiday_classid,0,0,0,0);

           IF NOT EXISTS(SELECT * FROM work_user_holiday_rec WHERE in_date=v_startday AND user_id=cv_user_id) THEN

             INSERT INTO work_user_holiday_rec (in_date,in_time,user_id,dictionary_id,reason,organ_id,depart_id,official_depart_id,user_holiday_id)
               VALUES(v_startday,v_in_time,cv_user_id,cv_dictionary_id,cv_reason,cv_organ_id,cv_depart_id,cv_official_depart_id,cv_id);
             UPDATE arc_user_holiday SET used_count=used_count+1,last_date=v_startday WHERE id=cv_id;

             UPDATE work_user_state SET holiday_time=IF(v_in_time=2,2,1),dictionary_id=7 WHERE in_date=v_startday AND user_id=cv_user_id AND dictionary_id=0;
           ELSE

             UPDATE arc_user_holiday SET last_date=v_startday WHERE id=cv_id;
           END IF;
         IF v_error=1 THEN
           ROLLBACK;
           SET v_err_count=v_err_count+1;
         ELSE
           COMMIT;
           SET v_suc_count=v_suc_count+1;
         END IF;
       ELSE

         UPDATE arc_user_holiday SET last_date=v_startday WHERE id=cv_id;
       END IF;

       SET v_startday=ADDDATE(v_startday,INTERVAL 1 DAY);
      END WHILE;

  END IF;
UNTIL v_done END REPEAT;
CLOSE cur_holiday;
IF v_succ=1 THEN

  INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo)
    VALUES(2,v_Date,v_start_time,NOW(),CONCAT('成功:',v_suc_count,'  失败:',v_err_count));
END IF;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `pgene_user_state` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_VALUE_ON_ZERO' */ ;
DELIMITER ;;
CREATE PROCEDURE `pgene_user_state`(v_Date date,OUT v_succ tinyint)
BEGIN
DECLARE v_start_time datetime;
DECLARE v_iswork tinyint;
DECLARE EXIT HANDLER FOR SQLEXCEPTION
BEGIN
  ROLLBACK;
END;
SET v_succ=0;
SET v_start_time=NOW();

IF EXISTS(SELECT id FROM sys_dispatch_log WHERE di_date=v_Date AND di_type=1) THEN
    SIGNAL SQLSTATE '45001' SET MYSQL_ERRNO=2000,MESSAGE_TEXT='当前日期已执行过用户状态记录!';
END IF;

DROP TEMPORARY TABLE IF EXISTS temp_users;
CREATE TEMPORARY TABLE temp_users
SELECT a.id,a.organ_id,a.depart_id,b.official_depart_id,a.is_assess,a.transfer_type_id,1 AS is_workday,1 AS is_must,0 AS dictionary_id,
  (CASE WHEN a.leader_flag=1 AND d.is_competent=1 THEN 1 ELSE 0 END) AS is_leader,a.official_type_id,
  0 as holiday_time
  FROM sys_user a
  INNER JOIN arc_department b ON a.depart_id=b.id
  INNER JOIN arc_organ d on a.organ_id=d.id and d.organ_type>0 and d.report_type in (1,3) AND d.is_revoked=0
  WHERE a.is_job=1 AND a.is_disabled=0 AND a.entry_date<=v_Date
  ORDER BY a.id;

UPDATE temp_users a
  INNER JOIN arc_organ_holiday b ON a.organ_id=b.organ_id AND b.end_date>=v_Date AND b.begin_date<=v_Date
  SET a.is_workday=0,a.is_must=0,a.dictionary_id=1;

SELECT is_work INTO v_iswork FROM arc_holiday WHERE adjust_date=v_Date;
IF NOT ISNULL(v_iswork) THEN

  IF v_iswork=0 THEN
    UPDATE temp_users a
      SET a.is_workday=0,a.is_must=0,a.dictionary_id=2
      WHERE a.dictionary_id=0;
  END IF;
ELSE

  DROP TEMPORARY TABLE IF EXISTS temp_organ_week;
  SET @sql=CONCAT('create temporary table temp_organ_week select organ_id,is_work',WEEKDAY(v_Date)+1,' as is_work from arc_organ_week');
  PREPARE stmt FROM @sql;
  EXECUTE stmt;
  DEALLOCATE PREPARE stmt;
  UPDATE temp_users a
    INNER JOIN temp_organ_week b ON a.organ_id=b.organ_id AND b.is_work=0
    SET a.is_workday=0,a.is_must=0,a.dictionary_id=3
    WHERE a.dictionary_id=0;

  IF WEEKDAY(v_Date)>=5 THEN
    UPDATE temp_users a
      LEFT JOIN temp_organ_week b ON a.organ_id=b.organ_id
      SET a.is_workday=0,a.is_must=0,a.dictionary_id=4
      WHERE a.dictionary_id=0 AND ISNULL(b.organ_id);
  END IF;
END IF;

UPDATE temp_users SET is_workday=0,is_must=0,dictionary_id=6 WHERE is_leader=1 AND dictionary_id=0;

UPDATE temp_users SET is_workday=0,is_must=0,dictionary_id=5 WHERE is_assess=0 AND dictionary_id=0;

START TRANSACTION;

  IF EXISTS(SELECT in_date FROM work_user_state WHERE in_date = v_Date LIMIT 1) THEN
    DELETE FROM work_user_state WHERE in_date=v_Date;
  END IF;

  INSERT INTO work_user_state (in_date,user_id,organ_id,depart_id,official_depart_id,
    is_assess,is_workday,is_must,is_report,transfer_type_id,dictionary_id,holiday_time)
    SELECT v_Date,`id`,organ_id,depart_id,official_depart_id,is_assess,is_workday,is_must,0,transfer_type_id,dictionary_id,holiday_time
    FROM temp_users;

  INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo) VALUES(1,v_Date,v_start_time,NOW(),'ok');
COMMIT;
SET v_succ=1;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `pgene_user_workplan` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE PROCEDURE `pgene_user_workplan`(v_begin_date date,v_date date,OUT v_succ tinyint)
BEGIN
DECLARE v_counter tinyint UNSIGNED;
DECLARE v_start_time datetime;
DECLARE v_err_count int UNSIGNED;
DECLARE v_suc_count int UNSIGNED;
DECLARE v_done INT DEFAULT 0;
DECLARE v_error INT DEFAULT 0;
DECLARE v_startday,v_endday date;
DECLARE v_day_hour decimal(6,1);
DECLARE v_new_logid bigint UNSIGNED;
DECLARE v_new_queueid bigint UNSIGNED;
DECLARE v_holiday_time tinyint UNSIGNED;
DECLARE v_in_time tinyint UNSIGNED;
DECLARE v_public_holiday tinyint UNSIGNED;
DECLARE v_organ_id,v_depart_id,v_official_depart_id bigint UNSIGNED;
DECLARE v_cost_time decimal(6,1);

DECLARE cv_plan_id,cv_user_id bigint UNSIGNED;
DECLARE cv_report_date datetime;
DECLARE cv_cost_time decimal(6,1);
DECLARE cv_auditor_ids,cv_audited_ids,cv_un_audit_ids varchar(100);
DECLARE cv_content varchar(200);
DECLARE cv_time_type_id,cv_space_type_id tinyint UNSIGNED;
DECLARE cv_begin_date,cv_end_date,cv_last_run_date date;
DECLARE cv_run_times int UNSIGNED;
DECLARE cv_class_id,cv_duty_id,cv_yearly_plan_id,cv_supervision_id,cv_approval_id bigint UNSIGNED;
DECLARE cur_plan CURSOR FOR
  SELECT b.id AS plan_id,b.user_id,NOW() AS report_date,b.cost_time,
    b.auditor_ids,b.auto_auditor_ids AS audited_ids,CONVERT('',char(100)) AS un_audit_ids,
    b.content,b.time_type_id,b.space_type_id,
    b.begin_date,b.end_date,b.last_run_date,b.run_times,
    b.class_id,b.duty_id,b.plan_id AS yearly_plan_id,b.supervision_id,b.approval_id
    FROM work_log_plan b
    WHERE b.end_date>=IFNULL(v_begin_date,v_Date) AND b.begin_date<=v_Date AND b.is_finish=0
    ORDER BY b.id;
DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done=1;
DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET v_error=1;
DECLARE EXIT HANDLER FOR SQLSTATE '45001' SET v_succ=0;
SET v_succ=1;
SET v_err_count=0;
SET v_suc_count=0;
SET v_start_time=NOW();

IF EXISTS(SELECT id FROM sys_dispatch_log WHERE di_date=v_Date AND di_type=3) THEN
    SIGNAL SQLSTATE '45001' SET MYSQL_ERRNO=2000,MESSAGE_TEXT='当前日期已执行过用户假期调度!';
END IF;
IF ISNULL(v_begin_date) THEN
  SET v_begin_date=v_date;
END IF;

SET v_day_hour=(SELECT CONVERT(parameter_value,DECIMAL) FROM sys_parameter WHERE code='0001');
IF ISNULL(v_day_hour) THEN
  SET v_day_hour=7;
END IF;

DROP TEMPORARY TABLE IF EXISTS temp_log_review;
CREATE TEMPORARY TABLE temp_log_review
SELECT a.id AS plan_id,b.id as user_id
    FROM work_log_plan a
		INNER JOIN sys_user b ON FIND_IN_SET(b.id,a.auditor_ids)
    WHERE a.end_date>=IFNULL(v_begin_date,v_Date) AND a.begin_date<=v_Date AND a.is_finish=0
		ORDER BY plan_id;
ALTER TABLE temp_log_review ADD INDEX idx_temp_log_review(plan_id);

OPEN cur_plan;
REPEAT
  FETCH cur_plan INTO cv_plan_id,cv_user_id,cv_report_date,cv_cost_time,
    cv_auditor_ids,cv_audited_ids,cv_un_audit_ids,
    cv_content,cv_time_type_id,cv_space_type_id,
    cv_begin_date,cv_end_date,cv_last_run_date,cv_run_times,
    cv_class_id,cv_duty_id,cv_yearly_plan_id,cv_supervision_id,cv_approval_id ;
  IF NOT v_done THEN
    SET v_counter=0;


     SET v_startday=IFNULL(cv_begin_date,v_date);
     IF IFNULL(cv_last_run_date,'1900-01-01')>=v_startday THEN
       SET v_startday=ADDDATE(cv_last_run_date,INTERVAL 1 DAY);
     END IF;
     IF v_begin_date>v_startday THEN
       SET v_startday=v_begin_date;
     END IF;
     SET v_endday=cv_end_date;
     IF v_endday>v_date THEN
       SET v_endday=v_date;
     END IF;
     label_newday:
     WHILE (v_startday<=v_endday) AND (v_counter<=31) DO
       SET v_counter=v_counter+1;

       block1:BEGIN
         DECLARE v_no_record tinyint DEFAULT 0;
         DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_no_record=1;
         SELECT organ_id,depart_id,official_depart_id,dictionary_id,holiday_time
           INTO v_organ_id,v_depart_id,v_official_depart_id,v_public_holiday,v_holiday_time
           FROM work_user_state
           WHERE in_date=v_startday AND user_id=cv_user_id LIMIT 1;
         IF (v_no_record=1) OR (IFNULL(v_public_holiday,0) IN (1,2,3,4)) THEN
           UPDATE work_log_plan SET last_run_date=v_startday,is_finish=IF(v_startday>=end_date,1,0) WHERE id=cv_plan_id;
           SET v_startday=ADDDATE(v_startday,INTERVAL 1 DAY);
           ITERATE label_newday;
         END IF;
       END block1;
       SET v_in_time=IF(v_holiday_time=1,1,3);
       SET v_cost_time=IF(v_holiday_time=1,3.5,cv_cost_time);

       IF NOT EXISTS(SELECT * FROM work_log WHERE user_id=cv_user_id AND log_date=v_startday) THEN


         START TRANSACTION;
         SET v_error=0;
          INSERT INTO work_log(user_id,log_date,state,report_date,organ_id,depart_id,official_depart_id,cost_time,
            un_audit_ids,audited_ids,plan_ids,return_memo,term_type)
          VALUES(cv_user_id,v_startday,1,cv_report_date,v_organ_id,v_depart_id,v_official_depart_id,v_cost_time,
            cv_auditor_ids,'',CONCAT(cv_plan_id,','),'',0);
          SET v_new_logid=@@identity;
          INSERT INTO work_log_list(work_log_id,sort,content,ampm,cost_time,time_type_id,space_type_id,
            attachment_ids,class_id,duty_id,plan_id,supervision_id,approval_id)
          VALUES(v_new_logid,1,cv_content,v_in_time,v_cost_time,cv_time_type_id,cv_space_type_id,
            '',cv_class_id,cv_duty_id,cv_yearly_plan_id,cv_supervision_id,cv_approval_id);

          INSERT INTO work_review_queue(audit_user_id,log_date,work_log_id,organ_id,depart_id,user_id)
            SELECT user_id,v_startday,v_new_logid,v_organ_id,v_depart_id,cv_user_id
            FROM temp_log_review WHERE plan_id=cv_plan_id;

          INSERT INTO work_log_queue(user_id,log_date,organ_id,official_depart_id,cost_time,work_log_id,is_revoke,flag)
            VALUES(cv_user_id,v_startday,v_organ_id,v_official_depart_id,v_cost_time,v_new_logid,0,0);
          SET v_new_queueid=@@identity;

          INSERT INTO work_log_list_queue(queue_id,cost_time,time_type_id,space_type_id,
            class_id,duty_id,plan_id,supervision_id,approval_id)
          VALUES(v_new_queueid,v_cost_time,cv_time_type_id,cv_space_type_id,
            cv_class_id,cv_duty_id,cv_yearly_plan_id,cv_supervision_id,cv_approval_id);

          UPDATE work_log_plan SET run_times=run_times+1,last_run_date=v_startday,is_finish=IF(v_startday>=end_date,1,0) WHERE id=cv_plan_id;
         IF v_error=1 THEN
           ROLLBACK;
           SET v_err_count=v_err_count+1;
         ELSE
           COMMIT;
           SET v_suc_count=v_suc_count+1;
         END IF;
       ELSE

         UPDATE work_log_plan SET last_run_date=v_startday,is_finish=IF(v_startday>=end_date,1,0) WHERE id=cv_plan_id;
       END IF;

       SET v_startday=ADDDATE(v_startday,INTERVAL 1 DAY);
      END WHILE;

  END IF;
UNTIL v_done END REPEAT;
CLOSE cur_plan;
IF v_succ=1 THEN

  INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo)
    VALUES(3,v_Date,v_start_time,NOW(),CONCAT('成功:',v_suc_count,'  失败:',v_err_count));
END IF;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `psum_work` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_VALUE_ON_ZERO' */ ;
DELIMITER ;;
CREATE PROCEDURE `psum_work`(v_Date date,OUT v_succ tinyint)
BEGIN
DECLARE v_start_time datetime;
DECLARE v_err_count int UNSIGNED;
DECLARE v_suc_count int UNSIGNED;
DECLARE v_iswork tinyint;
DECLARE v_firstday date;
DECLARE v_lastday date;
DECLARE v_period mediumint;
DECLARE v_dayhour decimal(12,2);
DECLARE v_holiday_classid bigint UNSIGNED;
DECLARE v_error tinyint UNSIGNED;
DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET v_error=1;
DECLARE EXIT HANDLER FOR SQLSTATE '45001' SET v_succ=0;
SET v_succ=0;
SET v_err_count=0;
SET v_suc_count=0;
SET v_start_time=NOW();
SET v_period=CONVERT(DATE_FORMAT(v_Date,'%Y%m'),SIGNED);
SELECT parameter_value INTO v_dayhour FROM sys_parameter WHERE code='0001';

IF EXISTS(SELECT id FROM sys_dispatch_log WHERE di_date=v_Date AND di_type=4) THEN
    SIGNAL SQLSTATE '45001' SET MYSQL_ERRNO=2000,MESSAGE_TEXT='当前日期已执行数据汇总!';
END IF;

SELECT CONVERT(parameter_value,UNSIGNED) INTO v_holiday_classid FROM sys_parameter WHERE code='0004';
IF ISNULL(v_holiday_classid) THEN
  SET v_holiday_classid=48;
END IF;


SELECT STR_TO_DATE(DATE_FORMAT(v_Date,'%Y-%m-01'),'%Y-%m-%d') INTO v_firstday;
SELECT LAST_DAY(v_Date) INTO v_lastday;
IF (v_Date=v_firstday) OR (v_Date=v_lastday) THEN
  START TRANSACTION;
  SET v_error=0;
      IF v_Date=v_lastday THEN
        DELETE FROM rep_keep_official WHERE period=v_period;
      ELSE
        IF EXISTS(SELECT id FROM rep_keep_official WHERE period=v_period LIMIT 1) THEN
          DELETE FROM rep_keep_official WHERE period=v_period;
        END IF;
      END IF;

      INSERT INTO rep_keep_official(period,organ_id,depart_id,official_num,actual_num)
        SELECT v_period,id,0 AS depart_id,official_num,actual_num FROM arc_organ
          WHERE organ_type<>0 AND is_revoked=0;
      INSERT INTO rep_keep_official(period,organ_id,depart_id,official_num,actual_num)
        SELECT v_period,a.organ_id,a.id,0 AS official_num,0 AS actual_num FROM arc_department a
          INNER JOIN arc_organ b ON a.organ_id=b.id AND b.is_revoked=0 AND b.organ_type<>0
          WHERE a.is_disabled=0 AND a.is_official=1;
   IF v_error=1 THEN
     ROLLBACK;
     SET v_err_count=v_err_count+1;
   ELSE
     COMMIT;
     SET v_suc_count=v_suc_count+1;
   END IF;
END IF;


IF ISNULL(v_dayhour) THEN
  SET v_dayhour=7;
END IF;

DROP TEMPORARY TABLE IF EXISTS temp_user_state;
CREATE TEMPORARY TABLE temp_user_state
  SELECT v_period AS period,
  organ_id,official_depart_id,user_id,SUM(is_workday) AS standard_days,
  SUM(is_workday)*v_dayhour AS standard_hours,0 AS up_flag
  FROM work_user_state
  WHERE in_date=v_Date AND is_assess=1 AND is_workday=1
  GROUP BY organ_id,official_depart_id,user_id
  ORDER BY organ_id,official_depart_id,user_id;
DROP TEMPORARY TABLE IF EXISTS temp_assess_depart;
CREATE TEMPORARY TABLE temp_assess_depart
  SELECT v_period AS period,organ_id,official_depart_id,SUM(standard_days) AS standard_days,SUM(standard_hours) AS standard_hours,0 AS up_flag
  FROM temp_user_state
  GROUP BY organ_id,official_depart_id
  ORDER BY organ_id,official_depart_id;
DROP TEMPORARY TABLE IF EXISTS temp_assess_organ;
CREATE TEMPORARY TABLE temp_assess_organ
  SELECT v_period AS period,organ_id,SUM(standard_days) AS standard_days,SUM(standard_hours) AS standard_hours,0 AS up_flag
  FROM temp_assess_depart
  GROUP BY organ_id
  ORDER BY organ_id;

START TRANSACTION;
SET v_error=0;

  UPDATE rep_assess_user a
    INNER JOIN temp_user_state b ON a.period=v_period AND a.organ_id=b.organ_id AND a.depart_id=b.official_depart_id AND a.user_id=b.user_id
    SET a.standard_days=a.standard_days+b.standard_days,a.standard_hours=a.standard_hours+b.standard_hours,b.up_flag=1;
  INSERT INTO rep_assess_user(period,organ_id,depart_id,user_id,standard_days,standard_hours)
    SELECT period,organ_id,official_depart_id,user_id,standard_days,standard_hours FROM temp_user_state WHERE up_flag=0;

  UPDATE rep_assess_depart a
    INNER JOIN temp_assess_depart b ON a.period=v_period AND a.organ_id=b.organ_id AND a.depart_id=b.official_depart_id
    SET a.standard_days=a.standard_days+b.standard_days,a.standard_hours=a.standard_hours+b.standard_hours,b.up_flag=1;
  INSERT INTO rep_assess_depart(period,organ_id,depart_id,standard_days,standard_hours)
    SELECT period,organ_id,official_depart_id,standard_days,standard_hours FROM temp_assess_depart WHERE up_flag=0;

  UPDATE rep_assess_organ a
    INNER JOIN temp_assess_organ b ON a.period=v_period AND a.organ_id=b.organ_id
    SET a.standard_days=a.standard_days+b.standard_days,a.standard_hours=a.standard_hours+b.standard_hours,b.up_flag=1;
  INSERT INTO rep_assess_organ(period,organ_id,standard_days,standard_hours)
    SELECT period,organ_id,standard_days,standard_hours FROM temp_assess_organ WHERE up_flag=0;
IF v_error=1 THEN
 ROLLBACK;
 SET v_err_count=v_err_count+1;
ELSE
 COMMIT;
 SET v_suc_count=v_suc_count+1;
END IF;

START TRANSACTION;
SET v_error=0;

UPDATE work_log_queue a
  INNER JOIN work_user_state b ON a.log_date=b.in_date AND a.user_id=b.user_id
  SET a.flag=1 WHERE a.flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_work_queue;
CREATE TEMPORARY TABLE temp_work_queue
SELECT a.id,MAX(a.user_id) AS user_id,MAX(a.log_date) AS log_date,
  MAX(c.organ_id) AS organ_id,MAX(c.official_depart_id) AS depart_id,
  MAX(a.work_log_id) AS work_log_id,MAX(a.is_revoke) AS is_revoke,
  IF(MAX(a.is_revoke=0),1,-1) AS actual_days,
  IF(MAX(a.is_revoke=0),COUNT(*),COUNT(*)*-1) AS actual_rows,
  IF(MAX(a.is_revoke=0),SUM(b.cost_time),SUM(b.cost_time)*-1) AS actual_hours,
  IF(MAX(c.is_workday=1),IF(MAX(a.is_revoke=0),1,-1),0) AS workday_days,
  IF(MAX(c.is_workday=1),IF(MAX(a.is_revoke=0),COUNT(*),COUNT(*)*-1),0) AS workday_rows,
  IF(MAX(c.is_workday=1),IF(MAX(a.is_revoke=0),SUM(b.cost_time),SUM(b.cost_time)*-1),0) AS workday_hours,
  IF(SUM(CASE WHEN b.time_type_id=2 THEN 1 ELSE 0 END)>=1,1,0)*IF(MAX(a.is_revoke=0),1,-1) AS overtime_days,
  SUM(CASE WHEN b.time_type_id=2 THEN 1 ELSE 0 END)*IF(MAX(a.is_revoke=0),1,-1) AS overtime_rows,
  SUM(CASE WHEN b.time_type_id=2 THEN b.cost_time ELSE 0 END)*IF(MAX(a.is_revoke=0),1,-1) AS overtime_hours,
  IF(SUM(CASE WHEN b.space_type_id=2 THEN 1 ELSE 0 END)>=1,1,0)*IF(MAX(a.is_revoke=0),1,-1) AS travel_days,
  SUM(CASE WHEN b.space_type_id=2 THEN 1 ELSE 0 END)*IF(MAX(a.is_revoke=0),1,-1) AS travel_rows,
  SUM(CASE WHEN b.space_type_id=2 THEN b.cost_time ELSE 0 END)*IF(MAX(a.is_revoke=0),1,-1) AS travel_hours,
  (CASE WHEN MAX(duty_id)>0 THEN 1 ELSE 0 END) AS pt_duty,
  (CASE WHEN MAX(plan_id)>0 THEN 1 ELSE 0 END) AS pt_yearly_plan,
  (CASE WHEN MAX(supervision_id)>0 THEN 1 ELSE 0 END) AS pt_supervision,
  (CASE WHEN MAX(approval_id)>0 THEN 1 ELSE 0 END) AS pt_approval,
  MAX(c.is_workday) AS is_workday,MAX(c.is_must) AS is_must,MAX(c.is_report) AS is_report,
  MAX(c.dictionary_id) AS dictionary_id,MAX(c.holiday_time) AS holiday_time
  FROM work_log_queue a
  INNER JOIN work_log_list_queue b ON a.id=b.queue_id AND b.class_id!=v_holiday_classid
  INNER JOIN work_user_state c ON a.log_date=c.in_date AND a.user_id=c.user_id AND c.is_assess=1
  WHERE a.flag=1
  GROUP BY a.id;

DROP TEMPORARY TABLE IF EXISTS temp_assess_user;
CREATE TEMPORARY TABLE temp_assess_user
SELECT CONVERT(DATE_FORMAT(log_date,'%Y%m'),SIGNED) AS period,
  organ_id,depart_id,user_id,
  SUM(actual_days) AS actual_days,SUM(actual_rows) AS actual_rows,SUM(actual_hours) AS actual_hours,
  SUM(workday_days) AS workday_days,SUM(workday_rows) AS workday_rows,SUM(workday_hours) AS workday_hours,
  SUM(overtime_days) AS overtime_days,SUM(overtime_rows) AS overtime_rows,SUM(overtime_hours) AS overtime_hours,
  SUM(travel_days) AS travel_days,SUM(travel_rows) AS travel_rows,SUM(travel_hours) AS travel_hours,
  SUM(pt_duty) AS pt_duty,SUM(pt_yearly_plan) AS pt_yearly_plan,SUM(pt_supervision) AS pt_supervision,
  SUM(pt_approval) AS pt_approval,0 AS up_flag
  FROM temp_work_queue
  GROUP BY period,organ_id,depart_id,user_id
  ORDER BY period,organ_id,depart_id,user_id;

UPDATE rep_assess_user a
  INNER JOIN temp_assess_user b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id
  SET b.up_flag=1,
  a.actual_days=a.actual_days+b.actual_days,a.actual_rows=a.actual_rows+b.actual_rows,a.actual_hours=a.actual_hours+b.actual_hours,
  a.workday_days=a.workday_days+b.workday_days,a.workday_rows=a.workday_rows+b.workday_rows,a.workday_hours=a.workday_hours+b.workday_hours,
  a.overtime_days=a.overtime_days+b.overtime_days,a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours,
  a.travel_days=a.travel_days+b.travel_days,a.travel_rows=a.travel_rows+b.travel_rows,a.travel_hours=a.travel_hours+b.travel_hours,
  a.pt_duty=a.pt_duty+b.pt_duty,a.pt_yearly_plan=a.pt_yearly_plan+b.pt_yearly_plan,
  a.pt_supervision=a.pt_supervision+b.pt_supervision,a.pt_approval=a.pt_approval+b.pt_approval;

INSERT INTO rep_assess_user(period,organ_id,depart_id,user_id,standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty,pt_yearly_plan,pt_supervision,pt_approval)
  SELECT period,organ_id,depart_id,user_id,0 AS standard_days,0 AS standard_hours,
    actual_days,actual_rows,actual_hours,
    workday_days,workday_rows,workday_hours,
    overtime_days,overtime_rows,overtime_hours,
    travel_days,travel_rows,travel_hours,
    pt_duty,pt_yearly_plan,pt_supervision,pt_approval
  FROM temp_assess_user WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_assess_depart;
CREATE TEMPORARY TABLE temp_assess_depart
SELECT period,organ_id,depart_id,
  SUM(actual_days) AS actual_days,SUM(actual_rows) AS actual_rows,SUM(actual_hours) AS actual_hours,
  SUM(workday_days) AS workday_days,SUM(workday_rows) AS workday_rows,SUM(workday_hours) AS workday_hours,
  SUM(overtime_days) AS overtime_days,SUM(overtime_rows) AS overtime_rows,SUM(overtime_hours) AS overtime_hours,
  SUM(travel_days) AS travel_days,SUM(travel_rows) AS travel_rows,SUM(travel_hours) AS travel_hours,
  SUM(pt_duty) AS pt_duty,SUM(pt_yearly_plan) AS pt_yearly_plan,SUM(pt_supervision) AS pt_supervision,
  SUM(pt_approval) AS pt_approval,0 AS up_flag
  FROM temp_assess_user
  GROUP BY period,organ_id,depart_id
  ORDER BY period,organ_id,depart_id;

UPDATE rep_assess_depart a
  INNER JOIN temp_assess_depart b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id
  SET b.up_flag=1,
  a.actual_days=a.actual_days+b.actual_days,a.actual_rows=a.actual_rows+b.actual_rows,a.actual_hours=a.actual_hours+b.actual_hours,
  a.workday_days=a.workday_days+b.workday_days,a.workday_rows=a.workday_rows+b.workday_rows,a.workday_hours=a.workday_hours+b.workday_hours,
  a.overtime_days=a.overtime_days+b.overtime_days,a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours,
  a.travel_days=a.travel_days+b.travel_days,a.travel_rows=a.travel_rows+b.travel_rows,a.travel_hours=a.travel_hours+b.travel_hours,
  a.pt_duty=a.pt_duty+b.pt_duty,a.pt_yearly_plan=a.pt_yearly_plan+b.pt_yearly_plan,
  a.pt_supervision=a.pt_supervision+b.pt_supervision,a.pt_approval=a.pt_approval+b.pt_approval;

INSERT INTO rep_assess_depart(period,organ_id,depart_id,standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty,pt_yearly_plan,pt_supervision,pt_approval)
  SELECT period,organ_id,depart_id,0 AS standard_days,0 AS standard_hours,
    actual_days,actual_rows,actual_hours,
    workday_days,workday_rows,workday_hours,
    overtime_days,overtime_rows,overtime_hours,
    travel_days,travel_rows,travel_hours,
    pt_duty,pt_yearly_plan,pt_supervision,pt_approval
  FROM temp_assess_depart WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_assess_organ;
CREATE TEMPORARY TABLE temp_assess_organ
SELECT period,organ_id,
  SUM(actual_days) AS actual_days,SUM(actual_rows) AS actual_rows,SUM(actual_hours) AS actual_hours,
  SUM(workday_days) AS workday_days,SUM(workday_rows) AS workday_rows,SUM(workday_hours) AS workday_hours,
  SUM(overtime_days) AS overtime_days,SUM(overtime_rows) AS overtime_rows,SUM(overtime_hours) AS overtime_hours,
  SUM(travel_days) AS travel_days,SUM(travel_rows) AS travel_rows,SUM(travel_hours) AS travel_hours,
  SUM(pt_duty) AS pt_duty,SUM(pt_yearly_plan) AS pt_yearly_plan,SUM(pt_supervision) AS pt_supervision,
  SUM(pt_approval) AS pt_approval,0 AS up_flag
  FROM temp_assess_depart
  GROUP BY period,organ_id
  ORDER BY period,organ_id;

UPDATE rep_assess_organ a
  INNER JOIN temp_assess_organ b
  ON a.period=b.period AND a.organ_id=b.organ_id
  SET b.up_flag=1,
  a.actual_days=a.actual_days+b.actual_days,a.actual_rows=a.actual_rows+b.actual_rows,a.actual_hours=a.actual_hours+b.actual_hours,
  a.workday_days=a.workday_days+b.workday_days,a.workday_rows=a.workday_rows+b.workday_rows,a.workday_hours=a.workday_hours+b.workday_hours,
  a.overtime_days=a.overtime_days+b.overtime_days,a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours,
  a.travel_days=a.travel_days+b.travel_days,a.travel_rows=a.travel_rows+b.travel_rows,a.travel_hours=a.travel_hours+b.travel_hours,
  a.pt_duty=a.pt_duty+b.pt_duty,a.pt_yearly_plan=a.pt_yearly_plan+b.pt_yearly_plan,
  a.pt_supervision=a.pt_supervision+b.pt_supervision,a.pt_approval=a.pt_approval+b.pt_approval;

INSERT INTO rep_assess_organ(period,organ_id,standard_days,standard_hours,
  actual_days,actual_rows,actual_hours,
  workday_days,workday_rows,workday_hours,
  overtime_days,overtime_rows,overtime_hours,
  travel_days,travel_rows,travel_hours,
  pt_duty,pt_yearly_plan,pt_supervision,pt_approval)
  SELECT period,organ_id,0 AS standard_days,0 AS standard_hours,
    actual_days,actual_rows,actual_hours,
    workday_days,workday_rows,workday_hours,
    overtime_days,overtime_rows,overtime_hours,
    travel_days,travel_rows,travel_hours,
    pt_duty,pt_yearly_plan,pt_supervision,pt_approval
  FROM temp_assess_organ WHERE up_flag=0;
IF v_error=1 THEN
 ROLLBACK;
 SET v_err_count=v_err_count+1;
ELSE
 COMMIT;
 SET v_suc_count=v_suc_count+1;
END IF;

DROP TEMPORARY TABLE IF EXISTS temp_work_queue;
CREATE TEMPORARY TABLE temp_work_queue
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,a.is_revoke,
  a.id,a.work_log_id,a.log_date,a.organ_id,a.official_depart_id AS depart_id,a.user_id,
  b.cost_time,b.time_type_id,b.class_id,b.duty_id,b.plan_id,b.supervision_id,b.approval_id
  FROM work_log_queue a
  INNER JOIN work_log_list_queue b ON a.id=b.queue_id AND b.class_id!=v_holiday_classid
  INNER JOIN work_user_state c ON a.log_date=c.in_date AND a.user_id=c.user_id AND c.is_assess=1
  WHERE a.flag=1;

START TRANSACTION;
SET v_error=0;
DROP TEMPORARY TABLE IF EXISTS temp_perform_approval;
CREATE TEMPORARY TABLE temp_perform_approval
SELECT period,organ_id,depart_id,user_id,approval_id,
  SUM(IF(is_revoke=0,1,-1)) AS work_rows,
  SUM(IF(is_revoke=0,cost_time,-1*cost_time)) AS work_hours,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,1,-1) ELSE 0 END) AS overtime_rows,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,cost_time,-1*cost_time) ELSE 0 END) AS overtime_hours,
  0 AS up_flag
  FROM temp_work_queue
  WHERE approval_id>=1
  GROUP BY period,organ_id,depart_id,user_id,approval_id
  ORDER BY period,organ_id,depart_id,user_id,approval_id;
UPDATE rep_perform_approval a
  INNER JOIN temp_perform_approval b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.approval_id=b.approval_id
  SET b.up_flag=1,
    a.work_rows=a.work_rows+b.work_rows,a.work_hours=a.work_hours+b.work_hours,
    a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours;
INSERT INTO rep_perform_approval(period,organ_id,depart_id,user_id,approval_id,
    work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT period,organ_id,depart_id,user_id,approval_id,
         work_rows,work_hours,overtime_rows,overtime_hours
  FROM temp_perform_approval WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_perform_duty;
CREATE TEMPORARY TABLE temp_perform_duty
SELECT period,organ_id,depart_id,user_id,duty_id,
  SUM(IF(is_revoke=0,1,-1)) AS work_rows,
  SUM(IF(is_revoke=0,cost_time,-1*cost_time)) AS work_hours,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,1,-1) ELSE 0 END) AS overtime_rows,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,cost_time,-1*cost_time) ELSE 0 END) AS overtime_hours,
  0 AS up_flag
  FROM temp_work_queue
  WHERE duty_id>=1
  GROUP BY period,organ_id,depart_id,user_id,duty_id
  ORDER BY period,organ_id,depart_id,user_id,duty_id;
UPDATE rep_perform_duty a
  INNER JOIN temp_perform_duty b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.duty_id=b.duty_id
  SET b.up_flag=1,
    a.work_rows=a.work_rows+b.work_rows,a.work_hours=a.work_hours+b.work_hours,
    a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours;
INSERT INTO rep_perform_duty(period,organ_id,depart_id,user_id,duty_id,
    work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT period,organ_id,depart_id,user_id,duty_id,
         work_rows,work_hours,overtime_rows,overtime_hours
  FROM temp_perform_duty WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_perform_supervision;
CREATE TEMPORARY TABLE temp_perform_supervision
SELECT period,organ_id,depart_id,user_id,supervision_id,
  SUM(IF(is_revoke=0,1,-1)) AS work_rows,
  SUM(IF(is_revoke=0,cost_time,-1*cost_time)) AS work_hours,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,1,-1) ELSE 0 END) AS overtime_rows,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,cost_time,-1*cost_time) ELSE 0 END) AS overtime_hours,
  0 AS up_flag
  FROM temp_work_queue
  WHERE supervision_id>=1
  GROUP BY period,organ_id,depart_id,user_id,supervision_id
  ORDER BY period,organ_id,depart_id,user_id,supervision_id;
UPDATE rep_perform_supervision a
  INNER JOIN temp_perform_supervision b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.supervision_id=b.supervision_id
  SET b.up_flag=1,
    a.work_rows=a.work_rows+b.work_rows,a.work_hours=a.work_hours+b.work_hours,
    a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours;
INSERT INTO rep_perform_supervision(period,organ_id,depart_id,user_id,supervision_id,
    work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT period,organ_id,depart_id,user_id,supervision_id,
         work_rows,work_hours,overtime_rows,overtime_hours
  FROM temp_perform_supervision WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_perform_yearly_plan;
CREATE TEMPORARY TABLE temp_perform_yearly_plan
SELECT period,organ_id,depart_id,user_id,plan_id,
  SUM(IF(is_revoke=0,1,-1)) AS work_rows,
  SUM(IF(is_revoke=0,cost_time,-1*cost_time)) AS work_hours,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,1,-1) ELSE 0 END) AS overtime_rows,
  SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,cost_time,-1*cost_time) ELSE 0 END) AS overtime_hours,
  0 AS up_flag
  FROM temp_work_queue
  WHERE plan_id>=1
  GROUP BY period,organ_id,depart_id,user_id,plan_id
  ORDER BY period,organ_id,depart_id,user_id,plan_id;
UPDATE rep_perform_yearly_plan a
  INNER JOIN temp_perform_yearly_plan b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.yearly_plan_id=b.plan_id
  SET b.up_flag=1,
    a.work_rows=a.work_rows+b.work_rows,a.work_hours=a.work_hours+b.work_hours,
    a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours;
INSERT INTO rep_perform_yearly_plan(period,organ_id,depart_id,user_id,yearly_plan_id,
    work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT period,organ_id,depart_id,user_id,plan_id,
         work_rows,work_hours,overtime_rows,overtime_hours
  FROM temp_perform_yearly_plan WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_perform_class;
CREATE TEMPORARY TABLE temp_perform_class
WITH etc_class AS (SELECT period,organ_id,depart_id,user_id,class_id,
    SUM(IF(is_revoke=0,1,-1)) AS work_rows,
    SUM(IF(is_revoke=0,cost_time,-1*cost_time)) AS work_hours,
    SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,1,-1) ELSE 0 END) AS overtime_rows,
    SUM(CASE WHEN time_type_id=2 THEN IF(is_revoke=0,cost_time,-1*cost_time) ELSE 0 END) AS overtime_hours
    FROM temp_work_queue
    WHERE class_id>=1
    GROUP BY period,organ_id,depart_id,user_id,class_id)
SELECT a1.*,b1.level1_id,0 AS up_flag
  FROM etc_class a1
  INNER JOIN (SELECT id,IF(depth=1,id,CONVERT(SUBSTRING(SUBSTRING_INDEX(parent_ids,',',2),3),UNSIGNED)) AS level1_id FROM arc_work_class) b1 ON a1.class_id=b1.id
  ORDER BY period,organ_id,depart_id,user_id,class_id;
UPDATE rep_perform_class a
  INNER JOIN temp_perform_class b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.level2_id=b.class_id
  SET b.up_flag=1,
    a.work_rows=a.work_rows+b.work_rows,a.work_hours=a.work_hours+b.work_hours,
    a.overtime_rows=a.overtime_rows+b.overtime_rows,a.overtime_hours=a.overtime_hours+b.overtime_hours;
INSERT INTO rep_perform_class(period,organ_id,depart_id,user_id,level1_id,level2_id,
    work_rows,work_hours,overtime_rows,overtime_hours)
  SELECT period,organ_id,depart_id,user_id,level1_id,class_id,
         work_rows,work_hours,overtime_rows,overtime_hours
  FROM temp_perform_class WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_class_pre;
CREATE TEMPORARY TABLE temp_class_pre
SELECT a.id,b.level1_id as class_id,IF(MAX(a.is_revoke)=0,1,-1) AS pt_num
  FROM temp_work_queue a
  INNER JOIN (SELECT id,IF(depth=1,id,CONVERT(SUBSTRING(SUBSTRING_INDEX(parent_ids,',',2),3),UNSIGNED)) AS level1_id FROM arc_work_class) b ON a.class_id=b.id
  GROUP BY a.id,b.level1_id;

DROP TEMPORARY TABLE IF EXISTS temp_perform_class_pt;
CREATE TEMPORARY TABLE temp_perform_class_pt
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,
  a.organ_id,a.official_depart_id AS depart_id,a.user_id,b.class_id,SUM(pt_num) AS pt_num,0 AS up_flag
  FROM work_log_queue a
  INNER JOIN temp_class_pre b ON a.id=b.id
  GROUP BY period,organ_id,depart_id,user_id,class_id
  ORDER BY period,organ_id,depart_id,user_id,class_id;

UPDATE rep_perform_class_pt a
  INNER JOIN temp_perform_class_pt b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.class_id=b.class_id
  SET b.up_flag=1,a.pt_num=a.pt_num+b.pt_num;

INSERT INTO rep_perform_class_pt(period,organ_id,depart_id,user_id,class_id,pt_num)
SELECT period,organ_id,depart_id,user_id,class_id,pt_num
  FROM temp_perform_class_pt
  WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_class_pre;
CREATE TEMPORARY TABLE temp_class_pre
SELECT id,class_id,IF(MAX(is_revoke)=0,1,-1) AS pt_num
  FROM temp_work_queue
  WHERE class_id>=1
  GROUP BY id,class_id;

DROP TEMPORARY TABLE IF EXISTS temp_perform_class_ptl;
CREATE TEMPORARY TABLE temp_perform_class_ptl
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,
  a.organ_id,a.official_depart_id AS depart_id,a.user_id,b.class_id,SUM(pt_num) AS pt_num,0 AS up_flag
  FROM work_log_queue a
  INNER JOIN temp_class_pre b ON a.id=b.id
  GROUP BY period,organ_id,depart_id,user_id,class_id
  ORDER BY period,organ_id,depart_id,user_id,class_id;

UPDATE rep_perform_class_ptl a
  INNER JOIN temp_perform_class_ptl b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.depart_id=b.depart_id AND a.user_id=b.user_id AND a.class_id=b.class_id
  SET b.up_flag=1,a.pt_num=a.pt_num+b.pt_num;

INSERT INTO rep_perform_class_ptl(period,organ_id,depart_id,user_id,class_id,pt_num)
SELECT period,organ_id,depart_id,user_id,class_id,pt_num
  FROM temp_perform_class_ptl
  WHERE up_flag=0;
IF v_error=1 THEN
 ROLLBACK;
 SET v_err_count=v_err_count+1;
ELSE
 COMMIT;
 SET v_suc_count=v_suc_count+1;
END IF;



START TRANSACTION;
SET v_error=0;
DROP TEMPORARY TABLE IF EXISTS temp_user_target;
CREATE TEMPORARY TABLE temp_user_target
SELECT CONVERT(DATE_FORMAT(in_date,'%Y%m'),UNSIGNED) AS period,organ_id,official_depart_id,user_id,
  count(*) AS target_number,0 AS up_flag
  FROM work_user_state
  WHERE in_date=v_Date AND is_must=1
  GROUP BY period,organ_id,official_depart_id,user_id;
UPDATE rep_rate_user a
  INNER JOIN temp_user_target b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.official_depart_id=b.official_depart_id AND a.user_id=b.user_id
  SET b.up_flag=1,a.target_number=a.target_number+b.target_number;
INSERT INTO rep_rate_user(period,organ_id,official_depart_id,user_id,target_number,actual_number)
  SELECT period,organ_id,official_depart_id,user_id,target_number,0
  FROM temp_user_target WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_depart_target;
CREATE TEMPORARY TABLE temp_depart_target
SELECT period,organ_id,official_depart_id,
  SUM(target_number) AS target_number,0 AS up_flag
  FROM temp_user_target
  GROUP BY period,organ_id,official_depart_id;
UPDATE rep_rate_depart a
  INNER JOIN temp_depart_target b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.official_depart_id=b.official_depart_id
  SET b.up_flag=1,a.target_number=a.target_number+b.target_number;
INSERT INTO rep_rate_depart(period,organ_id,official_depart_id,target_number,actual_number)
  SELECT period,organ_id,official_depart_id,target_number,0
  FROM temp_depart_target WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_organ_target;
CREATE TEMPORARY TABLE temp_organ_target
SELECT period,organ_id,SUM(target_number) AS target_number,0 AS up_flag
  FROM temp_depart_target
  GROUP BY period,organ_id;
UPDATE rep_rate_organ a
  INNER JOIN temp_organ_target b
  ON a.period=b.period AND a.organ_id=b.organ_id

  SET b.up_flag=1,a.target_number=a.target_number+b.target_number;
INSERT INTO rep_rate_organ(period,organ_id,target_number,actual_number)
  SELECT period,organ_id,target_number,0
  FROM temp_organ_target WHERE up_flag=0;


DROP TEMPORARY TABLE IF EXISTS temp_user_isreport;
CREATE TEMPORARY TABLE temp_user_isreport
SELECT log_date,user_id,SUM(IF(is_revoke=0,1,-1)) AS actual_number
    FROM work_log_queue WHERE flag=1
    GROUP BY log_date,user_id
    ORDER BY log_date,user_id;
UPDATE work_user_state a
  INNER JOIN temp_user_isreport b ON a.in_date=b.log_date AND a.user_id=b.user_id
  SET a.is_report=IF((a.is_report+b.actual_number)>=1,1,0);

DROP TEMPORARY TABLE IF EXISTS temp_user_actual;
CREATE TEMPORARY TABLE temp_user_actual
SELECT CONVERT(DATE_FORMAT(a.log_date,'%Y%m'),UNSIGNED) AS period,
  b.organ_id,b.official_depart_id,a.user_id,SUM(actual_number) AS actual_number,0 AS up_flag
  FROM temp_user_isreport a
  INNER JOIN work_user_state b ON a.log_date=b.in_date AND a.user_id=b.user_id AND b.is_must=1
  GROUP BY period,organ_id,official_depart_id,user_id;
UPDATE rep_rate_user a
  INNER JOIN temp_user_actual b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.official_depart_id=b.official_depart_id AND a.user_id=b.user_id
  SET a.actual_number=a.actual_number+b.actual_number,b.up_flag=1;
INSERT INTO rep_rate_user (period,organ_id,official_depart_id,user_id,target_number,actual_number)
  SELECT period,organ_id,official_depart_id,user_id,0,actual_number FROM temp_user_actual WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_depart_actual;
CREATE TEMPORARY TABLE temp_depart_actual
SELECT period,organ_id,official_depart_id,SUM(actual_number) AS actual_number,0 AS up_flag
  FROM temp_user_actual
  GROUP BY period,organ_id,official_depart_id;
UPDATE rep_rate_depart a
  INNER JOIN temp_depart_actual b
  ON a.period=b.period AND a.organ_id=b.organ_id AND a.official_depart_id=b.official_depart_id
  SET a.actual_number=a.actual_number+b.actual_number,b.up_flag=1;
INSERT INTO rep_rate_depart (period,organ_id,official_depart_id,target_number,actual_number)
  SELECT period,organ_id,official_depart_id,0,actual_number FROM temp_depart_actual WHERE up_flag=0;

DROP TEMPORARY TABLE IF EXISTS temp_organ_actual;
CREATE TEMPORARY TABLE temp_organ_actual
SELECT period,organ_id,SUM(actual_number) AS actual_number,0 AS up_flag
  FROM temp_depart_actual
  GROUP BY period,organ_id;
UPDATE rep_rate_organ a
  INNER JOIN temp_organ_actual b
  ON a.period=b.period AND a.organ_id=b.organ_id
  SET a.actual_number=a.actual_number+b.actual_number,b.up_flag=1;
INSERT INTO rep_rate_organ (period,organ_id,target_number,actual_number)
  SELECT period,organ_id,0,actual_number FROM temp_organ_actual WHERE up_flag=0;

IF v_error=1 THEN
 ROLLBACK;
 SET v_err_count=v_err_count+1;
ELSE
 COMMIT;
 SET v_suc_count=v_suc_count+1;
END IF;

DELETE a,b
  FROM work_log_queue a
  INNER JOIN work_log_list_queue b ON a.id=b.queue_id
  WHERE a.flag=1;

DELETE a,b
  FROM work_log_queue a
  INNER JOIN work_log_list_queue b ON a.id=b.queue_id
  INNER JOIN work_user_state c ON a.log_date=c.in_date AND a.user_id=c.user_id AND c.is_assess=0;
SET v_succ=1;

INSERT INTO sys_dispatch_log(di_type,di_date,begin_time,end_time,memo)
  VALUES(4,v_Date,v_start_time,NOW(),CONCAT('成功:',v_suc_count,'  失败:',v_err_count));
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-18 13:00:51
