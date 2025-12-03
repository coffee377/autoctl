-- Modify "bid_info" table
ALTER TABLE `bid_info` DROP INDEX `fk_pid_03`, ADD UNIQUE INDEX `project_id` (`project_id`);
