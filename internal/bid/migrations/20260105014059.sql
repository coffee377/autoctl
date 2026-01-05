-- 0822	刘全乐
select p.id, p.name, i.group_leader, i.group_leader_name, a.business_id, i.created_at, i.created_by
from bid_info i
         left join bid_project p on p.id = i.project_id
         left join bid_apply a on a.project_id = i.project_id
where p.name like '%清镇市%';

select *
from bid_expense e
         inner join bid_project p on p.id = e.project_id
where p.name like '%清镇市%';

-- 清除清镇市已中标拆分项目数据
delete
from bid_info i
-- 清镇市第一人民医院（医共体）基本公共卫生系统接口服务（二次发布）、清镇市中医医院（医共体）基本公共卫生系统提供的接口服务
-- 11db67f157f33cd61251660ff31c31f7
where i.project_id = '11db67f157f33cd61251660ff31c31f7';

-- 更新已拆分数据组长信息
-- 551ac05d337ccc4a743d6f2a4808b625	清镇市第一人民医院（医共体）基本公共卫生系统接口服务（二次发布）
-- 76ef08388e60e1659107a13f43a21948	清镇市中医医院（医共体）基本公共卫生系统提供的接口服务
update bid_info i
set i.group_leader      = '0822',
    i.group_leader_name = '刘全乐'
where i.bid_status = 'W'
  and i.project_id in ('551ac05d337ccc4a743d6f2a4808b625', '76ef08388e60e1659107a13f43a21948');

-- 将申请记录拆分
-- 11db67f157f33cd61251660ff31c31f7
--  - 551ac05d337ccc4a743d6f2a4808b625	清镇市第一人民医院（医共体）基本公共卫生系统接口服务（二次发布）
--  - 76ef08388e60e1659107a13f43a21948	清镇市中医医院（医共体）基本公共卫生系统提供的接口服务
update bid_apply a
set a.project_id = '551ac05d337ccc4a743d6f2a4808b625'
where a.business_id = '202507141027000501769'
  and a.project_id = '11db67f157f33cd61251660ff31c31f7';

-- 清镇市申请数据拆分一份
delete from bid_apply a where a.id = '607431c3e9d311f0a2d458112299b16c';
insert into bid_apply
select '607431c3e9d311f0a2d458112299b16c' as id,
       business_id,
       instance_id,
       purchaser_name,
       bid_type,
       agency_name,
       agency_contact,
       opening_date,
       notice_url,
       budget_amount,
       remark,
       attachments,
       approval_status,
       done,
       created_at,
       created_by,
       updated_at,
       updated_by,
       '76ef08388e60e1659107a13f43a21948' as project_id
from bid_apply a
where a.business_id = '202507141027000501769';

-- 平分 50000 预算
update bid_apply a
set a.budget_amount = 25000
where a.business_id = '202507141027000501769';

-- 清镇市中医医院采购人更新
update bid_apply a
set a.purchaser_name = '清镇市中医医院'
where a.business_id = '202507141027000501769'
  and a.id = '607431c3e9d311f0a2d458112299b16c';