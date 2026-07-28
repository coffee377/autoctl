-- Modify "bid_apply" table
ALTER TABLE `bid_apply` ADD COLUMN `registration_status` enum('RP','RO','RF','RS') NULL COMMENT "报名情况 RP:待报名 RO:报名中 RF:报名失败 RS:报名成功" AFTER `remark`, ADD COLUMN `registration_failure_desc` longtext NULL COMMENT "报名失败描述" AFTER `registration_status`;
