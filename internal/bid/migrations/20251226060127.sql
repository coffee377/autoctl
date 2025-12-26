-- 更新历史付款时间字段
update `bid_expense`
set `pay_time` = `plan_pay_time`
where `pay_time` is null;