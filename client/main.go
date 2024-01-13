package main

import (
	"context"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"log"
	"log/slog"
	"strconv"
	"time"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

func orderClientFilter(next tars.ClientFilter) tars.ClientFilter {
	return func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
		if msg.Req.Context == nil {
			msg.Req.Context = make(map[string]string)
		}
		var (
			s  string
			ok bool
		)
		if s, ok = msg.Req.Context["time"]; !ok {
			s = "inter" + strconv.FormatInt(time.Now().UnixNano(), 10)
			msg.Req.Context["time"] = s
		}
		log.Printf("call timestamp: %s", s)
		// Invoking the remote method
		err = next(ctx, msg, invoke, timeout)
		return err
	}
}

func main() {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)
	tars.UseClientFilterMiddleware(orderClientFilter)
	ctxTimeCall(client)
	ctxCall(client)
}

func ctxTimeCall(client *order.OrderManagement) {
	ctx := current.ContextWithTarsCurrent(context.Background())
	ctx = current.ContextWithClientCurrent(ctx)
	mcontext := map[string]string{
		"time": "raw" + strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	order, err := client.GetOrderWithContext(ctx, "1", mcontext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ctx: %+v\n", order)
}

// 有context.Content调用
func ctxCall(client *order.OrderManagement) {
	ctx := current.ContextWithTarsCurrent(context.Background())
	ctx = current.ContextWithClientCurrent(ctx)
	client.TarsSetTimeout(10000)
	mcontext := map[string]string{"client-key1": "client-value1", "client-key2": "client-value2"}
	status := map[string]string{"client-key1": "client-value1", "client-key2": "client-value2"}
	current.SetClientTimeout(ctx, 10000)
	order, err := client.GetOrderWithContext(ctx, "1", mcontext, status)
	if err != nil {
		panic(err)
	}
	slog.InfoContext(ctx, "serverContext", "context", mcontext)
	slog.InfoContext(ctx, "serverStatus", "status", status)

	fmt.Printf("ctx: %+v\n", order)
}
