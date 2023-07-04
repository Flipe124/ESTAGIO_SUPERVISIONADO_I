-- INSERTS TO STATUS
INSERT INTO `statuses` (`code`, `name`)
SELECT 0, 'pendente'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 0 AND `name` = 'pendente');

INSERT INTO `statuses` (`code`, `name`)
SELECT 1, 'concluido'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 1 AND `name` = 'concluido');
