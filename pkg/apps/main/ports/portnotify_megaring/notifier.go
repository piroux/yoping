package portnotify_megaring

import (
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"piroux.dev/yoping/api/pkg/apps/main/domain/adapters"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
)

/*
import "piroux.dev/yoping/api/pkg/apps/main/domain/models"

type PingNotifier interface {
	Notify(ping models.Ping) error
}
*/

type NotifierMegaring struct {
}

var _ adapters.PingNotifier = &NotifierMegaring{}

func (ntf *NotifierMegaring) Notify(ping models.Ping) error {
	pingDebug := (&spew.ConfigState{}).Sdump(ping)
	slog.Info("notify on Megaring", slog.String("ping", pingDebug))
	return nil
}
