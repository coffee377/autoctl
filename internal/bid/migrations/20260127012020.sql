-- Modify "bid_info" table
ALTER TABLE `bid_info` MODIFY COLUMN `bid_status` enum('RP','RO','RF','RS','DP','B','E','W','L','F','A','0') NOT NULL DEFAULT "0" COMMENT "投标状态 RP:待报名 RO:报名中 RF:报名失败 RS:报名成功 DP:标书编制中 B:投标中 E:项目入围 W:已中标 L:未中标 F:流标 A:弃标 0:-";
