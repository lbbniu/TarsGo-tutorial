package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/current"

	"github.com/lbbniu/TarsGo-tutorial/order"
)

func main() {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)
	tars.RegisterClientFilter(func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
		// TODO: 增加请求前业务逻辑
		// 这里可以增加：日志记录、权限认证、限流、链路追踪等
		s := time.Now()
		err = invoke(ctx, msg, timeout)
		// TODO: 增加请求后业务逻辑
		log.Printf("ServantName: %s, FuncName: %s, req: %v, resp: %v, latency: %s\n",
			msg.Req.SServantName, msg.Req.SFuncName, msg.Req, msg.Resp, time.Now().Sub(s))
		return err
	})

	tars.RegisterPreClientFilter(func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
		// TODO: 增加请求前业务逻辑
		return err // 此处返回err，客户端会继续处理，只会记录对应的错误日志
	})

	tars.RegisterPreClientFilter(func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
		// TODO: 增加请求后业务逻辑
		return err // 此处返回err，客户端会继续处理，只会记录对应的错误日志
	})

	tars.UseClientFilterMiddleware(func(next tars.ClientFilter) tars.ClientFilter {
		return func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
			// TODO: 增加请求前业务逻辑
			// 这里可以增加：日志记录、权限认证、限流、链路追踪等
			s := time.Now()
			err = next(ctx, msg, invoke, timeout)
			// TODO: 增加请求后业务逻辑
			log.Printf("ServantName: %s, FuncName: %s, latency: %s\n",
				msg.Req.SServantName, msg.Req.SFuncName, time.Now().Sub(s))
			return err
		}
	})

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
