package repository

const (
	createUserQuery                = `INSERT INTO users(uuid, username, first_name, last_name, password, email, email_verified, enabled, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	findUserByIDQuery              = `SELECT id, username, first_name, last_name, password, email, created_at, updated_at, managing_id, enabled, email_verified, uuid FROM users WHERE id = $1`
	updateUserQuery                = `UPDATE users SET username = $1, email = $2, updated_at = $3 WHERE user_id = $4`
	destroyUserQuery               = `DELETE FROM users WHERE user_id = $1`
	findUserByUsernameOrEmailQuery = `SELECT id, username, first_name, last_name, password, email, created_at, updated_at, managing_id, enabled, email_verified, uuid FROM users WHERE username = $1 OR email = $2`
)
