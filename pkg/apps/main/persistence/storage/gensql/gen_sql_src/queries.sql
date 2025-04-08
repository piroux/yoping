
-- Users ---------------------------------------------------------------------------------

-- name: CreateUser :one
INSERT
INTO users (
  id, name_full, phone
) VALUES (
  $1, $2, $3
)
RETURNING *
;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1
;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY name_full
;

-- name: GetContacts :many
SELECT DISTINCT users.*
FROM users
JOIN pings pings_from on pings_from.phone_from = users.phone
JOIN pings pings_to on pings_to.phone_to = users.phone
WHERE pings_from.phone_from = $1 OR pings_from.phone_to = $1
ORDER BY name_full
;

-- name: GetContactsBis :many
(
  SELECT users.*
  FROM users
  JOIN pings on pings.phone_to = $1
  ORDER BY name_full
)
UNION
(
  SELECT users.*
  FROM users
  JOIN pings on pings.phone_from = $1
  ORDER BY name_full
)
;

-- Pings ---------------------------------------------------------------------------------

-- name: CreatePing :one
INSERT
INTO pings (
  phone_to, phone_from, time_created
) VALUES (
  $1, $2, $3
)
RETURNING *
;

-- name: DeletePing :exec
DELETE
FROM pings
WHERE phone_to = $1 AND phone_from = $2
;

-- name: GetPing :one
SELECT *
FROM pings
WHERE phone_to = $1 AND phone_from = $2
LIMIT 1
;

-- name: GetPings :many
SELECT *
FROM pings
ORDER BY phone_to, phone_from, time_created
;
