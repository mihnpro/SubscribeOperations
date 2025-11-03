package database

import (
	"time"

	"github.com/gofrs/uuid"
)

type SubscriptionDB struct {
	SubscriptionID int       `db:"subscription_id" json:"subscription_id,omitempty"`
	ServiceName    string    `db:"service_name" json:"service_name"`
	Price          int       `db:"price" json:"price"`
	UserID         uuid.UUID `db:"user_id" json:"user_id"`
	StartDate      time.Time `db:"start_date" json:"start_date"`
	EndDate        time.Time    `db:"end_date" json:"end_date"`
}

type FullPriceByPeriodDbResp struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	TotalCost int       `db:"total_cost" json:"total_cost"`
	Count     int       `db:"count" json:"count"`
}
