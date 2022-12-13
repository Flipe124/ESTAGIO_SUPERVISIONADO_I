-- Adminer 4.8.1 MySQL 8.0.30 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

INSERT INTO `account` (`id`, `name`, `balance`, `created_at`, `updated_at`) VALUES
(1,	'nubank',	1000,	'2022-11-21 10:48:17',	'2022-11-21 10:48:17');

INSERT INTO `category` (`id`, `name`, `icon`, `type`, `created_at`, `updated_at`) VALUES
(1,	'casa',	'/',	'EXPENSE',	'2022-11-21 10:37:40',	'2022-11-21 10:37:40'),
(2,	'serviço',	'/',	'EXPENSE',	'2022-11-21 10:39:24',	'2022-11-21 10:39:24'),
(3,	'mercado',	'/',	'EXPENSE',	'2022-11-21 10:40:07',	'2022-11-21 10:40:07'),
(4,	'Salário',	'/',	'REVENUE',	'2022-11-27 20:06:06',	'2022-11-27 20:06:06');

DROP TABLE IF EXISTS `finance`;
CREATE TABLE `finance` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `category_id` int NOT NULL,
  `value` double NOT NULL,
  `type` varchar(45) NOT NULL,
  `status` varchar(45) NOT NULL,
  `description` varchar(50) NOT NULL,
  `date` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`account_id`,`category_id`),
  KEY `fk_finance_category_idx` (`category_id`),
  KEY `fk_finance_account1_idx` (`account_id`),
  CONSTRAINT `fk_finance_account1` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`),
  CONSTRAINT `fk_finance_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `finance` (`id`, `account_id`, `category_id`, `value`, `type`, `status`, `description`, `date`, `created_at`, `updated_at`) VALUES
(1,	1,	1,	100,	'EXPENSE',	'PAID',	'Aluguel',	'2022-11-17 10:49:05',	'2022-11-21 10:49:55',	'2022-11-27 19:46:47'),
(3,	1,	2,	200,	'EXPENSE',	'NOT_PAID',	'Gasolina',	'2022-11-09 11:41:38',	'2022-11-21 11:42:25',	'2022-11-27 19:46:47'),
(5,	1,	3,	1000,	'EXPENSE',	'PAID',	'Supermercado',	'2022-11-10 16:32:27',	'2022-11-22 16:33:13',	'2022-11-27 19:53:19'),
(6,	1,	4,	1500,	'REVENUE',	'PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:36'),
(7,	1,	2,	5000,	'EXPENSE',	'NOT_PAID',	'mercado',	'2022-12-05 00:00:00',	'2022-12-08 00:34:57',	'2022-12-08 00:34:57'),
(8,	1,	4,	1000,	'REVENUE',	'NOT_PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:20'),
(9,	1,	4,	1000,	'REVENUE',	'NOT_PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:20'),
(10,	1,	4,	1000,	'REVENUE',	'NOT_PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:20'),
(12,	1,	4,	1500,	'REVENUE',	'PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:36'),
(13,	1,	4,	1000,	'REVENUE',	'NOT_PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:20'),
(14,	1,	4,	1000,	'REVENUE',	'NOT_PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:20'),
(16,	1,	4,	1500,	'REVENUE',	'PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:36'),
(17,	1,	4,	1500,	'REVENUE',	'PAID',	'salário',	'2022-11-27 20:06:43',	'2022-11-27 20:08:02',	'2022-12-12 04:45:36');

SET NAMES utf8mb4;

INSERT INTO `table` (`id`, `column_1`, `column_2`, `column_3`) VALUES
(1,	1,	'1970-01-01',	'yesss'),
(2,	1,	'1970-01-01',	'yesss'),
(3,	1,	'1970-01-01',	'yesss'),
(4,	139,	'2021-12-01',	NULL),
(5,	12,	'2022-11-13',	NULL),
(6,	NULL,	NULL,	NULL),
(7,	14,	'2022-11-14',	'rgdgrdg'),
(8,	NULL,	NULL,	'hello world'),
(9,	25,	'1999-01-01',	'no'),
(10,	25,	'1999-01-01',	'no'),
(11,	1,	'1970-01-01',	'yesss');

-- 2022-12-13 21:26:36
