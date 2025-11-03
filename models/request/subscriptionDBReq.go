package request

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateSubscriptionReq struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type UpdateSubscriptionReq struct {
	SubscriptionID int       `json:"subscription_id"`
	ServiceName    string    `json:"service_name,omitempty"`
	Price          int       `json:"price,omitempty"`
	StartDate      time.Time `json:"start_date,omitempty"`
	EndDate        time.Time `json:"end_date,omitempty"`
}

type DeleteSubscriptionrReq struct {
	SubscriptionID int `json:"subscription_id"`
}

type GetFullPriceByPeriodReq struct {
	UserID    uuid.UUID `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
