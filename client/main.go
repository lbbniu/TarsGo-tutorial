package main

import (
	"context"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/current"

	"github.com/lbbniu/TarsGo-tutorial/order"
)

func main() {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@ssl -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)

	noCtxCall(client)
	ctxCall(client)
	oneWayCall(client)
	modHashCall(client)
	consistentHashCall(client)
}

// 无context.Content调用
func noCtxCall(client *order.OrderManagement) {
	order, err := client.GetOrder("1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("noctx: %+v\n", order)
}

// 有context.Content调用
func ctxCall(client *order.OrderManagement) {
	order, err := client.GetOrderWithContext(context.Background(), "1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ctx: %+v\n", order)
}

// 单向调用，无返回值，目前函数前面会有返回值，后续tars2go中会去掉
func oneWayCall(client *order.OrderManagement) {
	_, err := client.GetOrderOneWayWithContext(context.Background(), "1")
	if err != nil {
		panic(err)
	}
	fmt.Println("oneway")
}

// 取模调用
func modHashCall(client *order.OrderManagement) {
	ctx := current.ContextWithClientCurrent(context.Background())
	var hashCode uint32 = 1
	current.SetClientHash(ctx, int(tars.ModHash), hashCode)
	order, err := client.GetOrderWithContext(context.Background(), "1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ModHash: %+v\n", order)
}

// 一致性哈希调用
func consistentHashCall(client *order.OrderManagement) {
	ctx := current.ContextWithClientCurrent(context.Background())
	var hashCode uint32 = 1
	current.SetClientHash(ctx, int(tars.ConsistentHash), hashCode)
	order, err := client.GetOrderWithContext(context.Background(), "1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ConsistentHash: %+v\n", order)
}
