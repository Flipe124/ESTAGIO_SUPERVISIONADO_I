package user

var postCreate = `
	INSERT INTO accounts (user_id, name, balance, created_at, updated_at)
	VALUES (?, 'Carteira', 0, NOW(), NOW())
`
