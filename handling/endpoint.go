package handling

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	cargo "github.com/days85/shipping_cargo"
	location "github.com/days85/shipping_location"
	voyage "github.com/days85/shipping_voyage"
)

type registerIncidentRequest struct {
	ID             cargo.TrackingID
	Location       location.UNLocode
	Voyage         voyage.Number
	EventType      cargo.HandlingEventType
	CompletionTime time.Time
}

type registerIncidentResponse struct {
	Err error `json:"error,omitempty"`
}

func (r registerIncidentResponse) error() error { return r.Err }

func makeRegisterIncidentEndpoint(hs Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerIncidentRequest)
		err := hs.RegisterHandlingEvent(req.CompletionTime, req.ID, req.Voyage, req.Location, req.EventType)
		return registerIncidentResponse{Err: err}, nil
	}
}
