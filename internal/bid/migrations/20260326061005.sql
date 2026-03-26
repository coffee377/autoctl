-- Modify "bid_ca_certificate" table
ALTER TABLE `bid_ca_certificate` MODIFY COLUMN `code` varchar(24) NOT NULL COMMENT "CA证书编码", MODIFY COLUMN `name` varchar(32) NOT NULL COMMENT "CA证书名称", MODIFY COLUMN `expiry_time` datetime NOT NULL COMMENT "过期时间", MODIFY COLUMN `last_renewal_at` datetime(3) NULL COMMENT "最后续费时间";
-- Modify "bid_member_account" table
ALTER TABLE `bid_member_account` MODIFY COLUMN `username` varchar(32) NOT NULL COMMENT "账号", MODIFY COLUMN `password` varchar(32) NULL COMMENT "密码", MODIFY COLUMN `register_person` varchar(16) NULL COMMENT "注册人员", MODIFY COLUMN `register_mobile` varchar(16) NULL COMMENT " 注册手机号";
