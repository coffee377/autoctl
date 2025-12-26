-- Modify "bid_expense" table
ALTER TABLE `bid_expense`
    MODIFY COLUMN `guarantee_deadline` datetime NULL COMMENT "保函期限" after `pay_method`,
    MODIFY COLUMN `plan_pay_time` datetime NULL COMMENT "计划（最迟）转账时间",
    ADD COLUMN `guarantee_denomination` decimal(16,2) NOT NULL DEFAULT 0.00 COMMENT "保函面额（元）" after `guarantee_deadline`,
    ADD COLUMN `pay_time` datetime(3) NULL COMMENT "付款时间" AFTER `plan_pay_time`;