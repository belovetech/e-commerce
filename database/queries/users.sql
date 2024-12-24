-- name: CreateUser :one
INSERT INTO users (email, password, role)
VALUES ($1, $2, $3)
RETURNING id, email, role;

-- name: GetUserByEmail :one
SELECT id, email, password,  role FROM users WHERE email = $1;


-- name: GetAdmins :many
SELECT id, email, role FROM users WHERE role = 'admin';
