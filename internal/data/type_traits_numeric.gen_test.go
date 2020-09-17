// Code generated by type_traits_numeric.gen_test.go.tmpl. DO NOT EDIT.
package data_test

import (
	"reflect"
	"testing"

	"github.com/factset/go-drill/internal/data"
)

func TestInt64Traits(t *testing.T) {
	const N = 10
	b1 := data.Int64Traits.CastToBytes([]int64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Int64Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, int64(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]int64, N)
	data.Int64Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestInt32Traits(t *testing.T) {
	const N = 10
	b1 := data.Int32Traits.CastToBytes([]int32{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Int32Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, int32(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]int32, N)
	data.Int32Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestFloat64Traits(t *testing.T) {
	const N = 10
	b1 := data.Float64Traits.CastToBytes([]float64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Float64Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, float64(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]float64, N)
	data.Float64Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestUint64Traits(t *testing.T) {
	const N = 10
	b1 := data.Uint64Traits.CastToBytes([]uint64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Uint64Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, uint64(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]uint64, N)
	data.Uint64Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestUint32Traits(t *testing.T) {
	const N = 10
	b1 := data.Uint32Traits.CastToBytes([]uint32{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Uint32Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, uint32(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]uint32, N)
	data.Uint32Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestFloat32Traits(t *testing.T) {
	const N = 10
	b1 := data.Float32Traits.CastToBytes([]float32{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Float32Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, float32(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]float32, N)
	data.Float32Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestInt16Traits(t *testing.T) {
	const N = 10
	b1 := data.Int16Traits.CastToBytes([]int16{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Int16Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, int16(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]int16, N)
	data.Int16Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestUint16Traits(t *testing.T) {
	const N = 10
	b1 := data.Uint16Traits.CastToBytes([]uint16{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Uint16Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, uint16(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]uint16, N)
	data.Uint16Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestInt8Traits(t *testing.T) {
	const N = 10
	b1 := data.Int8Traits.CastToBytes([]int8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Int8Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, int8(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]int8, N)
	data.Int8Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}

func TestUint8Traits(t *testing.T) {
	const N = 10
	b1 := data.Uint8Traits.CastToBytes([]uint8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v1 := data.Uint8Traits.CastFromBytes(b1)
	for i, v := range v1 {
		if got, want := v, uint8(i); got != want {
			t.Fatalf("invalid value[%d]. got=%v, want=%v", i, got, want)
		}
	}

	v2 := make([]uint8, N)
	data.Uint8Traits.Copy(v2, v1)

	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("invalid values:\nv1=%v\nv2=%v\n", v1, v2)
	}
}
