package main

import (
	"context"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"log"
	"time"

	"github.com/lbbniu/TarsGo-tutorial/internal/servant"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

// --config/config.conf
func main() {
	cfg := tars.GetServerConfig()
	imp := servant.NewOrderCtx()
	app := new(order.OrderManagement)
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".OrderObj")

	// 注册服务端过滤器
	tars.RegisterServerFilter(func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
		// TODO: 增加服务端处理前业务逻辑
		// 这里可以增加：日志记录、权限认证、限流、链路追踪等
		s := time.Now()
		err = d(ctx, f, req, resp, withContext)
		// TODO: 增加服务端处理后业务逻辑
		log.Printf("ServantName: %s, FuncName: %s, req: %v, resp: %v, latency: %s\n",
			req.SServantName, req.SFuncName, req, resp, time.Now().Sub(s))
		return err
	})

	tars.RegisterPreServerFilter(func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
		// TODO: 增加服务端处理前业务逻辑
		return err // 此处返回err，服务端会继续处理，只会记录对应的错误日志
	})

	tars.RegisterPostServerFilter(func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
		// TODO: 增加服务端处理前业务逻辑
		return err // 此处返回err，服务端会继续处理，只会记录对应的错误日志
	})

	tars.UseServerFilterMiddleware(func(next tars.ServerFilter) tars.ServerFilter {
		return func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
			// TODO: 增加服务端处理前业务逻辑
			// 这里可以增加：日志记录、权限认证、限流、链路追踪等
			s := time.Now()
			err = next(ctx, d, f, req, resp, withContext)
			// TODO: 增加服务端处理后业务逻辑
			log.Printf("ServantName: %s, FuncName: %s, latency: %s\n",
				req.SServantName, req.SFuncName, time.Now().Sub(s))
			return err
		}
	})

	tars.Run()
}
