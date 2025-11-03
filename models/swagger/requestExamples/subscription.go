package requestExamples

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateSubscriptionReqExample struct {
	ServiceName string    `json:"service_name" example:"Netflix"`
	Price       int       `json:"price" example:"100"`
	UserID      string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426655440000"`
	StartDate   time.Time `json:"start_date"  example:"2023-01-01T00:00:00Z"`
	EndDate     time.Time `json:"end_date" example:"2023-12-31T00:00:00Z"`
}

type UpdateSubscriptionReqExample struct {
	SubscriptionID uint      `json:"subscription_id" example:"1"`
	ServiceName    string    `json:"service_name,omitempty" example:"Netflix"`
	Price          int       `json:"price,omitempty" example:"100"`
	StartDate      time.Time `json:"start_date,omitempty" example:"2023-01-01T00:00:00Z"`
	EndDate        time.Time `json:"end_date,omitempty" example:"2023-12-31T00:00:00Z"`
}

type DeleteSubscriptionReqExample struct {
	ID uint `json:"subscription_id" example:"1"`
}

type GetFullPriceByPeriodReqExample struct {
	UserID    uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426655440000"`
	StartDate time.Time `json:"start_date" example:"2023-01-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" example:"2023-12-31T00:00:00Z"`
}
