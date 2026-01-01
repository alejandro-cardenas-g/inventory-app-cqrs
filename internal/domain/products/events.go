package products

import (
	"time"
)

type ProductCreated struct {
	ProductID  int64
	OccurredAt time.Time
}
