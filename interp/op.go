package interp

import (
	"go/token"
	"reflect"
)

type binop func(x, y interface{}) interface{}

var binops = map[token.Token]map[reflect.Kind]binop{
	token.ADD:     {},
	token.SUB:     {},
	token.MUL:     {},
	token.QUO:     {},
	token.REM:     {},
	token.AND:     {},
	token.OR:      {},
	token.XOR:     {},
	token.AND_NOT: {},
}

var kindtypes = map[reflect.Kind]reflect.Type{
	reflect.String:     reflect.TypeOf(""),
	reflect.Float32:    reflect.TypeOf(float32(0)),
	reflect.Float64:    reflect.TypeOf(float64(0)),
	reflect.Complex64:  reflect.TypeOf(complex64(0)),
	reflect.Complex128: reflect.TypeOf(complex128(0)),
	reflect.Uint8:      reflect.TypeOf(uint8(0)),
	reflect.Uint16:     reflect.TypeOf(uint16(0)),
	reflect.Uint32:     reflect.TypeOf(uint32(0)),
	reflect.Uint64:     reflect.TypeOf(uint64(0)),
	reflect.Int8:       reflect.TypeOf(int8(0)),
	reflect.Int16:      reflect.TypeOf(int16(0)),
	reflect.Int32:      reflect.TypeOf(int32(0)),
	reflect.Int64:      reflect.TypeOf(int64(0)),
	reflect.Uint:       reflect.TypeOf(uint(0)),
	reflect.Int:        reflect.TypeOf(int(0)),
	reflect.Uintptr:    reflect.TypeOf(uintptr(0)),
}

