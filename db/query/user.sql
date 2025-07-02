-- name: CreateUser :one
   insert into users (
      username,
      mail,
      password_hash
   ) values ( $1,
              $2,
              $3 ) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;