-- ################ Queries ################ --

-- name: QueryUserByID :one
SELECT * FROM users
WHERE user_id =  $1 LIMIT 1;

-- name: QueryUserByEmail :one
SELECT * FROM users
WHERE email =  $1 LIMIT 1;

-- ################# Commands ################ --

-- name: CreateUser :exec
INSERT INTO users (user_id, name, email, password_hash, roles, department, enabled, date_created, date_updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
