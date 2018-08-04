drop table if exists arc_approval;
CREATE TABLE `arc_approval` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `organ_id` bigint(20) unsigned NOT NULL COMMENT '机构id',
  `name` varchar(200) NOT NULL COMMENT '审批(服务)事项名称',
  `outline` varchar(20) NOT NULL DEFAULT '' COMMENT '纲要',
  `sort` smallint(4) NOT NULL DEFAULT 0 COMMENT '显示序号',
  `category` varchar(20) NOT NULL DEFAULT '' COMMENT '分类',
  `is_disabled` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否停用',
  `disable_date` date NOT NULL DEFAULT '2999-12-31' COMMENT '停用日期',
  `gmt_create` datetime(6) NOT NULL DEFAULT current_timestamp(6) COMMENT '创建时间',
  `gmt_modified` datetime(6) NOT NULL DEFAULT '1900-01-01 00:00:00.000000' ON UPDATE current_timestamp(6) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_organ_id` (`organ_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8 COMMENT='基础_审批服务事项';
drop table if exists arc_assessment_state;
CREATE TABLE `arc_assessment_state` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `organ_id` bigint(20) DEFAULT NULL,
  `item_id` bigint(20) DEFAULT NULL,
  `yearmonth` varchar(255) DEFAULT NULL,
  `assessment_state` varchar(10) DEFAULT '',
  `is_assessment` varchar(10) DEFAULT '' COMMENT '整体考核状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=249 DEFAULT CHARSET=utf8;
drop table if exists arc_common_indicator_type_dict;
drop table if exists arc_category;
CREATE TABLE `arc_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '类别名称',
  `organ_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
CREATE TABLE `arc_common_indicator_type_dict` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `type_code` bigint(11) unsigned DEFAULT NULL COMMENT '类别code',
  `gmt_create` datetime(6) NOT NULL DEFAULT current_timestamp(6) COMMENT '创建时间',
  `gmt_modified` datetime(6) NOT NULL DEFAULT '1900-01-01 00:00:00.000000' ON UPDATE current_timestamp(6) COMMENT '更新时间',
  `design_year` varchar(6) DEFAULT '' COMMENT '共性指标定义年份',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '考核内容',
  `organ_id` bigint(20) NOT NULL,
  `score` bigint(20) DEFAULT NULL,
  `yearmonth` varchar(20) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5181 DEFAULT CHARSET=utf8 COMMENT='绩效考核共性指标评分标准字典表';
drop table if exists arc_common_indicator_score_refer;
CREATE TABLE `arc_common_indicator_score_refer` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `type_value` varchar(100) NOT NULL COMMENT '评分标准内容',
  `fraction` bigint(5) unsigned NOT NULL COMMENT '分数',
  `gmt_create` datetime(6) NOT NULL DEFAULT current_timestamp(6) COMMENT '创建时间',
  `gmt_modified` datetime(6) NOT NULL DEFAULT '1900-01-01 00:00:00.000000' ON UPDATE current_timestamp(6) COMMENT '更新时间',
  `item_id` bigint(20) DEFAULT NULL,
  `sort` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=231 DEFAULT CHARSET=utf8 COMMENT='绩效考核共性指标评分标准内容表';
