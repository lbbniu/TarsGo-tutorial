package main

import (
	"context"
	"time"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
	"github.com/golang-jwt/jwt/v5"

	"github.com/lbbniu/TarsGo-tutorial/order"
)

type JwtAuthentication struct {
	key []byte
}

func NewJwtAuthentication(key []byte) *JwtAuthentication {
	return &JwtAuthentication{key: key}
}

func (j *JwtAuthentication) Build() tars.ClientFilterMiddleware {
	return func(next tars.ClientFilter) tars.ClientFilter {
		return func(ctx context.Context, msg *tars.Message, invoke tars.Invoke, timeout time.Duration) (err error) {
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				ID:        "lbbniu",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			})
			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString(j.key)
			if err != nil {
				return err
			}
			if msg.Req.Context == nil {
				msg.Req.Context = make(map[string]string)
			}
			//msg.Req.Context["authorization"] = "Bearer " + tokenString
			msg.Req.Context["authorization"] = tokenString
			return next(ctx, msg, invoke, timeout)
		}
	}
}

func main() {
	log := rogger.GetLogger("order")
	defer rogger.FlushLogger()
	jwtAuth := NewJwtAuthentication([]byte("84810c08ab8805c376b1bf614b8446c6"))
	tars.UseClientFilterMiddleware(jwtAuth.Build())

	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order, err := client.GetOrderWithContext(ctx, "1")
	if err != nil {
		panic(err)
	}
	log.Infof("GetOrderWithContext Response: %+v", order)
}
