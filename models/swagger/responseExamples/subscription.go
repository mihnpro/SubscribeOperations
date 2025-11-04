package responseExamples

import (
	"time"

	"github.com/gofrs/uuid"
)

type GetSudscriptionExample struct {
	ServiceName string    `json:"service_name" example:"Netflix"`
	Price       int       `json:"price" example:"100"`
	UserID      string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426655440000"`
	StartDate   time.Time `json:"start_date" example:"2023-01-01T00:00:00Z"`
	EndDate     time.Time `json:"end_date" example:"2023-12-31T00:00:00Z"`
}

type GetFullPriceByPeriodExample struct {
	UserID    uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426655440000"`
	TotalCost int       `json:"total_cost" example:"1000"`
	Count     int       `json:"count" example:"2"`
}

type SubCreatedExample struct {
	Message string `json:"message" example:"Subscription created"`
}

type EmptySubDateExample struct {
	Error string `json:"error" example:"start_date and end_date are required"`
}

type InvalidSubDateExample struct {
	Error string `json:"error" example:"start_date must be before end_date"`
}

type SomeServerErrorsDuriingCreatingSubExample struct {
	Error string `json:"error" example:"some server errors during creating subscription"`
}

type SubDeletedExample struct {
	Message string `json:"message" example:"Subscription deleted"`
}

type SomeServerErrorsDuriingDeletingSubExample struct {
	Error string `json:"error" example:"some server errors during deleting subscription"`
}

type SubUpdatedExample struct{
	Message string `json:"message" example:"Subscription updated"`
}

type SubNotFoundExample struct{
	Error string `json:"error" example:"Subscription not found"`
}

type SomeServerErrorsDuriingUpdatingSubExample struct {
	Error string `json:"error" example:"some server errors during updating subscription"`
}

type EmptyDateOrInvalidExample struct {
	Error string `json:"error" example:"start_date and end_date are required or start_date must be before end_date"`
}

type SomeErrorGetingAllSubsExample struct {
	Error string `json:"error" example:"some server errors during getting all subscriptions"`
}

type InvalidSubIDExample struct {
	Error string `json:"error" example:"Invalid subscription ID"`
}

type EmptySubOrInvalidIDExample struct {
	Error string `json:"error" example:"Subscription ID is required or Invalid subscription ID"`
}

type SomeServerErrorsDuringGettingSubExample struct {
	Error string `json:"error" example:"some server errors during getting subscription"`
}

type SomeServerErrorsDuringGetingFullSubPriceByPeriod struct {
	Error string `json:"error" example:"some server errors during getting full price by period"`
}

