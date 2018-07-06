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