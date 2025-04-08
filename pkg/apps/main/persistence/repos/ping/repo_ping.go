package repo_ping

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domainAdapters "piroux.dev/yoping/api/pkg/apps/main/domain/adapters"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
	"piroux.dev/yoping/api/pkg/apps/main/persistence/repos"
	"piroux.dev/yoping/api/pkg/apps/main/persistence/storage/gensql/gen_sql_dst"
	// _ "gitlab.com/cznic/sqlite"
)

type RepoDB struct {
	poolDB  *pgxpool.Pool
	queries gen_sql_dst.Queries
}

func NewRepoDB(dbURL string) *RepoDB {

	poolDB, err := pgxpool.New(context.Background(), dbURL)
	// db, err := sql.Open("sqlite", ":memory:") // TODO
	if err != nil {
		log.Fatalf("failed to create DB connection: %s", err.Error())
	}

	return &RepoDB{
		poolDB:  poolDB,
		queries: *gen_sql_dst.New(poolDB),
	}
}

var _ domainAdapters.PingRespository = &RepoDB{} // remove after impl

func (rp *RepoDB) Create(ping *models.Ping) (*models.Ping, error) {
	var err error
	ctx := context.TODO()

	tx, err := rp.poolDB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	params := gen_sql_dst.CreatePingParams{
		PhoneTo:   ping.PhoneNumbers.To,
		PhoneFrom: ping.PhoneNumbers.From,
	}
	err = params.TimeCreated.Scan(ping.TimeCreated)
	if err != nil {
		return nil, err
	}

	resultDbPing, err := rp.queries.CreatePing(ctx, params)
	if err != nil {
		return nil, err
	}

	resultModelPing := &models.Ping{
		PhoneNumbers: models.PhoneNumberPair{
			From: resultDbPing.PhoneFrom,
			To:   resultDbPing.PhoneTo,
		},
		TimeCreated: resultDbPing.TimeCreated.Time,
	}

	return resultModelPing, err
}

func (rp *RepoDB) Update(ping *models.Ping) (*models.Ping, error) {
	return nil, errors.New("TODO: unimplemented")
}

func (rp *RepoDB) Delete(ping *models.Ping) error {
	var err error
	ctx := context.TODO()

	tx, err := rp.poolDB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	params := gen_sql_dst.DeletePingParams{
		PhoneTo:   ping.PhoneNumbers.To,
		PhoneFrom: ping.PhoneNumbers.From,
	}

	err = rp.queries.DeletePing(ctx, params)
	if err == pgx.ErrNoRows {
		return repos.ErrDataNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
func (rp *RepoDB) GetOne(from, to string) (*models.Ping, error) {
	var err error
	ctx := context.TODO()

	params := gen_sql_dst.GetPingParams{
		PhoneTo:   to,
		PhoneFrom: from,
	}

	resultDbPing, err := rp.queries.GetPing(ctx, params)
	if err == pgx.ErrNoRows {
		return nil, repos.ErrDataNotFound
	}
	if err != nil {
		return nil, err
	}

	resultModelPing := &models.Ping{
		PhoneNumbers: models.PhoneNumberPair{
			From: resultDbPing.PhoneFrom,
			To:   resultDbPing.PhoneFrom, // BUG
		},
		TimeCreated: resultDbPing.TimeCreated.Time,
	}

	return resultModelPing, err
}

func (rp *RepoDB) GetAll() ([]*models.Ping, error) {
	var err error
	ctx := context.TODO()

	resultDbPings, err := rp.queries.GetPings(ctx)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	resultModelPings := make([]*models.Ping, 0, len(resultDbPings))
	for _, resultDbPing := range resultDbPings {
		resultModelPings = append(resultModelPings, &models.Ping{
			PhoneNumbers: models.PhoneNumberPair{
				From: resultDbPing.PhoneFrom,
				To:   resultDbPing.PhoneTo,
			},
			TimeCreated: resultDbPing.TimeCreated.Time,
		})
	}

	return resultModelPings, err
}
