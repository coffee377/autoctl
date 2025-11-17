--  FIX 中标信息关联项目错误
update bid_info i
set i.project_id = '36c5f9215c689d7fa12c021b7b929041',
    i.id         = '202508131418000512239'
where i.id = '202505091957000135778';

SET @json_data = '[
  {
    "project_name": "定远县卫健委基本公共卫生服务信息系统升级改造项目",
    "project_code": "GS2024-021",
    "total_amount": 58200,
    "contract_signed": true,
    "contract_no": "A261-12",
    "contract_sign_date": "2025-01-10"
  },
  {
    "project_name": "肇源县中医院临床路径系统项目(第4次)",
    "project_code": "HLJGCYC16990100Z20241958968",
    "total_amount": 30000,
    "contract_signed": true,
    "contract_no": "A267-16",
    "contract_sign_date": "2025-01-22"
  },
  {
    "project_name": "合肥市社会救助综合服务平台运维",
    "project_code": "2025BFFFD00036",
    "total_amount": 750200,
    "contract_signed": true,
    "contract_no": "A263-17",
    "contract_sign_date": "2025-01-23"
  },
  {
    "project_name": "桐城市基层卫生医疗机构信息化系统运维服务项目",
    "project_code": "/",
    "total_amount": 273040,
    "contract_signed": true,
    "contract_no": "A261-10",
    "contract_sign_date": "2025-01-25"
  },
  {
    "project_name": "经开区检验检查互认和公卫慢阻肺功能升级项目",
    "project_code": "CG-AQ-W2025-025",
    "total_amount": 59000,
    "contract_signed": true,
    "contract_no": "A262-25",
    "contract_sign_date": "2025-01-27"
  },
  {
    "project_name": "智慧共享中药房基层HIS接口服务项目",
    "project_code": "/",
    "total_amount": 60000,
    "contract_signed": true,
    "contract_no": "A261-23",
    "contract_sign_date": "2025-02-13"
  },
  {
    "project_name": "中国石油加油站管理系统运行维护（2025年）项目智慧销售信息系统一线运维服务采购（二次）",
    "project_code": "ZY25-Z227-FW010",
    "total_amount": 4480000,
    "contract_signed": true,
    "contract_no": "A264-34",
    "contract_sign_date": "2025-04-21"
  },
  {
    "project_name": "安庆市大观区检验检查结果互认和基本公共卫生服务慢阻肺功能升级项目",
    "project_code": "HYDL-2025021",
    "total_amount": 60000,
    "contract_signed": true,
    "contract_no": "A263-3",
    "contract_sign_date": "2025-03-07"
  },
  {
    "project_name": "马鞍山市智慧养老服务信息平台项目",
    "project_code": "MASCG-0-J-F-2025-0105",
    "total_amount": 460000,
    "contract_signed": true,
    "contract_no": "A263-6",
    "contract_sign_date": "2025-03-07"
  },
  {
    "project_name": "中国石油天然气零售系统运行维护（2025年）项目天然气零售系统及配套的一线运维服务采购",
    "project_code": "ZY25-Z227-FW035",
    "total_amount": 928000,
    "contract_signed": true,
    "contract_no": "A265-7",
    "contract_sign_date": "2025-04-28"
  },
  {
    "project_name": "昆仑大模型产品接收测试服务采购",
    "project_code": "/",
    "total_amount": 760000,
    "contract_signed": true,
    "contract_no": "A264-9",
    "contract_sign_date": "2025-04-10"
  },
  {
    "project_name": "安徽省某单位产品及服务030040甄选项目",
    "project_code": "ZYJC-ZX-202503-0135",
    "total_amount": 10258000,
    "contract_signed": true,
    "contract_no": "A264-49",
    "contract_sign_date": "2025-04-22"
  },
  {
    "project_name": "肇源县总医院HIS系统和重症监护信息管理系统升级项目",
    "project_code": "DTZB-2025-0005",
    "total_amount": 240000,
    "contract_signed": true,
    "contract_no": "A266-9",
    "contract_sign_date": "2025-04-03"
  },
  {
    "project_name": "六安市医保数据专区建设及数据开发利用采购项目第3包",
    "project_code": "FS34150120250066号",
    "total_amount": 688000,
    "contract_signed": true,
    "contract_no": "A264-14",
    "contract_sign_date": "2025-04-09"
  },
  {
    "project_name": "南阳市第一人民医院医养信息系统采购项目",
    "project_code": "卧龙政采磋商-2025-12",
    "total_amount": 1186000,
    "contract_signed": true,
    "contract_no": "A269-25",
    "contract_sign_date": "2025-06-05"
  },
  {
    "project_name": "肇源县中医院国家传染病智能监测接口开发服务项目",
    "project_code": "ZC【2022】001",
    "total_amount": 50000,
    "contract_signed": true,
    "contract_no": "A267-18",
    "contract_sign_date": "2025-05-19"
  },
  {
    "project_name": "云梦泽智慧平台项目测试服务",
    "project_code": "ZY25-Z227-FW046",
    "total_amount": 4100000,
    "contract_signed": true,
    "contract_no": "A266-27",
    "contract_sign_date": "2025-05-06"
  },
  {
    "project_name": "长丰县卫生健康委信息化运维项目",
    "project_code": "2025ACCFN00070",
    "total_amount": 1507000,
    "contract_signed": true,
    "contract_no": "A265-6",
    "contract_sign_date": "2025-04-14"
  },
  {
    "project_name": "肇源县智慧医院移动端便民服务平台",
    "project_code": "源财采备[2025]00080号",
    "total_amount": 189000,
    "contract_signed": true,
    "contract_no": "A264-36",
    "contract_sign_date": "2025-04-17"
  },
  {
    "project_name": "医院管理系统升级",
    "project_code": "[230624]FDGJ[DY]20250001",
    "total_amount": 585000,
    "contract_signed": true,
    "contract_no": "A265-28",
    "contract_sign_date": "2025-04-30"
  },
  {
    "project_name": "远程心电市平台运维服务项目",
    "project_code": "2025WGXYW02",
    "total_amount": 61000,
    "contract_signed": true,
    "contract_no": "A269-27",
    "contract_sign_date": "2025-06-27"
  },
  {
    "project_name": "中国电信股份有限公司淮北分公司淮北市民政局智慧养老助餐监管一体化平台OCR 技术服务项目",
    "project_code": "ZYT-20250428-035",
    "total_amount": 10000,
    "contract_signed": true,
    "contract_no": "A268-46",
    "contract_sign_date": "2025-06-10"
  },
  {
    "project_name": "固镇县卫健委基层卫生院国家传染病预警监测前置软件部署项目",
    "project_code": "/",
    "total_amount": 151800,
    "contract_signed": true,
    "contract_no": "A267-19",
    "contract_sign_date": "2025-05-19"
  },
  {
    "project_name": "安徽省民政厅信息系统运维项目（二次）",
    "project_code": "AZZB-2025-FW040902B",
    "total_amount": 213900,
    "contract_signed": true,
    "contract_no": "A267-24",
    "contract_sign_date": "2025-05-21"
  },
  {
    "project_name": "濉溪县妇幼保健院信息化建设采购项目",
    "project_code": "SXQT-25014",
    "total_amount": 4370000,
    "contract_signed": true,
    "contract_no": "A267-13",
    "contract_sign_date": "2025-05-20"
  },
  {
    "project_name": "亳州市谯城区民政局衔接并轨应用信息化采购项目",
    "project_code": "/",
    "total_amount": 255300,
    "contract_signed": true,
    "contract_no": "A268-16",
    "contract_sign_date": "2025-06-09"
  },
  {
    "project_name": "青海省社会组织法人单位信息资源库第三方运维保障服务项目",
    "project_code": "/",
    "total_amount": 68000,
    "contract_signed": true,
    "contract_no": "A269-11",
    "contract_sign_date": "2025-06-10"
  },
  {
    "project_name": "绵阳市居民家庭经济状况核对系统和绵阳市低收入人口动态监测预警平台系统信息化建设服务项目",
    "project_code": "SCGX2025031",
    "total_amount": 278000,
    "contract_signed": true,
    "contract_no": "A269-7",
    "contract_sign_date": "2025-06-16"
  },
  {
    "project_name": "博爱成都信息化建设服务采购项目",
    "project_code": "0701-2541SC110436",
    "total_amount": 849000,
    "contract_signed": true,
    "contract_no": "A269-19",
    "contract_sign_date": "2025-06-20"
  },
  {
    "project_name": "“天府救助通”四川省低收入人口动态监测信息平台电子签章设备采购项目",
    "project_code": "/",
    "total_amount": 50400,
    "contract_signed": true,
    "contract_no": "A268-38",
    "contract_sign_date": "2025-06-19"
  },
  {
    "project_name": "安居区民政局“安逸救”一体化综合服务平台项目",
    "project_code": "SCXY-2025-101号",
    "total_amount": 242500,
    "contract_signed": true,
    "contract_no": "A269-43",
    "contract_sign_date": "2025-06-25"
  },
  {
    "project_name": "四川省低收入人口动态监测信息平台刚性支出困难家庭认定子系统拓展建设项目",
    "project_code": "/",
    "total_amount": 210000,
    "contract_signed": true,
    "contract_no": "A270-18",
    "contract_sign_date": "2025-07-16"
  },
  {
    "project_name": "桐城市医共体检验互联互通服务项目",
    "project_code": "/",
    "total_amount": 246000,
    "contract_signed": true,
    "contract_no": "A269-37",
    "contract_sign_date": "2025-06-28"
  },
  {
    "project_name": "望江县开展全省“两项政策”衔接并轨改革试点工作之民政低收入平台系统升级项目",
    "project_code": "WJXM-2025046",
    "total_amount": 150000,
    "contract_signed": true,
    "contract_no": "A270-12",
    "contract_sign_date": "2025-06-30"
  },
  {
    "project_name": "徽州区“两项政策”衔接并轨采购项目",
    "project_code": "ahzx20250605",
    "total_amount": 155000,
    "contract_signed": true,
    "contract_no": "A269-33",
    "contract_sign_date": "2025-06-30"
  },
  {
    "project_name": "安徽省社会救助一件事联办应用开发项目第1包",
    "project_code": "2025BFAFZ01289",
    "total_amount": 1120000,
    "contract_signed": true,
    "contract_no": "A270-13",
    "contract_sign_date": "2025-07-11"
  },
  {
    "project_name": "阜阳市防止返贫帮扶政策和农村低收入人口常态化帮扶政策衔接并轨监测模块建设项目",
    "project_code": "AHZCZB(2025)032",
    "total_amount": 77500,
    "contract_signed": true,
    "contract_no": "A272-18",
    "contract_sign_date": "2025-07-16"
  },
  {
    "project_name": "桐城市区域合理用药管理系统项目",
    "project_code": "H1QT25Z990154",
    "total_amount": 729000,
    "contract_signed": true,
    "contract_no": "A270-9",
    "contract_sign_date": "2025-07-16"
  },
  {
    "project_name": "修文县六广镇中心卫生院慢病管理系统与贵州省公卫系统数据对接采购项目",
    "project_code": "[XWX-LGZZXWSY-2025-002]",
    "total_amount": 48000,
    "contract_signed": true,
    "contract_no": "A271-35",
    "contract_sign_date": "2025-07-22"
  },
  {
    "project_name": "清镇市第一人民医院（医共体）基本公共卫生系统接口服务（二次发布）",
    "project_code": "QZSYY2025-7-14-2",
    "total_amount": 24000,
    "contract_signed": true,
    "contract_no": "A273-6",
    "contract_sign_date": "2025-08-19"
  },
  {
    "project_name": "铜陵市郊区衔接并轨应用信息化项目",
    "project_code": "GDZB-2025-044 ",
    "total_amount": 96000,
    "contract_signed": true,
    "contract_no": "A271-20",
    "contract_sign_date": "2025-07-29"
  },
  {
    "project_name": "巢湖市基层卫生院国家传染病预警监测前置软件部署建设项目",
    "project_code": "2025JRAZB0025",
    "total_amount": 108800,
    "contract_signed": true,
    "contract_no": "A271-18",
    "contract_sign_date": "2025-08-11"
  },
  {
    "project_name": "宿州市低收入人口动态监测信息平台升级改造项目",
    "project_code": "AHWL20250721",
    "total_amount": 83000,
    "contract_signed": true,
    "contract_no": "A272-26",
    "contract_sign_date": "2025-08-26"
  },
  {
    "project_name": "清镇市中医医院（医共体）基本公共卫生系统提供的接口服务",
    "project_code": "QZSZYY2025-07-23-01",
    "total_amount": 23000,
    "contract_signed": true,
    "contract_no": "A273-31",
    "contract_sign_date": "2025-09-18"
  },
  {
    "project_name": "攀枝花共同富裕试验区低收入群体监测帮扶服务平台（数字驾驶舱）二期建设采购项目",
    "project_code": "SCJC-FZ(P)-202507032",
    "total_amount": 798000,
    "contract_signed": true,
    "contract_no": "A272-27",
    "contract_sign_date": "2025-08-28"
  },
  {
    "project_name": "三明市养老服务综合管理平台升级改造项目（包一）",
    "project_code": "闽宏正[2025]明招010号",
    "total_amount": 216330,
    "contract_signed": true,
    "contract_no": "A272-5",
    "contract_sign_date": "2025-08-19"
  },
  {
    "project_name": "中电信数智科技有限公司安徽分公司阜阳市县域基层医疗信息化建设项目（包一）",
    "project_code": "AHDX250438",
    "total_amount": 2900000,
    "contract_signed": true,
    "contract_no": "A273-21",
    "contract_sign_date": "2025-09-18"
  },
  {
    "project_name": "肥东县2020年卫生健康信息化智慧医疗建设项目运维项目",
    "project_code": "2025ADDFN00120 ",
    "total_amount": 2566000,
    "contract_signed": true,
    "contract_no": "A271-23",
    "contract_sign_date": "2025-08-13"
  },
  {
    "project_name": "中国电信四川公司2025年四川省民政厅业务系统综合运维项目",
    "project_code": "HX-202500060-013",
    "total_amount": 1287900,
    "contract_signed": true,
    "contract_no": "A269-2",
    "contract_sign_date": null
  },
  {
    "project_name": "泗县妇幼保健计划生育服务中心（泗县妇幼保健院）智慧医院建设项目",
    "project_code": "EP-SXQT2025039",
    "total_amount": 389800,
    "contract_signed": true,
    "contract_no": "A272-24",
    "contract_sign_date": "2025-08-27"
  },
  {
    "project_name": "合肥市医保数据专区建设和场景应用项目",
    "project_code": "2025BFFFZ01975",
    "total_amount": 2050000,
    "contract_signed": true,
    "contract_no": "A272-34",
    "contract_sign_date": "2025-08-29"
  },
  {
    "project_name": "淮南市低收入人口动态监测信息平台升级改造工作",
    "project_code": "HNBJ-2025CG108",
    "total_amount": 84199,
    "contract_signed": true,
    "contract_no": "A273-34",
    "contract_sign_date": "2025-09-19"
  },
  {
    "project_name": "马鞍山市社会救助管理中心“核对系统”、“大数据系统”运维项目",
    "project_code": "AHSYZTB-2025-365",
    "total_amount": 194000,
    "contract_signed": true,
    "contract_no": "A274-32",
    "contract_sign_date": "2025-09-23"
  },
  {
    "project_name": "芜湖市社会救助大数据信息系统升级优化项目",
    "project_code": "JSDC[Q]20250909",
    "total_amount": 258860,
    "contract_signed": true,
    "contract_no": "A274-30",
    "contract_sign_date": "2025-09-29"
  },
  {
    "project_name": "芜湖市居民家庭经济状况核对信息系统运维服务项目",
    "project_code": "/",
    "total_amount": 55000,
    "contract_signed": true,
    "contract_no": "A274-31",
    "contract_sign_date": "2025-10-11"
  }
]';

-- 中标项目签约合同信息
update bid_info i
    left join bid_project p
    on i.project_id = p.id
    inner join JSON_TABLE(
            @json_data,
            '$[*]' COLUMNS (
                project_name VARCHAR(125) PATH '$.project_name',
                project_code VARCHAR(64) PATH '$.project_code',
                total_amount decimal(16, 4) PATH '$.total_amount',
                contract_signed TINYINT(1) PATH '$.contract_signed',
                contract_no VARCHAR(64) PATH '$.contract_no',
                contract_sign_date DATETIME PATH '$.contract_sign_date'
                )
               ) AS jt on p.name = jt.project_name and p.code = jt.project_code
        and i.software_amount + i.hardware_amount + i.operation_amount = jt.total_amount
set i.contract_signed    = jt.contract_signed,
    i.contract_no        = jt.contract_no,
    i.contract_sign_date = jt.contract_sign_date,
    i.updated_at         = now(3),
    i.updated_by         = '0634',
    i.remark             = '销售合同信息导入'
where i.contract_no is null
  and i.contract_signed = false;