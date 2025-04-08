package controllers

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"piroux.dev/yoping/api/pkg/apps/main/domain/models"
	"piroux.dev/yoping/api/pkg/apps/main/domain/services"
)

/*
TODO: Follow Cloud-Events json-schema format:
https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/formats/cloudevents.json
*/

type ControllerPing struct {
	ServicePing *services.PingService
}

type PingExRequest struct {
	PhoneNumberFrom string `path:"phoneFrom"`
	PhoneNumberTo   string `path:"phoneTo"`
}

type PingExResponse struct {
	Ping *models.Ping
}

func (ctr *ControllerPing) PingEx(ctx context.Context, req *PingExRequest) (rsp *PingExResponse, err error) {

	pingExReq := services.PingExRequest{
		PingData: models.PingData{
			PhoneNumberFrom: req.PhoneNumberFrom,
			PhoneNumberTo:   req.PhoneNumberTo,
		},
	}

	pingExDomainRsp, err := ctr.ServicePing.PingEx(ctx, pingExReq)
	if err != nil {
		return nil, err
	}

	return &PingExResponse{
		Ping: pingExDomainRsp.Ping,
	}, nil
}

type PingInRequest struct {
	ContentType string `header:"Content-Type" enum:"application/cloudevents+json"`
	Body        event.Event
}

// github.com/cloudevents/sdk-go/v2@v2.15.2/event
type PingInRequestData struct {
	PhoneNumberFrom string `json:"phone_from"`
	PhoneNumberTo   string `json:"phone_to"`
}

type PingInResponse struct {
	Ping *models.Ping
}

func (ctr *ControllerPing) PingIn(ctx context.Context, req *PingInRequest) (rsp *PingInResponse, err error) {

	reqData := PingInRequestData{}

	err = req.Body.DataAs(&reqData)
	if err != nil {
		return nil, err
	}

	pingInReq := services.PingInRequest{
		PingData: models.PingData{
			PhoneNumberFrom: reqData.PhoneNumberFrom,
			PhoneNumberTo:   reqData.PhoneNumberTo,
		},
	}

	pingInDomainRsp, err := ctr.ServicePing.PingIn(ctx, pingInReq)
	if err != nil {
		return nil, err
	}

	return &PingInResponse{
		Ping: pingInDomainRsp.Ping,
	}, nil
}
