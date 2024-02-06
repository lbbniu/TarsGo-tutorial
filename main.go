package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
	"github.com/golang-jwt/jwt/v5"

	"github.com/lbbniu/TarsGo-tutorial/internal/servant"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

// --config/config.conf
func main() {
	log := rogger.GetLogger("order")
	defer rogger.FlushLogger()
	cfg := tars.GetServerConfig()
	imp := servant.NewOrderCtx()
	app := new(order.OrderManagement)
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".OrderObj")

	tars.UseServerFilterMiddleware(func(next tars.ServerFilter) tars.ServerFilter {
		return func(ctx context.Context, d tars.Dispatch, f interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
			if req.Context == nil {
				return tars.Errorf(http.StatusForbidden, "missing context")
			}
			// tokenString := strings.TrimPrefix(req.Context["authorization"], "Bearer ")
			tokenString := req.Context["authorization"]

			token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte("84810c08ab8805c376b1bf614b8446c6"), nil
			})
			claims, ok := token.Claims.(*jwt.RegisteredClaims)
			if !ok || !token.Valid {
				return tars.Errorf(http.StatusUnauthorized, err.Error())
			}
			log.Infof("user_id: %s", claims.ID)
			return next(ctx, d, f, req, resp, withContext)
		}
	})

	tars.Run()
}
