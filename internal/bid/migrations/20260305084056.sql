-- Create "sys_task_log" table
CREATE TABLE `sys_task_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `biz_type` enum('1','2','0') NOT NULL DEFAULT "0" COMMENT "任务流转业务类型 1:商品采购 2:项目投标 0:其他",
  `biz_id` varchar(32) NOT NULL COMMENT "业务标识",
  `assign_seq` int unsigned NOT NULL COMMENT "同一业务标识下的指派序号（从1开始递增）",
  `assign_time` datetime NULL COMMENT "任务指派时间",
  `handler_no` varchar(8) NULL COMMENT "受理人工号",
  `start_time` datetime NULL COMMENT "任务开始时间",
  `end_time` datetime NULL COMMENT "任务结束时间（任务完成/终止的时间）",
  `remark` longtext NULL COMMENT "备注（如重新指派原因、任务终止说明等）",
  `created_at` datetime(3) NOT NULL COMMENT "创建时间",
  `created_by` varchar(32) NULL COMMENT "创建人",
  `updated_at` datetime(3) NOT NULL COMMENT "更新时间",
  `updated_by` varchar(32) NULL COMMENT "更新人",
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uk_bti_as` (`biz_type`, `biz_id`, `assign_seq`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "任务流转记录";
