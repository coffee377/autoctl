-- 重复导入数据处理

-- 269f6283090b608ef6c9b9578b8e2839
-- e7e54b941bc7efea0f8295f84a83c0a2
-- “天府救助通”四川省低收入人口动态监测信息平台电子签章设备采购项目
update bid_info i
set i.project_id = 'e7e54b941bc7efea0f8295f84a83c0a2'
where i.project_id = '269f6283090b608ef6c9b9578b8e2839';

delete
from bid_project p
where p.name = '“天府救助通”四川省低收入人口动态监测信息平台电子签章设备采购项目'
  and p.source = '0'
  and id = '269f6283090b608ef6c9b9578b8e2839';

-- c340608fc1fe35a81cb43f04f4385a94
-- 08ad63cb7cdd8377b5f1538e561a8df4
-- 定远县卫健委基本公共卫生服务信息系统升级改造项目
update bid_info i
set i.project_id = '08ad63cb7cdd8377b5f1538e561a8df4'
where i.project_id = 'c340608fc1fe35a81cb43f04f4385a94';

delete
from bid_project p
where p.name = '定远县卫健委基本公共卫生服务信息系统升级改造项目'
  and p.source = '0'
  and id = 'c340608fc1fe35a81cb43f04f4385a94';
