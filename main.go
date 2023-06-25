package main

import (
	"github.com/TarsCloud/TarsGo/tars"

	"github.com/lbbniu/TarsGo-tutorial/internal/servant"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

// --config/config.conf
func main() {
	cfg := tars.GetServerConfig()
	imp := servant.NewOrderCtx()
	app := new(order.OrderManagement)
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".OrderObj")
	tars.Run()
}
