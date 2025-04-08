package controllers

import (
	"context"

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

/*
// CloudEvent JSON Payload: github.com/cloudevents/sdk-go/v2@v2.15.2/event
type PingInRequest struct {
	ContentType string `header:"Content-Type" enum:"application/cloudevents+json" required:"true"`
	Body        event.Event
}
*/

// Simple JSON Payload
type PingInRequest struct {
	ContentType string `header:"Content-Type" enum:"application/json" required:"true"`
	Body        PingInRequestData
}

// NOTE: Sourcing this data from the input payload is effectively redundant with the URL path parameters,
// but this is a project to test stuff.
type PingInRequestData struct {
	PhoneNumberFrom string `json:"phone_from" required:"true"`
	PhoneNumberTo   string `json:"phone_to" required:"true"`
}

type PingInResponse struct {
	Ping *models.Ping
}

func (ctr *ControllerPing) PingIn(ctx context.Context, req *PingInRequest) (rsp *PingInResponse, err error) {

	//reqData := PingInRequestData{}

	/*
		err = req.Body.DataAs(&reqData)
		if err != nil {
			return nil, err
		}
	*/

	pingInReq := services.PingInRequest{
		PingData: models.PingData{
			PhoneNumberFrom: req.Body.PhoneNumberFrom,
			PhoneNumberTo:   req.Body.PhoneNumberTo,
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

/*
	// 3. Create a new CloudEvent
	event := cloudevents.NewEvent() // Uses the CloudEvents 1.0 spec by default

	// 4. Set required CloudEvents attributes
	event.SetID(uuid.New().String()) // Unique ID for the event
	event.SetSource("example/myapp/instance-1") // Identifies the origin of the event
	event.SetType("com.example.myevent.v1")    // Describes the event type (often in reverse domain notation)

	// 5. Set optional CloudEvents attributes
	event.SetTime(time.Now())          // Timestamp of the event
	event.SetSubject("ResourceUpdate") // Describes the subject of the event

*/
