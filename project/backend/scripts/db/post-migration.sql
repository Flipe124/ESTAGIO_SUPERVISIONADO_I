-- INSERTS TO TYPES
INSERT INTO `types` (`code`, `name`)
SELECT 0, 'entrada'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `types` WHERE `code` = 0 AND `name` = 'entrada');

INSERT INTO `types` (`code`, `name`)
SELECT 1, 'saida'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `types` WHERE `code` = 1 AND `name` = 'saida');

-- INSERTS TO STATUSES
INSERT INTO `statuses` (`code`, `name`)
SELECT 0, 'pendente'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 0 AND `name` = 'pendente');

INSERT INTO `statuses` (`code`, `name`)
SELECT 1, 'concluido'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 1 AND `name` = 'concluido');

-- INSERTS TO CATEGORIES
INSERT INTO `categories` (`user_id`, `name`)
SELECT 0, 'Educação'
WHERE NOT EXISTS (SELECT `user_id`, `name` FROM `categories` WHERE `user_id` = 0 AND `name` = 'Educação');

INSERT INTO `categories` (`user_id`, `name`)
SELECT 0, 'Lazer'
WHERE NOT EXISTS (SELECT `user_id`, `name` FROM `categories` WHERE `user_id` = 0 AND `name` = 'Lazer');

INSERT INTO `categories` (`user_id`, `name`)
SELECT 0, 'Alimentação'
WHERE NOT EXISTS (SELECT `user_id`, `name` FROM `categories` WHERE `user_id` = 0 AND `name` = 'Alimentação');

INSERT INTO `categories` (`user_id`, `name`)
SELECT 0, 'Saúde'
WHERE NOT EXISTS (SELECT `user_id`, `name` FROM `categories` WHERE `user_id` = 0 AND `name` = 'Saúde');

INSERT INTO `categories` (`user_id`, `name`)
SELECT 0, 'Outros'
WHERE NOT EXISTS (SELECT `user_id`, `name` FROM `categories` WHERE `user_id` = 0 AND `name` = 'Outros');
