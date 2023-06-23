package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"graph-gateway/protos/catalog"
)

// AppointmentCharge is the resolver for the appointment_charge field.
func (r *appointmentChargeResolver) AppointmentCharge(ctx context.Context, obj *catalog.AppointmentCharge) (float64, error) {
	return float64(obj.AppointmentCharge), nil
}

// Price is the resolver for the price field.
func (r *clientPriceResolver) Price(ctx context.Context, obj *catalog.ClientPrice) (float64, error) {
	return float64(obj.Price), nil
}

// Price is the resolver for the price field.
func (r *serviceResolver) Price(ctx context.Context, obj *catalog.Service) (float64, error) {
	return float64(obj.Price), nil
}

// Price is the resolver for the price field.
func (r *addServiceRequestResolver) Price(ctx context.Context, obj *catalog.AddServiceRequest, data float64) error {
	obj.Price = float32(data)
	return nil
}

// Price is the resolver for the price field.
func (r *editServiceRequestResolver) Price(ctx context.Context, obj *catalog.EditServiceRequest, data float64) error {
	obj.Price = float32(data)
	return nil
}

// Price is the resolver for the price field.
func (r *setClientPriceRequestResolver) Price(ctx context.Context, obj *catalog.SetClientPriceRequest, data float64) error {
	obj.Price = float32(data)
	return nil
}

// AppointmentCharge returns AppointmentChargeResolver implementation.
func (r *Resolver) AppointmentCharge() AppointmentChargeResolver {
	return &appointmentChargeResolver{r}
}

// ClientPrice returns ClientPriceResolver implementation.
func (r *Resolver) ClientPrice() ClientPriceResolver { return &clientPriceResolver{r} }

// Service returns ServiceResolver implementation.
func (r *Resolver) Service() ServiceResolver { return &serviceResolver{r} }

// AddServiceRequest returns AddServiceRequestResolver implementation.
func (r *Resolver) AddServiceRequest() AddServiceRequestResolver {
	return &addServiceRequestResolver{r}
}

// EditServiceRequest returns EditServiceRequestResolver implementation.
func (r *Resolver) EditServiceRequest() EditServiceRequestResolver {
	return &editServiceRequestResolver{r}
}

// SetClientPriceRequest returns SetClientPriceRequestResolver implementation.
func (r *Resolver) SetClientPriceRequest() SetClientPriceRequestResolver {
	return &setClientPriceRequestResolver{r}
}

type appointmentChargeResolver struct{ *Resolver }
type clientPriceResolver struct{ *Resolver }
type serviceResolver struct{ *Resolver }
type addServiceRequestResolver struct{ *Resolver }
type editServiceRequestResolver struct{ *Resolver }
type setClientPriceRequestResolver struct{ *Resolver }
