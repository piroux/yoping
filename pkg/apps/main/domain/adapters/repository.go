package adapters

import "piroux.dev/yoping/api/pkg/apps/main/domain/models"

type UserRespository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) error

	GetOne(userID string) (*models.User, error)
	GetAll() ([]*models.User, error)

	GetContacts(userID string) ([]*models.User, error)
}

type PingRespository interface {
	Create(ping *models.Ping) (*models.Ping, error)
	Update(ping *models.Ping) (*models.Ping, error)
	Delete(ping *models.Ping) error

	GetOne(from, to string) (*models.Ping, error)
	GetAll() ([]*models.Ping, error)

	// GetForUserTo(to string) (*models.Ping, error)
}
