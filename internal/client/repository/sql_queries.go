package repository

const (
	createClientQuery              = `INSERT INTO clients(client_id, client_secret, user_id, managing_id, enabled, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	findClientByTextAndUserIDQuery = `SELECT id, client_id FROM clients WHERE client_id = $1 AND user_id = $2`
	findClientByIDAndSecret        = `SELECT id, client_id, client_secret FROM clients WHERE client_id = $1 AND client_secret = $2`
)

const (
	findManagingByTextQuery = `SELECT id, managing_id FROM managing WHERE managing_id = $1 AND enabled = TRUE`
)
