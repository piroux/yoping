// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package gen_sql_dst

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Ping struct {
	PhoneTo     string
	PhoneFrom   string
	TimeCreated pgtype.Timestamp
}

type User struct {
	ID       pgtype.UUID
	NameFull string
	Phone    string
	Bio      pgtype.Text
}
