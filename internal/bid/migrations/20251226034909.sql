-- Modify "bid_project" table
ALTER TABLE `bid_project` MODIFY COLUMN `department_name` varchar(64) NOT NULL DEFAULT "" COMMENT "所属部门名称";
