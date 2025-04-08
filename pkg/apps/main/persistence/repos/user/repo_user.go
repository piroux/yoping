package repo_user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	domainAdapters "piroux.dev/yoping/api/pkg/apps/main/domain/adapters"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
	"piroux.dev/yoping/api/pkg/apps/main/persistence/repos"
	"piroux.dev/yoping/api/pkg/apps/main/persistence/storage/gensql/gen_sql_dst"
)

type RepoDB struct {
	db      pgxpool.Conn
	queries gen_sql_dst.Queries
}

var _ domainAdapters.UserRespository = &RepoDB{} // remove after impl

func (rp *RepoDB) Create(user *models.User) (*models.User, error) {
	var err error
	ctx := context.TODO()

	tx, err := rp.db.BeginTx(ctx, pgx.TxOptions{})
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

	params := gen_sql_dst.CreateUserParams{
		NameFull: user.NameFull,
		Phone:    user.PhoneNumber,
	}
	err = params.ID.Scan(user.Id)
	if err != nil {
		return nil, err
	}

	resultDbUser, err := rp.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	resultModelUser := &models.User{
		Id:          uuid.MustParse(resultDbUser.ID.String()),
		NameFull:    resultDbUser.NameFull,
		PhoneNumber: resultDbUser.Phone,
	}

	return resultModelUser, err
}

func (rp *RepoDB) Update(user *models.User) (*models.User, error) {
	return nil, errors.New("TODO: unimplemented")
}

func (rp *RepoDB) Delete(user *models.User) error {
	var err error
	ctx := context.TODO()

	tx, err := rp.db.BeginTx(ctx, pgx.TxOptions{})
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

	var paramUserId pgtype.UUID
	err = paramUserId.Scan(user.Id)
	if err != nil {
		return err
	}

	err = rp.queries.DeleteUser(ctx, paramUserId)
	if err == pgx.ErrNoRows {
		return repos.ErrDataNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
func (rp *RepoDB) GetOne(userId string) (*models.User, error) {
	var err error
	ctx := context.TODO()

	var paramUserId pgtype.UUID
	err = paramUserId.Scan(userId)
	if err != nil {
		return nil, err
	}

	resultDbUser, err := rp.queries.GetUser(ctx, paramUserId)
	if err == pgx.ErrNoRows {
		return nil, repos.ErrDataNotFound
	}
	if err != nil {
		return nil, err
	}

	resultModelUser := &models.User{
		Id:          uuid.MustParse(resultDbUser.ID.String()),
		NameFull:    resultDbUser.NameFull,
		PhoneNumber: resultDbUser.Phone,
	}

	return resultModelUser, err
}

func (rp *RepoDB) GetAll() ([]*models.User, error) {
	var err error
	ctx := context.TODO()

	resultDbUsers, err := rp.queries.GetUsers(ctx)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	resultModelUsers := make([]*models.User, 0, len(resultDbUsers))
	for _, resultDbUser := range resultDbUsers {
		resultModelUsers = append(resultModelUsers, &models.User{
			Id:          uuid.MustParse(resultDbUser.ID.String()),
			NameFull:    resultDbUser.NameFull,
			PhoneNumber: resultDbUser.Phone,
		})
	}

	return resultModelUsers, err
}

func (rp *RepoDB) GetContacts(userID string) ([]*models.User, error) {
	var err error
	ctx := context.TODO()

	resultDbUsers, err := rp.queries.GetContacts(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	resultModelUsers := make([]*models.User, 0, len(resultDbUsers))
	for _, resultDbUser := range resultDbUsers {
		resultModelUsers = append(resultModelUsers, &models.User{
			Id:          uuid.MustParse(resultDbUser.ID.String()),
			NameFull:    resultDbUser.NameFull,
			PhoneNumber: resultDbUser.Phone,
		})
	}

	return resultModelUsers, err
}
