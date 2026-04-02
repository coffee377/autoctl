-- Modify "bid_expense" table
ALTER TABLE `bid_expense` ADD COLUMN `deleted` bool NOT NULL DEFAULT 0 COMMENT "是否逻辑删除" AFTER `done`;
