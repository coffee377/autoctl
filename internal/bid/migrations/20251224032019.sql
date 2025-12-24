-- Modify "bid_info" table
ALTER TABLE `bid_info` MODIFY COLUMN `id` varchar(32) NOT NULL COMMENT "投标信息 ID", MODIFY COLUMN `bid_status` enum('RP','RO','RS','RF','DP','B','W','L','F','A','0') NOT NULL DEFAULT "0" COMMENT "投标状态 RP:待报名 RO:报名中 RS:报名成功 RF:报名失败 DP:标书编制中 B:投标中 W:已中标 L:未中标 F:流标 A:弃标 0:-";
