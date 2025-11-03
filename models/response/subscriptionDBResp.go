package response

import (
	"time"

	"github.com/gofrs/uuid"
)

type GetSudscription struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
type GetFullPriceByPeriodResp struct {
	UserID    uuid.UUID `json:"user_id"`
	TotalCost int       `json:"total_cost"`
	Count     int       `json:"count"`
}
