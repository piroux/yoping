package adapters

import "piroux.dev/yoping/api/pkg/apps/main/domain/models"

type PingNotifier interface {
	Notify(ping models.Ping) error
}