func init() {
	binops[token.ADD][reflect.String] = addString
	binops[token.ADD][reflect.Float32] = addFloat32
	binops[token.SUB][reflect.Float32] = subFloat32
	binops[token.MUL][reflect.Float32] = mulFloat32
	binops[token.QUO][reflect.Float32] = quoFloat32
	binops[token.ADD][reflect.Float64] = addFloat64
	binops[token.SUB][reflect.Float64] = subFloat64
	binops[token.MUL][reflect.Float64] = mulFloat64
	binops[token.QUO][reflect.Float64] = quoFloat64
	binops[token.ADD][reflect.Complex64] = addComplex64
	binops[token.SUB][reflect.Complex64] = subComplex64
	binops[token.MUL][reflect.Complex64] = mulComplex64
	binops[token.QUO][reflect.Complex64] = quoComplex64
	binops[token.ADD][reflect.Complex128] = addComplex128
	binops[token.SUB][reflect.Complex128] = subComplex128
	binops[token.MUL][reflect.Complex128] = mulComplex128
	binops[token.QUO][reflect.Complex128] = quoComplex128
	binops[token.ADD][reflect.Uint8] = addUint8
	binops[token.SUB][reflect.Uint8] = subUint8
	binops[token.MUL][reflect.Uint8] = mulUint8
	binops[token.QUO][reflect.Uint8] = quoUint8
	binops[token.REM][reflect.Uint8] = remUint8
	binops[token.AND][reflect.Uint8] = andUint8
	binops[token.OR][reflect.Uint8] = orUint8
	binops[token.XOR][reflect.Uint8] = xorUint8
	binops[token.AND_NOT][reflect.Uint8] = and_notUint8
	binops[token.ADD][reflect.Uint16] = addUint16
	binops[token.SUB][reflect.Uint16] = subUint16
	binops[token.MUL][reflect.Uint16] = mulUint16
	binops[token.QUO][reflect.Uint16] = quoUint16
	binops[token.REM][reflect.Uint16] = remUint16
	binops[token.AND][reflect.Uint16] = andUint16
	binops[token.OR][reflect.Uint16] = orUint16
	binops[token.XOR][reflect.Uint16] = xorUint16
	binops[token.AND_NOT][reflect.Uint16] = and_notUint16
	binops[token.ADD][reflect.Uint32] = addUint32
	binops[token.SUB][reflect.Uint32] = subUint32
	binops[token.MUL][reflect.Uint32] = mulUint32
	binops[token.QUO][reflect.Uint32] = quoUint32
	binops[token.REM][reflect.Uint32] = remUint32
	binops[token.AND][reflect.Uint32] = andUint32
	binops[token.OR][reflect.Uint32] = orUint32
	binops[token.XOR][reflect.Uint32] = xorUint32
	binops[token.AND_NOT][reflect.Uint32] = and_notUint32
	binops[token.ADD][reflect.Uint64] = addUint64
	binops[token.SUB][reflect.Uint64] = subUint64
	binops[token.MUL][reflect.Uint64] = mulUint64
	binops[token.QUO][reflect.Uint64] = quoUint64
	binops[token.REM][reflect.Uint64] = remUint64
	binops[token.AND][reflect.Uint64] = andUint64
	binops[token.OR][reflect.Uint64] = orUint64
	binops[token.XOR][reflect.Uint64] = xorUint64
	binops[token.AND_NOT][reflect.Uint64] = and_notUint64
	binops[token.ADD][reflect.Int8] = addInt8
	binops[token.SUB][reflect.Int8] = subInt8
	binops[token.MUL][reflect.Int8] = mulInt8
	binops[token.QUO][reflect.Int8] = quoInt8
	binops[token.REM][reflect.Int8] = remInt8
	binops[token.AND][reflect.Int8] = andInt8
	binops[token.OR][reflect.Int8] = orInt8
	binops[token.XOR][reflect.Int8] = xorInt8
	binops[token.AND_NOT][reflect.Int8] = and_notInt8
	binops[token.ADD][reflect.Int16] = addInt16
	binops[token.SUB][reflect.Int16] = subInt16
	binops[token.MUL][reflect.Int16] = mulInt16
	binops[token.QUO][reflect.Int16] = quoInt16
	binops[token.REM][reflect.Int16] = remInt16
	binops[token.AND][reflect.Int16] = andInt16
	binops[token.OR][reflect.Int16] = orInt16
	binops[token.XOR][reflect.Int16] = xorInt16
	binops[token.AND_NOT][reflect.Int16] = and_notInt16
	binops[token.ADD][reflect.Int32] = addInt32
	binops[token.SUB][reflect.Int32] = subInt32
	binops[token.MUL][reflect.Int32] = mulInt32
	binops[token.QUO][reflect.Int32] = quoInt32
	binops[token.REM][reflect.Int32] = remInt32
	binops[token.AND][reflect.Int32] = andInt32
	binops[token.OR][reflect.Int32] = orInt32
	binops[token.XOR][reflect.Int32] = xorInt32
	binops[token.AND_NOT][reflect.Int32] = and_notInt32
	binops[token.ADD][reflect.Int64] = addInt64
	binops[token.SUB][reflect.Int64] = subInt64
	binops[token.MUL][reflect.Int64] = mulInt64
	binops[token.QUO][reflect.Int64] = quoInt64
	binops[token.REM][reflect.Int64] = remInt64
	binops[token.AND][reflect.Int64] = andInt64
	binops[token.OR][reflect.Int64] = orInt64
	binops[token.XOR][reflect.Int64] = xorInt64
	binops[token.AND_NOT][reflect.Int64] = and_notInt64
	binops[token.ADD][reflect.Uint] = addUint
	binops[token.SUB][reflect.Uint] = subUint
	binops[token.MUL][reflect.Uint] = mulUint
	binops[token.QUO][reflect.Uint] = quoUint
	binops[token.REM][reflect.Uint] = remUint
	binops[token.AND][reflect.Uint] = andUint
	binops[token.OR][reflect.Uint] = orUint
	binops[token.XOR][reflect.Uint] = xorUint
	binops[token.AND_NOT][reflect.Uint] = and_notUint
	binops[token.ADD][reflect.Int] = addInt
	binops[token.SUB][reflect.Int] = subInt
	binops[token.MUL][reflect.Int] = mulInt
	binops[token.QUO][reflect.Int] = quoInt
	binops[token.REM][reflect.Int] = remInt
	binops[token.AND][reflect.Int] = andInt
	binops[token.OR][reflect.Int] = orInt
	binops[token.XOR][reflect.Int] = xorInt
	binops[token.AND_NOT][reflect.Int] = and_notInt
	binops[token.ADD][reflect.Uintptr] = addUintptr
	binops[token.SUB][reflect.Uintptr] = subUintptr
	binops[token.MUL][reflect.Uintptr] = mulUintptr
	binops[token.QUO][reflect.Uintptr] = quoUintptr
	binops[token.REM][reflect.Uintptr] = remUintptr
	binops[token.AND][reflect.Uintptr] = andUintptr
	binops[token.OR][reflect.Uintptr] = orUintptr
	binops[token.XOR][reflect.Uintptr] = xorUintptr
	binops[token.AND_NOT][reflect.Uintptr] = and_notUintptr

}

