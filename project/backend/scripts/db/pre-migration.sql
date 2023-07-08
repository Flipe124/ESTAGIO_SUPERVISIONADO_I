-- CREATE SYSTEM USER
CREATE TABLE IF NOT EXISTS `users` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(191) NULL DEFAULT 'unnamed',
	`username` VARCHAR(191) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `username` (`username`)
);

SET SQL_MODE='NO_AUTO_VALUE_ON_ZERO';

INSERT INTO `users` (`id`, `username`)
SELECT 0, 'system'
WHERE NOT EXISTS (SELECT `id`, `username` FROM `users` WHERE `id` = 0 AND `username` = 'system');
