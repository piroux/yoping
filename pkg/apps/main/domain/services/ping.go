package services

import (
	"context"
	"fmt"

	"piroux.dev/yoping/api/pkg/apps/main/domain/adapters"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
)

// Application Service
type PingService struct {
	notifier        adapters.PingNotifier
	userRespository adapters.UserRespository
	pingRespository adapters.PingRespository
}

func NewPingService(
	notifier adapters.PingNotifier,
	userRespository adapters.UserRespository,
	pingRespository adapters.PingRespository) *PingService {

	return &PingService{
		notifier:        notifier,
		userRespository: userRespository,
		pingRespository: pingRespository,
	}
}

type PingExRequest struct {
	//PhoneNumberPair models.PhoneNumberPair
	models.PingData
}

type PingExResponse struct {
	Status string
	Ping   *models.Ping
}

func (svc *PingService) PingEx(ctx context.Context, req PingExRequest) (rsp PingExResponse, err error) {
	ping, err := models.NewPing(req.PhoneNumberFrom, req.PhoneNumberTo)
	if err != nil {
		err = fmt.Errorf("failed to construct ping: %w", err)
		return PingExResponse{Status: "failed to build ping"}, err
	}

	_, err = svc.pingRespository.Create(ping) // BUG
	if err != nil {
		return PingExResponse{Status: "failed to store ping"}, err
	}

	err = svc.notifier.Notify(*ping)
	if err != nil {
		return PingExResponse{Status: "failed to notify ping"}, err
	}

	return PingExResponse{
		Status: "ok",
		Ping:   ping,
	}, nil
}

type PingInRequest struct {
	// PhoneNumberPair models.PhoneNumberPair
	models.PingData
}

type PingInResponse struct {
	Status string
	Ping   *models.Ping
}

func (svc *PingService) PingIn(ctx context.Context, req PingInRequest) (rsp PingInResponse, err error) {
	pnPair, err := models.NewPhoneNumberPair(req.PhoneNumberFrom, req.PhoneNumberTo)
	if err != nil {
		return PingInResponse{Status: "failed to validate phone numbers pair"}, err
	}

	rsp.Ping, err = svc.pingRespository.GetOne(
		pnPair.From,
		pnPair.To,
	)
	if err != nil {
		return PingInResponse{Status: "failed to get ping from store"}, err
	}

	err = svc.pingRespository.Delete(rsp.Ping)
	if err != nil {
		return PingInResponse{Status: "failed to delete ping from store"}, err
	}

	rsp.Status = "ok"

	return rsp, nil
}