func addString(x, y interface{}) interface{} {
	return x.(string) + y.(string)
}

func addFloat32(x, y interface{}) interface{} {
	return x.(float32) + y.(float32)
}

func subFloat32(x, y interface{}) interface{} {
	return x.(float32) - y.(float32)
}

func mulFloat32(x, y interface{}) interface{} {
	return x.(float32) * y.(float32)
}

func quoFloat32(x, y interface{}) interface{} {
	return x.(float32) / y.(float32)
}

func addFloat64(x, y interface{}) interface{} {
	return x.(float64) + y.(float64)
}

func subFloat64(x, y interface{}) interface{} {
	return x.(float64) - y.(float64)
}

func mulFloat64(x, y interface{}) interface{} {
	return x.(float64) * y.(float64)
}

func quoFloat64(x, y interface{}) interface{} {
	return x.(float64) / y.(float64)
}

func addComplex64(x, y interface{}) interface{} {
	return x.(complex64) + y.(complex64)
}

func subComplex64(x, y interface{}) interface{} {
	return x.(complex64) - y.(complex64)
}

func mulComplex64(x, y interface{}) interface{} {
	return x.(complex64) * y.(complex64)
}

func quoComplex64(x, y interface{}) interface{} {
	return x.(complex64) / y.(complex64)
}

func addComplex128(x, y interface{}) interface{} {
	return x.(complex128) + y.(complex128)
}

func subComplex128(x, y interface{}) interface{} {
	return x.(complex128) - y.(complex128)
}

func mulComplex128(x, y interface{}) interface{} {
	return x.(complex128) * y.(complex128)
}

func quoComplex128(x, y interface{}) interface{} {
	return x.(complex128) / y.(complex128)
}

func addUint8(x, y interface{}) interface{} {
	return x.(uint8) + y.(uint8)
}

func subUint8(x, y interface{}) interface{} {
	return x.(uint8) - y.(uint8)
}

func mulUint8(x, y interface{}) interface{} {
	return x.(uint8) * y.(uint8)
}

func quoUint8(x, y interface{}) interface{} {
	return x.(uint8) / y.(uint8)
}

func remUint8(x, y interface{}) interface{} {
	return x.(uint8) % y.(uint8)
}

func andUint8(x, y interface{}) interface{} {
	return x.(uint8) & y.(uint8)
}

func orUint8(x, y interface{}) interface{} {
	return x.(uint8) | y.(uint8)
}

func xorUint8(x, y interface{}) interface{} {
	return x.(uint8) ^ y.(uint8)
}

func and_notUint8(x, y interface{}) interface{} {
	return x.(uint8) &^ y.(uint8)
}

func addUint16(x, y interface{}) interface{} {
	return x.(uint16) + y.(uint16)
}

func subUint16(x, y interface{}) interface{} {
	return x.(uint16) - y.(uint16)
}

func mulUint16(x, y interface{}) interface{} {
	return x.(uint16) * y.(uint16)
}

func quoUint16(x, y interface{}) interface{} {
	return x.(uint16) / y.(uint16)
}

func remUint16(x, y interface{}) interface{} {
	return x.(uint16) % y.(uint16)
}

func andUint16(x, y interface{}) interface{} {
	return x.(uint16) & y.(uint16)
}

func orUint16(x, y interface{}) interface{} {
	return x.(uint16) | y.(uint16)
}

