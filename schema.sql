CREATE TABLE `users` (
                         `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
                         `name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT 'User name',
                         `age` INT NOT NULL DEFAULT 0 COMMENT 'User age',
                         `semester` INT NOT NULL DEFAULT 0 COMMENT 'Current semester',
                         `cgpa` DECIMAL(3,2) NOT NULL DEFAULT 0.00 COMMENT 'Cumulative Grade Point Average',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Users table';
