-- Modify "bid_info" table
ALTER TABLE `bid_info`
    MODIFY COLUMN `bid_subject_code` varchar(32) NULL COMMENT "投标主体编码",
    MODIFY COLUMN `bid_subject_name` varchar(32) NULL COMMENT "投标主体名称",
    ADD COLUMN `group_leader`      varchar(8) NOT NULL COMMENT "投标组长工号" AFTER `id`,
    ADD COLUMN `group_leader_name` varchar(8) NOT NULL COMMENT "投标组长" AFTER `group_leader`;