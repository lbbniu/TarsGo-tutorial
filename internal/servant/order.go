package servant

import (
	"context"
	"net/http"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/lbbniu/TarsGo-tutorial/order"
)

var orders = make(map[string]order.Order)

type OrderCtx struct {
}

// 有context.Context实现服务端
var _ order.OrderManagementServantWithContext = (*OrderCtx)(nil)

func NewOrderCtx() *OrderCtx {
	o := &OrderCtx{}
	o.init()
	return o
}

func (o *OrderCtx) init() {
	orders["1"] = order.Order{
		Id:          "1",
		Price:       100,
		Items:       []string{"iPhone 11", "MacBook Pro"},
		Description: "MacBook Pro",
		Destination: "Beijing",
	}
}

func (o *OrderCtx) GetOrder(tarsCtx context.Context, orderId string) (ret order.Order, err error) {
	ord, exists := orders[orderId]
	if exists {
		return ord, nil
	}

	return ord, tars.Errorf(http.StatusNotFound, "Order does not exist. : ", orderId)
}
