package user

// SystemUserID saves the id of system user.
const SystemUserID = 0

var postCreate = `
	INSERT INTO accounts (user_id, name, balance, created_at, updated_at)
	VALUES (?, 'Carteira', 0, NOW(), NOW())
`
