package repo_ping

import (
	"time"

	"github.com/objectbox/objectbox-go/objectbox"
)

type PingEntity struct {
	Id          uint64 `objectbox:"id"`
	PhoneFrom   string
	PhoneTo     string
	TimeCreated time.Time
}

type RepoBox struct {
	ob      *objectbox.ObjectBox
	pingBox *objectbox.Box
}

/*

func NewRepoBox(ob *objectbox.ObjectBox) *RepoBox {
	pingBox := ob.Box(PingEntity{}) // Pass the entity type to create a box
	return &RepoBox{
		ob:      ob,
		pingBox: pingBox,
	}
}


var _ domainAdapters.PingRespository = &RepoBox{}


func (rb *RepoBox) Create(ping *models.Ping) (*models.Ping, error) {
	entity := &PingEntity{
		PhoneFrom:   ping.PhoneNumbers.From,
		PhoneTo:     ping.PhoneNumbers.To,
		TimeCreated: ping.TimeCreated,
	}

	id, err := rb.pingBox.Put(entity)
	if err != nil {
		return nil, err
	}

	entity.Id = id
	return &models.Ping{
		PhoneNumbers: models.PhoneNumberPair{
			From: entity.PhoneFrom,
			To:   entity.PhoneTo,
		},
		TimeCreated: entity.TimeCreated,
	}, nil
}

func (rb *RepoBox) Update(ping *models.Ping) (*models.Ping, error) {
	query := rb.pingBox.Query(
		objectbox.Equals("PhoneFrom", ping.PhoneNumbers.From),
		objectbox.Equals("PhoneTo", ping.PhoneNumbers.To),
	)
	defer query.Close()

	entities, err := query.Find()
	if err != nil || len(entities) == 0 {
		return nil, repos.ErrDataNotFound
	}

	entity := entities[0].(*PingEntity)
	entity.TimeCreated = ping.TimeCreated

	_, err = rb.pingBox.Put(entity)
	if err != nil {
		return nil, err
	}

	return &models.Ping{
		PhoneNumbers: models.PhoneNumberPair{
			From: entity.PhoneFrom,
			To:   entity.PhoneTo,
		},
		TimeCreated: entity.TimeCreated,
	}, nil
}

func (rb *RepoBox) Delete(ping *models.Ping) error {
	query := rb.pingBox.Query(
		objectbox.Equals("PhoneFrom", ping.PhoneNumbers.From),
		objectbox.Equals("PhoneTo", ping.PhoneNumbers.To),
	)
	defer query.Close()

	entities, err := query.Find()
	if err != nil || len(entities) == 0 {
		return repos.ErrDataNotFound
	}

	entity := entities[0].(*PingEntity)
	err = rb.pingBox.Remove(entity.Id)
	if err != nil {
		return err
	}

	return nil
}

func (rb *RepoBox) GetOne(from, to string) (*models.Ping, error) {
	query := rb.pingBox.Query(
		objectbox.Equals("PhoneFrom", from),
		objectbox.Equals("PhoneTo", to),
	)
	defer query.Close()

	entities, err := query.Find()
	if err != nil || len(entities) == 0 {
		return nil, repos.ErrDataNotFound
	}

	entity := entities[0].(*PingEntity)
	return &models.Ping{
		PhoneNumbers: models.PhoneNumberPair{
			From: entity.PhoneFrom,
			To:   entity.PhoneTo,
		},
		TimeCreated: entity.TimeCreated,
	}, nil
}

func (rb *RepoBox) GetAll() ([]*models.Ping, error) {
	entities, err := rb.pingBox.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*models.Ping, 0, len(entities))
	for _, e := range entities {
		entity := e.(*PingEntity)
		result = append(result, &models.Ping{
			PhoneNumbers: models.PhoneNumberPair{
				From: entity.PhoneFrom,
				To:   entity.PhoneTo,
			},
			TimeCreated: entity.TimeCreated,
		})
	}

	return result, nil
}


*/
