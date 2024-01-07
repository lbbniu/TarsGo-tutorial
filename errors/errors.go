package errors

import "errors"

var (
	ParamsErr = errors.New("params err")
	BizErr    = errors.New("biz err")
)
