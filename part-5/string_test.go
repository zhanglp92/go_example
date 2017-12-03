package examplePart5

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStringSlice(t *testing.T) {
	s0 := "0123456789"
	s1 := s0[2:3]

	t.Logf("%+v", (*reflect.StringHeader)(unsafe.Pointer(&s0)))
	t.Logf("%+v", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func TestArraySlice(t *testing.T) {
	a0 := []int{1, 2, 3, 4}
	a1 := a0[1:2]

	// has same address.
	t.Logf("%+v", unsafe.Pointer(&a0[1]))
	t.Logf("%+v", unsafe.Pointer(&a1[0]))
}

func returnArg(t *testing.T, s interface{}) interface{} {
	// val := s.(string)
	switch s.(type) {
	case string:
		val := s.(string)
		t.Logf("=>%+v", (*reflect.StringHeader)(unsafe.Pointer(&val)))
	case int:
		val := s.(int)
		t.Logf("=>%+v", unsafe.Pointer(&val))
	}

	return s
}

func TestTransmitString(t *testing.T) {
	s0 := "test transmit string."

	s1 := returnArg(t, s0).(string)

	t.Logf("%+v", (*reflect.StringHeader)(unsafe.Pointer(&s0)))
	t.Logf("%+v", (*reflect.StringHeader)(unsafe.Pointer(&s1)))

	sint0 := 10
	sint1 := returnArg(t, sint0).(int)
	t.Logf("%+v", unsafe.Pointer(&sint0))
	t.Logf("%+v", unsafe.Pointer(&sint1))

}
