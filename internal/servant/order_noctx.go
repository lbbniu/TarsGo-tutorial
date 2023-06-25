package servant

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/lbbniu/TarsGo-tutorial/order"
	"net/http"
)

type Order struct {
}

// 无context.Context实现服务端
var _ order.OrderManagementServant = (*Order)(nil)

// GetOrder returns an order by id
func (o *Order) GetOrder(orderId string) (ret order.Order, err error) {
	ord, exists := orders[orderId]
	if exists {
		return ord, nil
	}

	return ord, tars.Errorf(http.StatusNotFound, "Order does not exist. : ", orderId)
}
