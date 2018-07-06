create database if not exists test;
create databases test;
CREATE TABLE if not exists `arc_duty_relate` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
	`organ_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '机构id',
	`depart_id` BIGINT(20) NOT NULL,
	`duty_id` BIGINT(20) NOT NULL,
	`category_id` BIGINT(20) NULL DEFAULT NULL COMMENT '(字典15)分类id',
	`name` VARCHAR(200) NOT NULL COMMENT '事项名称',
	`outline` VARCHAR(20) NOT NULL COMMENT '纲要',
	`sort` SMALLINT(4) NOT NULL DEFAULT '0',
	`category` VARCHAR(20) NOT NULL COMMENT '分类',
	`is_disabled` TINYINT(4) NOT NULL DEFAULT '0',
	`disable_date` DATE NOT NULL DEFAULT '2999-12-31' COMMENT '停用日期',
	`gmt_create` DATETIME(6) NOT NULL DEFAULT now(),
	`gmt_modified` DATETIME(6) NOT NULL DEFAULT '1900-01-01 00:00:00.000000',
	PRIMARY KEY (`id`)
)
COMMENT='基础_岗位职责关联事项'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;


CREATE TABLE if not exists `work_log_list_dutyrel` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
	`list_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '日志listid',
	`dutyrel_id` BIGINT(20) NULL DEFAULT NULL COMMENT '辅助关联项id',
	PRIMARY KEY (`id`)
)
COMMENT='基础_岗位职责关联事项'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;

REPLACE INTO `sys_enum` (`id`, `name`, `memo`, `gmt_create`, `gmt_modified`) VALUES (15, '岗位职责关联项', '安全生产,污染防治,招商引资等等等', '2018-06-23 14:48:11.887779', '1900-01-01 00:00:00.000000');
INSERT ignore INTO `sys_enum_item` (`enum_id`, `item_id`, `name`, `value`, `sort`, `is_disabled`, `gmt_create`, `gmt_modified`) VALUES 
(15, 1, '信访工作', 'letterandcall', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:38.138364'),
(15, 2, '污染防治', 'Pollutioncontrol', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:36.741568'),
(15, 3, '安全生产', 'Worksafety', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:32.882025'),
(15, 4, '放管服改革', 'Reforms', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:30.740323'),
(15, 5, '招商引资', 'Businessivitation', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:28.644530'),
(15, 6, '招才引智', 'Intelligenceimport', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:26.646752'),
(15, 7, '精准扶贫', 'Povertyalleviation', 0, 0, '2018-06-23 14:49:36.903621', '2018-06-25 15:23:24.270038');
