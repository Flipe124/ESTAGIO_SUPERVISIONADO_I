-- INSERTS TO TYPE
INSERT INTO `types` (`code`, `name`)
SELECT 0, 'entrada'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `types` WHERE `code` = 0 AND `name` = 'entrada');

INSERT INTO `types` (`code`, `name`)
SELECT 1, 'saida'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `types` WHERE `code` = 1 AND `name` = 'saida');

-- INSERTS TO STATUS
INSERT INTO `statuses` (`code`, `name`)
SELECT 0, 'pendente'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 0 AND `name` = 'pendente');

INSERT INTO `statuses` (`code`, `name`)
SELECT 1, 'concluido'
WHERE NOT EXISTS (SELECT `code`, `name` FROM `statuses` WHERE `code` = 1 AND `name` = 'concluido');
