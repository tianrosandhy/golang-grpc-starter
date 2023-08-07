package handler

import (
	"context"
	appointmentPB "grpc/protob/appointment"
)

type AppointmentService struct{}

func (a *AppointmentService) PostAppointment(ctx context.Context, req *appointmentPB.AppointmentData) (*appointmentPB.AppointmentData, error) {
	return &appointmentPB.AppointmentData{}, nil
}

func (a *AppointmentService) GetSingleAppointment(ctx context.Context, req *appointmentPB.GetAppointmentRequest) (*appointmentPB.AppointmentData, error) {
	return &appointmentPB.AppointmentData{}, nil
}

func (a *AppointmentService) GetAllAppointment(ctx context.Context, req *appointmentPB.GetAppointmentRequest) (*appointmentPB.AppointmentDataCollection, error) {
	return &appointmentPB.AppointmentDataCollection{}, nil
}
