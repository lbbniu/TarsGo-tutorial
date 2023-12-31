// Package order comment
// This file was generated by tars2go 1.2.2
// Generated from order.tars
package order

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/endpoint"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	tarstrace "github.com/TarsCloud/TarsGo/tars/util/trace"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
)

// OrderManagement struct
type OrderManagement struct {
	servant m.Servant
}

// SetServant sets servant for the service.
func (obj *OrderManagement) SetServant(servant m.Servant) {
	obj.servant = servant
}

// TarsSetTimeout sets the timeout for the servant which is in ms.
func (obj *OrderManagement) TarsSetTimeout(timeout int) {
	obj.servant.TarsSetTimeout(timeout)
}

// TarsSetProtocol sets the protocol for the servant.
func (obj *OrderManagement) TarsSetProtocol(p m.Protocol) {
	obj.servant.TarsSetProtocol(p)
}

// Endpoints returns all active endpoint.Endpoint
func (obj *OrderManagement) Endpoints() []*endpoint.Endpoint {
	return obj.servant.Endpoints()
}

// AddServant adds servant  for the service.
func (obj *OrderManagement) AddServant(imp OrderManagementServant, servantObj string) {
	tars.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *OrderManagement) AddServantWithContext(imp OrderManagementServantWithContext, servantObj string) {
	tars.AddServantWithContext(obj, imp, servantObj)
}

// GetOrder is the proxy function for the method defined in the tars file, with the context
func (obj *OrderManagement) GetOrder(orderId string, opts ...map[string]string) (Order, error) {
	return obj.GetOrderWithContext(context.Background(), orderId, opts...)
}

// GetOrderWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *OrderManagement) GetOrderWithContext(tarsCtx context.Context, orderId string, opts ...map[string]string) (ret Order, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(orderId, 1)
	if err != nil {
		return ret, err
	}

	trace, ok := current.GetTarsTrace(tarsCtx)
	if ok && trace.Call() {
		var traceParam string
		trace.NewSpan()
		traceParamFlag := trace.NeedTraceParam(tarstrace.EstCS, uint(buf.Len()))
		if traceParamFlag == tarstrace.EnpNormal {
			value := map[string]interface{}{}
			value["orderId"] = orderId
			jm, _ := json.Marshal(value)
			traceParam = string(jm)
		} else if traceParamFlag == tarstrace.EnpOverMaxLen {
			traceParam = `{"trace_param_over_max_len":true}`
		}
		tars.Trace(trace.GetTraceKey(tarstrace.EstCS), tarstrace.AnnotationCS, tars.GetClientConfig().ModuleName, obj.servant.Name(), "GetOrder", 0, traceParam, "")
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "getOrder", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = ret.ReadBlock(readBuf, 0, true)
	if err != nil {
		return ret, err
	}

	if ok && trace.Call() {
		var traceParam string
		traceParamFlag := trace.NeedTraceParam(tarstrace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == tarstrace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			jm, _ := json.Marshal(value)
			traceParam = string(jm)
		} else if traceParamFlag == tarstrace.EnpOverMaxLen {
			traceParam = `{"trace_param_over_max_len":true}`
		}
		tars.Trace(trace.GetTraceKey(tarstrace.EstCR), tarstrace.AnnotationCR, tars.GetClientConfig().ModuleName, obj.servant.Name(), "GetOrder", tarsResp.IRet, traceParam, "")
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// GetOrderOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *OrderManagement) GetOrderOneWayWithContext(tarsCtx context.Context, orderId string, opts ...map[string]string) (ret Order, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(orderId, 1)
	if err != nil {
		return ret, err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "getOrder", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}

type OrderManagementServant interface {
	GetOrder(orderId string) (ret Order, err error)
}
type OrderManagementServantWithContext interface {
	GetOrder(tarsCtx context.Context, orderId string) (ret Order, err error)
}

// Dispatch is used to call the server side implement for the method defined in the tars file. withContext shows using context or not.
func (obj *OrderManagement) Dispatch(tarsCtx context.Context, val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	buf := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "getOrder":
		var orderId string

		if tarsReq.IVersion == basef.TARSVERSION {

			err = readBuf.ReadString(&orderId, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("orderId", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = readBuf.ReadString(&orderId, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["orderId"])
				if err = json.Unmarshal(jsonStr, &orderId); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		trace, ok := current.GetTarsTrace(tarsCtx)
		if ok && trace.Call() {
			var traceParam string
			traceParamFlag := trace.NeedTraceParam(tarstrace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == tarstrace.EnpNormal {
				value := map[string]interface{}{}
				value["orderId"] = orderId
				jm, _ := json.Marshal(value)
				traceParam = string(jm)
			} else if traceParamFlag == tarstrace.EnpOverMaxLen {
				traceParam = `{"trace_param_over_max_len":true}`
			}
			tars.Trace(trace.GetTraceKey(tarstrace.EstSR), tarstrace.AnnotationSR, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "getOrder", 0, traceParam, "")
		}

		var funRet Order
		if !withContext {
			imp := val.(OrderManagementServant)
			funRet, err = imp.GetOrder(orderId)
		} else {
			imp := val.(OrderManagementServantWithContext)
			funRet, err = imp.GetOrder(tarsCtx, orderId)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

			err = funRet.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			err = funRet.WriteBlock(buf, 0)
			if err != nil {
				return err
			}

			rspTup.PutBuffer("", buf.ToBytes())
			rspTup.PutBuffer("tars_ret", buf.ToBytes())

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["tars_ret"] = funRet

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

		if ok && trace.Call() {
			var traceParam string
			traceParamFlag := trace.NeedTraceParam(tarstrace.EstSS, uint(buf.Len()))
			if traceParamFlag == tarstrace.EnpNormal {
				value := map[string]interface{}{}
				value[""] = funRet
				jm, _ := json.Marshal(value)
				traceParam = string(jm)
			} else if traceParamFlag == tarstrace.EnpOverMaxLen {
				traceParam = `{"trace_param_over_max_len":true}`
			}
			tars.Trace(trace.GetTraceKey(tarstrace.EstSS), tarstrace.AnnotationSS, tars.GetClientConfig().ModuleName, tarsReq.SServantName, "getOrder", 0, traceParam, "")
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(tarsCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(tarsCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}