func xorUint16(x, y interface{}) interface{} {
	return x.(uint16) ^ y.(uint16)
}

func and_notUint16(x, y interface{}) interface{} {
	return x.(uint16) &^ y.(uint16)
}

func addUint32(x, y interface{}) interface{} {
	return x.(uint32) + y.(uint32)
}

func subUint32(x, y interface{}) interface{} {
	return x.(uint32) - y.(uint32)
}

func mulUint32(x, y interface{}) interface{} {
	return x.(uint32) * y.(uint32)
}

func quoUint32(x, y interface{}) interface{} {
	return x.(uint32) / y.(uint32)
}

func remUint32(x, y interface{}) interface{} {
	return x.(uint32) % y.(uint32)
}

func andUint32(x, y interface{}) interface{} {
	return x.(uint32) & y.(uint32)
}

func orUint32(x, y interface{}) interface{} {
	return x.(uint32) | y.(uint32)
}

func xorUint32(x, y interface{}) interface{} {
	return x.(uint32) ^ y.(uint32)
}

func and_notUint32(x, y interface{}) interface{} {
	return x.(uint32) &^ y.(uint32)
}

func addUint64(x, y interface{}) interface{} {
	return x.(uint64) + y.(uint64)
}

func subUint64(x, y interface{}) interface{} {
	return x.(uint64) - y.(uint64)
}

func mulUint64(x, y interface{}) interface{} {
	return x.(uint64) * y.(uint64)
}

func quoUint64(x, y interface{}) interface{} {
	return x.(uint64) / y.(uint64)
}

func remUint64(x, y interface{}) interface{} {
	return x.(uint64) % y.(uint64)
}

func andUint64(x, y interface{}) interface{} {
	return x.(uint64) & y.(uint64)
}

func orUint64(x, y interface{}) interface{} {
	return x.(uint64) | y.(uint64)
}

func xorUint64(x, y interface{}) interface{} {
	return x.(uint64) ^ y.(uint64)
}

func and_notUint64(x, y interface{}) interface{} {
	return x.(uint64) &^ y.(uint64)
}

func addInt8(x, y interface{}) interface{} {
	return x.(int8) + y.(int8)
}

func subInt8(x, y interface{}) interface{} {
	return x.(int8) - y.(int8)
}

func mulInt8(x, y interface{}) interface{} {
	return x.(int8) * y.(int8)
}

func quoInt8(x, y interface{}) interface{} {
	return x.(int8) / y.(int8)
}

func remInt8(x, y interface{}) interface{} {
	return x.(int8) % y.(int8)
}

func andInt8(x, y interface{}) interface{} {
	return x.(int8) & y.(int8)
}

func orInt8(x, y interface{}) interface{} {
	return x.(int8) | y.(int8)
}

func xorInt8(x, y interface{}) interface{} {
	return x.(int8) ^ y.(int8)
}

func and_notInt8(x, y interface{}) interface{} {
	return x.(int8) &^ y.(int8)
}

func addInt16(x, y interface{}) interface{} {
	return x.(int16) + y.(int16)
}

func subInt16(x, y interface{}) interface{} {
	return x.(int16) - y.(int16)
}

func mulInt16(x, y interface{}) interface{} {
	return x.(int16) * y.(int16)
}

func quoInt16(x, y interface{}) interface{} {
	return x.(int16) / y.(int16)
}

func remInt16(x, y interface{}) interface{} {
	return x.(int16) % y.(int16)
}

func andInt16(x, y interface{}) interface{} {
	return x.(int16) & y.(int16)
}

func orInt16(x, y interface{}) interface{} {
	return x.(int16) | y.(int16)
}

func xorInt16(x, y interface{}) interface{} {
	return x.(int16) ^ y.(int16)
}

func and_notInt16(x, y interface{}) interface{} {
	return x.(int16) &^ y.(int16)
}

func addInt32(x, y interface{}) interface{} {
	return x.(int32) + y.(int32)
}

func subInt32(x, y interface{}) interface{} {
	return x.(int32) - y.(int32)
}

func mulInt32(x, y interface{}) interface{} {
	return x.(int32) * y.(int32)
}

