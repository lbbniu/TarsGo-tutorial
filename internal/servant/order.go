package servant

import (
	"context"
	"net/http"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/lbbniu/TarsGo-tutorial/order"
)

type Order struct {
}

var orders = make(map[string]order.Order)

var _ order.OrderManagementServantWithContext = (*Order)(nil)

func (o Order) GetOrder(tarsCtx context.Context, orderId string) (ret order.Order, err error) {
	ord, exists := orders[orderId]
	if exists {
		return ord, nil
	}

	return ord, tars.Errorf(http.StatusNotFound, "Order does not exist. : ", orderId)
}
