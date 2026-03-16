-- Modify "bid_member_account" table
ALTER TABLE `bid_member_account` MODIFY COLUMN `primary_ca_id` varchar(32) NULL COMMENT "主证书 ID";
