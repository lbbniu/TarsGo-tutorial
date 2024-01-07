package errors

import (
	"context"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"log"
	"net/http"
	"testing"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/lbbniu/TarsGo-tutorial/order"
)

func Invoke(i bool) error {
	if i {
		return ParamsErr
	} else {
		return BizErr
	}
}

func TestError(t *testing.T) {
	err := Invoke(true)
	if err != nil {
		switch {
		case errors.Is(err, ParamsErr):
			log.Println("params error")
		case errors.Is(err, BizErr):
			log.Println("biz error")
		}
	}
}

func TestRPCError(t *testing.T) {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)

	retrievedOrder, err := client.GetOrderWithContext(context.Background(), "1")
	if err != nil && errors.Is(err, ParamsErr) {
		// 不会走到这里，因为err和common.ParamsErr不相等
		t.Fatal(err)
	}
	t.Logf("order: %+v", retrievedOrder)
}

func TestRPCError1(t *testing.T) {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)

	//// Get Order
	//resp, err := client.GetOrderWithContext(context.Background(), "1")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if resp.Errno != order.BizErrno_Ok {
	//	t.Fatal(resp.Msg)
	//}
	//t.Logf("GetOrder Response -> : %+v", resp.Data)
}

func TestTarsError(t *testing.T) {
	ok := tars.Errorf(basef.TARSSERVERSUCCESS, "ok")
	fmt.Println(ok)

	serverNoFuncErr := tars.Errorf(basef.TARSSERVERNOFUNCERR, "服务器端没有该函数")
	fmt.Println(serverNoFuncErr)
}

func TestTarsError1(t *testing.T) {
	comm := tars.GetCommunicator()
	client := new(order.OrderManagement)
	obj := "Test.OrderServer.OrderObj@tcp -h 127.0.0.1 -p 8080 -t 60000"
	comm.StringToProxy(obj, client)

	// Get Order
	ord, err := client.GetOrderWithContext(context.Background(), "2")
	if err != nil {
		var tarsErr = new(tars.Error)
		if ok := errors.As(err, &tarsErr); ok && tarsErr.Code == http.StatusNotFound {
			t.Logf("code: %d, msg: %s", tarsErr.Code, tarsErr.Message)
		} else {
			t.Fatal(err)
		}
		return
	}
	t.Logf("GetOrder Response -> : %+v", ord)
}
