-- name: SaveUser :exec
INSERT INTO users (
    id, email, password, username
) VALUES (
    $1, $2, $3, $4
);

-- name: GetByEmail :one 
SELECT *
FROM users u 
WHERE u.email = $1;

-- name: GetUserByID :one
SELECT *
FROM users u
WHERE u.id = $1;