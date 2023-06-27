package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/lbbniu/TarsGo-tutorial/HelloGo/tars-protocol/Test"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("Test.HelloGo.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(Test.Hello)
	comm.StringToProxy(obj, app)
	var out, i int32
	i = 123
	ret, err := app.Add(i, i*2, &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret, out)
}
