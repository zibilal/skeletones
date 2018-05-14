package apitest

import "github.com/zibilal/skeletones/uuid"

type OrderDTO struct {
	OrderId uuid.ID `json:"order_id"`
}
