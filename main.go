package main

import (
	"context"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"

	"github.com/lbbniu/TarsGo-tutorial/internal/servant"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

// --config/config.conf
func main() {
	cfg := tars.GetServerConfig()
	imp := servant.NewOrderCtx()
	app := new(order.OrderManagement)
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".OrderObj")
	tars.UseServerFilterMiddleware(func(next tars.ServerFilter) tars.ServerFilter {
		return func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
			return next(ctx, d, f, req, resp, withContext)
		}
	})
	tars.Run()
}