func quoInt32(x, y interface{}) interface{} {
	return x.(int32) / y.(int32)
}

func remInt32(x, y interface{}) interface{} {
	return x.(int32) % y.(int32)
}

func andInt32(x, y interface{}) interface{} {
	return x.(int32) & y.(int32)
}

func orInt32(x, y interface{}) interface{} {
	return x.(int32) | y.(int32)
}

func xorInt32(x, y interface{}) interface{} {
	return x.(int32) ^ y.(int32)
}

func and_notInt32(x, y interface{}) interface{} {
	return x.(int32) &^ y.(int32)
}

func addInt64(x, y interface{}) interface{} {
	return x.(int64) + y.(int64)
}

func subInt64(x, y interface{}) interface{} {
	return x.(int64) - y.(int64)
}

func mulInt64(x, y interface{}) interface{} {
	return x.(int64) * y.(int64)
}

func quoInt64(x, y interface{}) interface{} {
	return x.(int64) / y.(int64)
}

func remInt64(x, y interface{}) interface{} {
	return x.(int64) % y.(int64)
}

func andInt64(x, y interface{}) interface{} {
	return x.(int64) & y.(int64)
}

func orInt64(x, y interface{}) interface{} {
	return x.(int64) | y.(int64)
}

func xorInt64(x, y interface{}) interface{} {
	return x.(int64) ^ y.(int64)
}

func and_notInt64(x, y interface{}) interface{} {
	return x.(int64) &^ y.(int64)
}

func addUint(x, y interface{}) interface{} {
	return x.(uint) + y.(uint)
}

func subUint(x, y interface{}) interface{} {
	return x.(uint) - y.(uint)
}

func mulUint(x, y interface{}) interface{} {
	return x.(uint) * y.(uint)
}

func quoUint(x, y interface{}) interface{} {
	return x.(uint) / y.(uint)
}

func remUint(x, y interface{}) interface{} {
	return x.(uint) % y.(uint)
}

func andUint(x, y interface{}) interface{} {
	return x.(uint) & y.(uint)
}

func orUint(x, y interface{}) interface{} {
	return x.(uint) | y.(uint)
}

func xorUint(x, y interface{}) interface{} {
	return x.(uint) ^ y.(uint)
}

func and_notUint(x, y interface{}) interface{} {
	return x.(uint) &^ y.(uint)
}

func addInt(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

func subInt(x, y interface{}) interface{} {
	return x.(int) - y.(int)
}

func mulInt(x, y interface{}) interface{} {
	return x.(int) * y.(int)
}

func quoInt(x, y interface{}) interface{} {
	return x.(int) / y.(int)
}

func remInt(x, y interface{}) interface{} {
	return x.(int) % y.(int)
}

func andInt(x, y interface{}) interface{} {
	return x.(int) & y.(int)
}

func orInt(x, y interface{}) interface{} {
	return x.(int) | y.(int)
}

func xorInt(x, y interface{}) interface{} {
	return x.(int) ^ y.(int)
}

func and_notInt(x, y interface{}) interface{} {
	return x.(int) &^ y.(int)
}

func addUintptr(x, y interface{}) interface{} {
	return x.(uintptr) + y.(uintptr)
}

func subUintptr(x, y interface{}) interface{} {
	return x.(uintptr) - y.(uintptr)
}

func mulUintptr(x, y interface{}) interface{} {
	return x.(uintptr) * y.(uintptr)
}

func quoUintptr(x, y interface{}) interface{} {
	return x.(uintptr) / y.(uintptr)
}

func remUintptr(x, y interface{}) interface{} {
	return x.(uintptr) % y.(uintptr)
}

func andUintptr(x, y interface{}) interface{} {
	return x.(uintptr) & y.(uintptr)
}

func orUintptr(x, y interface{}) interface{} {
	return x.(uintptr) | y.(uintptr)
}

func xorUintptr(x, y interface{}) interface{} {
	return x.(uintptr) ^ y.(uintptr)
}

func and_notUintptr(x, y interface{}) interface{} {
	return x.(uintptr) &^ y.(uintptr)
}
