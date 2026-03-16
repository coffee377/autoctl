-- Create "bid_web_site" table
CREATE TABLE `bid_web_site` (
  `id` varchar(32) NOT NULL,
  `province` varchar(8) NOT NULL COMMENT "省份",
  `city` varchar(8) NULL COMMENT "地市",
  `name` varchar(32) NOT NULL COMMENT "网站名称",
  `url` longtext NOT NULL COMMENT "网站地址",
  `active` bool NOT NULL DEFAULT 1 COMMENT "是否启用",
  `remark` longtext NULL COMMENT "备注",
  `created_at` datetime(3) NOT NULL COMMENT "创建时间",
  `created_by` varchar(32) NULL COMMENT "创建人",
  `updated_at` datetime(3) NOT NULL COMMENT "更新时间",
  `updated_by` varchar(32) NULL COMMENT "更新人",
  PRIMARY KEY (`id`),
  INDEX `idx_name` (`name`),
  INDEX `idx_province_city` (`province`, `city`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "网站信息";
-- Create "bid_member_account" table
CREATE TABLE `bid_member_account` (
  `id` varchar(32) NOT NULL,
  `owner_code` varchar(32) NOT NULL COMMENT "归属主体编码",
  `owner_name` varchar(32) NOT NULL COMMENT "归属主体名称",
  `username` varchar(255) NOT NULL COMMENT "账号",
  `password` varchar(255) NULL COMMENT "密码",
  `register_person` varchar(255) NULL COMMENT "注册人员",
  `register_mobile` varchar(255) NULL COMMENT " 注册手机号",
  `primary_ca_id` varchar(32) NOT NULL COMMENT "主证书",
  `account_status` enum('active','inactive','abandoned','suspended') NOT NULL DEFAULT "active" COMMENT "账号状态: active-正常/inactive-未激活/abandoned-废弃/suspended-暂停",
  `abandon_reason` longtext NULL COMMENT "废弃原因",
  `remark` longtext NULL COMMENT "备注",
  `website_id` varchar(32) NOT NULL COMMENT "归属网站",
  `created_at` datetime(3) NOT NULL COMMENT "创建时间",
  `created_by` varchar(32) NULL COMMENT "创建人",
  `updated_at` datetime(3) NOT NULL COMMENT "更新时间",
  `updated_by` varchar(32) NULL COMMENT "更新人",
  PRIMARY KEY (`id`),
  INDEX `idx_website_id` (`website_id`),
  CONSTRAINT `fk_website_id` FOREIGN KEY (`website_id`) REFERENCES `bid_web_site` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "会员账号";
-- Create "bid_ca_certificate" table
CREATE TABLE `bid_ca_certificate` (
  `id` varchar(32) NOT NULL,
  `code` varchar(255) NOT NULL COMMENT "CA证书编码",
  `name` varchar(255) NOT NULL COMMENT "CA证书名称",
  `expiry_time` timestamp NOT NULL COMMENT "过期时间",
  `password` varchar(255) NULL COMMENT "CA证书密码",
  `remark` longtext NULL COMMENT "备注",
  `primary` bool NOT NULL DEFAULT 0 COMMENT "是否为主证书",
  `last_renewal_at` timestamp NULL COMMENT "最后续费时间",
  `created_at` datetime(3) NOT NULL COMMENT "创建时间",
  `created_by` varchar(32) NULL COMMENT "创建人",
  `updated_at` datetime(3) NOT NULL COMMENT "更新时间",
  `updated_by` varchar(32) NULL COMMENT "更新人",
  PRIMARY KEY (`id`),
  INDEX `idx_code` (`code`),
  INDEX `idx_name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "CA 证书";
-- Create "bid_account_ca_rel" table
CREATE TABLE `bid_account_ca_rel` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account_id` varchar(32) NOT NULL COMMENT "会员账号ID",
  `ca_id` varchar(32) NOT NULL COMMENT "CA证书ID",
  `remark` varchar(255) NULL COMMENT "关联备注",
  `bind_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_account_id` (`account_id`),
  INDEX `idx_ca_id` (`ca_id`),
  UNIQUE INDEX `uk_account_id_ca_id` (`account_id`, `ca_id`),
  CONSTRAINT `fk_account_id` FOREIGN KEY (`account_id`) REFERENCES `bid_member_account` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `fk_ca_id` FOREIGN KEY (`ca_id`) REFERENCES `bid_ca_certificate` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "账号证书关联关系";
