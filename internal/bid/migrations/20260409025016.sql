-- Modify "bid_ca_certificate" table
ALTER TABLE `bid_ca_certificate` DROP INDEX `idx_code`, ADD UNIQUE INDEX `uk_code` (`code`);
