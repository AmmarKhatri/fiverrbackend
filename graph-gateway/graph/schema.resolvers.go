package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"
	"graph-gateway/graph/helpers"
	"graph-gateway/protos/comms"
	"graph-gateway/protos/schedule"
)

// UpdateCalendar is the resolver for the UpdateCalendar field.
func (r *mutationResolver) UpdateCalendar(ctx context.Context, input *schedule.UpdateCalendarRequest) (*schedule.UpdateCalendarResponse, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	res, err := c.UpdateCalendar(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// AddAppointment is the resolver for the AddAppointment field.
func (r *mutationResolver) AddAppointment(ctx context.Context, input *schedule.AddAppointmentRequest) (*schedule.Appointment, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	res, err := c.AddAppointment(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// CancelAppointment is the resolver for the CancelAppointment field.
func (r *mutationResolver) CancelAppointment(ctx context.Context, input *schedule.CancelAppointmentRequest) (*string, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	_, err = c.CancelAppointment(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return nil, nil
}

// AcceptedAppointment is the resolver for the AcceptedAppointment field.
func (r *mutationResolver) AcceptedAppointment(ctx context.Context, input *comms.AcceptedAppointmentInput) (*comms.CommsAppointment, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	res, err := c.AcceptAppointment(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// RejectedAppointment is the resolver for the RejectedAppointment field.
func (r *mutationResolver) RejectedAppointment(ctx context.Context, input *comms.RejectedAppointmentInput) (*string, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	_, err = c.RejectAppointment(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return nil, nil
}

// SendMessageFromConsole is the resolver for the SendMessageFromConsole field.
func (r *mutationResolver) SendMessageFromConsole(ctx context.Context, input *comms.ConsoleMessageInput) (*comms.ConsoleMessage, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	res, err := c.SendMessageFromConsole(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// ListProviderDays is the resolver for the ListProviderDays field.
func (r *queryResolver) ListProviderDays(ctx context.Context, input *schedule.ListDaysRequest) (*schedule.ListProviderDaysResponse, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	res, err := c.ListProviderDays(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// ListCustomerDays is the resolver for the ListCustomerDays field.
func (r *queryResolver) ListCustomerDays(ctx context.Context, input *schedule.ListDaysRequest) (*schedule.ListCustomerDaysResponse, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	res, err := c.ListCustomerDays(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// GetDayDetails is the resolver for the GetDayDetails field.
func (r *queryResolver) GetDayDetails(ctx context.Context, input *schedule.DayInput) (*schedule.DayDetails, error) {
	conn, err := helpers.ScheduleConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := schedule.NewSchedulerClient(conn)
	res, err := c.GetDayDetails(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// ProcessOutgoingCallback is the resolver for the ProcessOutgoingCallback field.
func (r *queryResolver) ProcessOutgoingCallback(ctx context.Context, input *comms.OutgoingCallbackRequest) (*string, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	_, err = c.ProcessOutgoingCallback(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return nil, nil
}

// ProcessIncomingSMSCallback is the resolver for the ProcessIncomingSMSCallback field.
func (r *queryResolver) ProcessIncomingSMSCallback(ctx context.Context, input *comms.IncomingSMSCallbackRequest) (*string, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	_, err = c.ProcessIncomingSMSCallback(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return nil, nil
}

// ProcessMultiChannelCallback is the resolver for the ProcessMultiChannelCallback field.
func (r *queryResolver) ProcessMultiChannelCallback(ctx context.Context, input *comms.MultiChannelCallbackRequest) (*string, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	_, err = c.ProcessMultiChannelCallback(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return nil, nil
}

// RequestAppointment is the resolver for the RequestAppointment field.
func (r *queryResolver) RequestAppointment(ctx context.Context, input *comms.CommsAppointmentInput) (*comms.CommsPendingAppointment, error) {
	conn, err := helpers.CommsConnection()
	if err != nil {
		fmt.Println("Error connecting to the client")
		return nil, err
	}
	defer conn.Close()
	c := comms.NewCommunicatorClient(conn)
	res, err := c.RequestAppointment(ctx, input)
	if err != nil {
		fmt.Println("Error getting response and calling function")
		return nil, err
	}
	return res, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }