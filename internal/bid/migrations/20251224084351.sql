-- Modify "bid_info" table
ALTER TABLE `bid_info` MODIFY COLUMN `group_leader` varchar(8) NULL DEFAULT "" COMMENT "投标组长工号", MODIFY COLUMN `group_leader_name` varchar(8) NULL DEFAULT "" COMMENT "投标组长";
