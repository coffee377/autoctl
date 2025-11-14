-- 设置时区
SET time_zone = 'Asia/Shanghai';

-- 1. 更新支出信息项目 ID
update bid_expense e
    inner join bid_project p on e.project_name = p.name and e.project_code = p.code
set e.project_id = p.id,
    e.updated_at = current_timestamp(3),
    e.updated_by = '0634'
where e.project_id is null
  and e.project_name = '安徽省某单位产品及服务030040甄选项目'
  and e.project_code = 'ZYJC-ZX-202503-0135';

-- 2. 费用支出表缺失的项目信息
SET @json_data = '[
  {
    "expense_business_id": "202501081140000199187",
    "id": "9a8ac6d7993413d53988ffeb3cccea76",
    "name": "合肥水务集团2024年营商环境相关功能新增",
    "code": "2024BFFWZ02321",
    "type": "S",
    "source": "0",
    "department_code": "00700",
    "department_name": "商务部",
    "biz_rep_no": "0028",
    "biz_rep_name": "杨彬彬",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202502111438000307157",
    "id": "99ee08cd07bca30fd91edc0c316788eb",
    "name": "石台县卫生健康委员会基本公共卫生服务信息系统升级改造项目",
    "code": "/",
    "type": "S",
    "source": "0",
    "department_code": "03900",
    "department_name": "智慧医疗事业部",
    "biz_rep_no": "0854",
    "biz_rep_name": "李杰",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202502211040000573315",
    "id": "d6b73b2510ff75a35cd4b1f08049a547",
    "name": "阿里地区藏医院“智慧医院”合作建设项目",
    "code": "XZCQ-2025-02",
    "type": "S",
    "source": "0",
    "department_code": "03900",
    "department_name": "智慧医疗事业部",
    "biz_rep_no": "/",
    "biz_rep_name": "/",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202502261602000472320",
    "id": "0e16be946af4138393a2e1c2c0b6f9cd",
    "name": "阿里地区藏医院“智慧医院” 合作建设项目",
    "code": "ZY25-Z227-FW010",
    "type": "S",
    "source": "0",
    "department_code": "03900",
    "department_name": "智慧医疗事业部",
    "biz_rep_no": "0007",
    "biz_rep_name": "卢秀良",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202504071548000163298",
    "id": "ae8d53bbcff226a79b330b99868a60e9",
    "name": "山东信达物联应用技术有限公司供应商入库",
    "code": "/",
    "type": "S",
    "source": "0",
    "department_code": "04100",
    "department_name": "数字政企事业部",
    "biz_rep_no": "0076",
    "biz_rep_name": "曹运节",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202508110959000594781",
    "id": "5aabcc6042ad2ee37f1cbd8cbbe3ed0e",
    "name": "威信县紧密型医共体服务平台建设项目",
    "code": "YNLL-202507172",
    "type": "S",
    "source": "0",
    "department_code": "00700",
    "department_name": "商务部",
    "biz_rep_no": "0028",
    "biz_rep_name": "杨彬彬",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202510210902000448045",
    "id": "75a754c96227518668a710f6a526437c",
    "name": "威信县紧密型医共体服务平台建设项目（二次）",
    "code": "ZTZC2025-G1-01155-YNLL-0041",
    "type": "UP",
    "source": "0",
    "department_code": "00700",
    "department_name": "商务部",
    "biz_rep_no": "0028",
    "biz_rep_name": "杨彬彬",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202510281358000460154",
    "id": "44f494f6977a7af73ec0b4d23c23f292",
    "name": "肇源县医共体平台升级改造项目（二次）",
    "code": "[230622]QC[GK]20250002-1",
    "type": "UP",
    "source": "0",
    "department_code": "03900",
    "department_name": "智慧医疗事业部",
    "biz_rep_no": "0007",
    "biz_rep_name": "卢秀良",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  },
  {
    "expense_business_id": "202511031442000096289",
    "id": "2979db955920e2043075b3622fd00a69",
    "name": "医院信息管理系统项目（安徽省未成年人强制隔离戒毒所）",
    "code": "/",
    "type": "S",
    "source": "0",
    "department_code": "04000",
    "department_name": "智慧医院事业部",
    "biz_rep_no": "0053",
    "biz_rep_name": "林宏",
    "created_at": "2025-11-04 15:34:54.564",
    "created_by": "0634",
    "updated_at": "2025-11-04 15:34:54.564",
    "updated_by": "0634",
    "remark": "费用支出项目数据导入"
  }
]';

--  3. 插入缺失的项目信息
INSERT INTO bid_project (id,
                         name,
                         code,
                         type,
                         source,
                         department_code,
                         department_name,
                         biz_rep_no,
                         biz_rep_name,
                         created_at,
                         created_by,
                         updated_at,
                         updated_by,
                         remark)
SELECT jt.id,
       jt.name,
       jt.code,
       jt.type,
       jt.source,
       jt.department_code,
       jt.department_name,
       jt.biz_rep_no,
       jt.biz_rep_name,
       STR_TO_DATE(jt.created_at, '%Y-%m-%d %H:%i:%s.%f') AS created_at,
       jt.created_by,
       STR_TO_DATE(jt.updated_at, '%Y-%m-%d %H:%i:%s.%f') AS updated_at,
       jt.updated_by,
       jt.remark
FROM JSON_TABLE(
             @json_data,
             '$[*]' COLUMNS (
                 id VARCHAR(32) PATH '$.id',
                 name VARCHAR(64) PATH '$.name',
                 code VARCHAR(64) PATH '$.code',
                 type VARCHAR(10) PATH '$.type',
                 source VARCHAR(10) PATH '$.source',
                 department_code VARCHAR(64) PATH '$.department_code',
                 department_name VARCHAR(64) PATH '$.department_name',
                 biz_rep_no VARCHAR(8) PATH '$.biz_rep_no',
                 biz_rep_name VARCHAR(16) PATH '$.biz_rep_name',
                 created_at DATETIME(3) PATH '$.created_at',
                 created_by VARCHAR(32) PATH '$.created_by',
                 updated_at DATETIME(3) PATH '$.updated_at',
                 updated_by VARCHAR(32) PATH '$.updated_by',
                 remark LONGTEXT PATH '$.remark'
                 )
     ) AS jt
ON DUPLICATE KEY UPDATE name            = VALUES(name),
                        code            = VALUES(code),
                        type            = VALUES(type),
                        source          = VALUES(source),
                        department_code = VALUES(department_code),
                        department_name = VALUES(department_name),
                        biz_rep_no      = VALUES(biz_rep_no),
                        biz_rep_name    = VALUES(biz_rep_name),
                        created_at      = VALUES(created_at),
                        created_by      = VALUES(created_by),
                        updated_at      = VALUES(updated_at),
                        updated_by      = VALUES(updated_by),
                        remark          = VALUES(remark);

-- 4. 费用支出表，更新项目关联关系
update bid_expense e
    inner join JSON_TABLE(@json_data,
                          '$[*]' COLUMNS (
                              business_id VARCHAR(32) PATH '$.expense_business_id',
                              id VARCHAR(32) PATH '$.id'
                              )
               ) AS jt on e.business_id = jt.business_id
set e.project_id = jt.id,
    e.updated_at = current_timestamp(3),
    e.updated_by = '0634'
where e.project_id is null;