-- 序号自增触发器
drop trigger if exists trg_sys_task_log_assign_seq;

create trigger trg_sys_task_log_assign_seq
    before insert
    on sys_task_log
    for each row
begin
    DECLARE max_seq INT;

    -- 查询当前请购单的最大序号（无需 FOR UPDATE，BEFORE INSERT 天然避免并发重复）
    SELECT IFNULL(MAX(assign_seq), 0)
    INTO max_seq
    FROM sys_task_log
    WHERE biz_type = NEW.biz_type
      and biz_id = NEW.biz_id;

    -- 插入前直接赋值给 NEW.assign_seq（此时数据还未写入表，可修改）
    SET NEW.assign_seq = max_seq + 1;
end;