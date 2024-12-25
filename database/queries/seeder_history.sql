-- name: GetSeederByName :one
SELECT seeder_name FROM seeding_history WHERE seeder_name = $1;


-- name: CreateSeederHistory :exec
INSERT INTO seeding_history (seeder_name) VALUES ($1);

