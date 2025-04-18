// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package gen_sql_dst

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPing = `-- name: CreatePing :one

INSERT
INTO pings (
  phone_to, phone_from, time_created
) VALUES (
  $1, $2, $3
)
RETURNING phone_to, phone_from, time_created
`

type CreatePingParams struct {
	PhoneTo     string
	PhoneFrom   string
	TimeCreated pgtype.Timestamp
}

// Pings ---------------------------------------------------------------------------------
func (q *Queries) CreatePing(ctx context.Context, arg CreatePingParams) (Ping, error) {
	row := q.db.QueryRow(ctx, createPing, arg.PhoneTo, arg.PhoneFrom, arg.TimeCreated)
	var i Ping
	err := row.Scan(&i.PhoneTo, &i.PhoneFrom, &i.TimeCreated)
	return i, err
}

const createUser = `-- name: CreateUser :one

INSERT
INTO users (
  id, name_full, phone
) VALUES (
  $1, $2, $3
)
RETURNING id, name_full, phone, bio
`

type CreateUserParams struct {
	ID       pgtype.UUID
	NameFull string
	Phone    string
}

// Users ---------------------------------------------------------------------------------
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.NameFull, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.NameFull,
		&i.Phone,
		&i.Bio,
	)
	return i, err
}

const deletePing = `-- name: DeletePing :exec
DELETE
FROM pings
WHERE phone_to = $1 AND phone_from = $2
`

type DeletePingParams struct {
	PhoneTo   string
	PhoneFrom string
}

func (q *Queries) DeletePing(ctx context.Context, arg DeletePingParams) error {
	_, err := q.db.Exec(ctx, deletePing, arg.PhoneTo, arg.PhoneFrom)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getContacts = `-- name: GetContacts :many
SELECT DISTINCT users.id, users.name_full, users.phone, users.bio
FROM users
JOIN pings pings_from on pings_from.phone_from = users.phone
JOIN pings pings_to on pings_to.phone_to = users.phone
WHERE pings_from.phone_from = $1 OR pings_from.phone_to = $1
ORDER BY name_full
`

func (q *Queries) GetContacts(ctx context.Context, phoneFrom string) ([]User, error) {
	rows, err := q.db.Query(ctx, getContacts, phoneFrom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.NameFull,
			&i.Phone,
			&i.Bio,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getContactsBis = `-- name: GetContactsBis :many
(
  SELECT users.id, users.name_full, users.phone, users.bio
  FROM users
  JOIN pings on pings.phone_to = $1
  ORDER BY name_full
)
UNION
(
  SELECT users.id, users.name_full, users.phone, users.bio
  FROM users
  JOIN pings on pings.phone_from = $1
  ORDER BY name_full
)
`

func (q *Queries) GetContactsBis(ctx context.Context, phoneTo string) ([]User, error) {
	rows, err := q.db.Query(ctx, getContactsBis, phoneTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.NameFull,
			&i.Phone,
			&i.Bio,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPing = `-- name: GetPing :one
SELECT phone_to, phone_from, time_created
FROM pings
WHERE phone_to = $1 AND phone_from = $2
LIMIT 1
`

type GetPingParams struct {
	PhoneTo   string
	PhoneFrom string
}

func (q *Queries) GetPing(ctx context.Context, arg GetPingParams) (Ping, error) {
	row := q.db.QueryRow(ctx, getPing, arg.PhoneTo, arg.PhoneFrom)
	var i Ping
	err := row.Scan(&i.PhoneTo, &i.PhoneFrom, &i.TimeCreated)
	return i, err
}

const getPings = `-- name: GetPings :many
SELECT phone_to, phone_from, time_created
FROM pings
ORDER BY phone_to, phone_from, time_created
`

func (q *Queries) GetPings(ctx context.Context) ([]Ping, error) {
	rows, err := q.db.Query(ctx, getPings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ping
	for rows.Next() {
		var i Ping
		if err := rows.Scan(&i.PhoneTo, &i.PhoneFrom, &i.TimeCreated); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name_full, phone, bio
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.NameFull,
		&i.Phone,
		&i.Bio,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name_full, phone, bio
FROM users
ORDER BY name_full
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.NameFull,
			&i.Phone,
			&i.Bio,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
