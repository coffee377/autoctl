-- Modify "bid_expense" table
ALTER TABLE `bid_expense` ADD COLUMN `remark` longtext NULL COMMENT "备注" AFTER `deleted`;
