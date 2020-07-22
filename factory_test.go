package factory_test

import (
	"errors"
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestPtrBuild(t *testing.T) {
	factory.SetSeed(601)
	s := &struct {
		IntField    int
		PtrIntField *int
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, 1222834422, s.IntField)
	assert.Equal(t, 1743740582, *s.PtrIntField)
}

func TestSliceBuild(t *testing.T) {
	factory.SetSeed(600)
	s := &struct {
		SliceIntField    []int  `factory:"len=1"`
		PtrSliceIntField *[]int `factory:"len=2"`
		SlicePtrIntField []*int `factory:"len=3"`
		SliceIntField2   []int
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, 1, len(s.SliceIntField))
	assert.Equal(t, -1982024679, s.SliceIntField[0])

	assert.Equal(t, 2, len(*s.PtrSliceIntField))
	assert.Equal(t, 818547590, (*s.PtrSliceIntField)[0])
	assert.Equal(t, -1723388666, (*s.PtrSliceIntField)[1])

	assert.Equal(t, 3, len(s.SlicePtrIntField))
	assert.Equal(t, -152085377, *s.SlicePtrIntField[0])
	assert.Equal(t, 1775168364, *s.SlicePtrIntField[1])
	assert.Equal(t, 546278985, *s.SlicePtrIntField[2])

	assert.Equal(t, 8, len(s.SliceIntField2))
}

func TestTagFuncCallBuild(t *testing.T) {
	factory.RegisterProvider("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})

	s := &struct {
		Ping string `factory:"Ping"`
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "pong", s.Ping)
}

func TestTagFuncCallCaseSensitiveBuild(t *testing.T) {
	factory.RegisterProvider("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})

	s := &struct {
		Ping1 string `factory:"Ping"`
		Ping2 string `factory:"PING"`
		Ping3 string `factory:"ping"`
		Ping4 string `factory:"pInG"`
		Ping5 string `factory:"ping(a,b,c)"`
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "pong", s.Ping1)
	assert.Equal(t, "pong", s.Ping2)
	assert.Equal(t, "pong", s.Ping3)
	assert.Equal(t, "pong", s.Ping4)
	assert.Equal(t, "pong", s.Ping5)
}

func TestTagFuncCallWithParamsBuild(t *testing.T) {
	factory.RegisterProvider("Temperature", "string", func(params ...string) (interface{}, error) {
		if len(params) == 1 {
			return params[0], nil
		}
		return "nil", nil
	})

	s := &struct {
		Temp1 string `factory:"Temperature"`
		Temp2 string `factory:"Temperature(hot)"`
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "nil", s.Temp1)
	assert.Equal(t, "hot", s.Temp2)
}

func TestTagFuncCallReturnErrorBuild(t *testing.T) {
	factory.RegisterProvider("Error", "string", func(params ...string) (interface{}, error) {
		return nil, errors.New("This is an error")
	})

	s := &struct {
		Err string `factory:"Error"`
	}{}

	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "This is an error", err.Error())
}

func TestTagFuncCallNotSupportedTypeBuild(t *testing.T) {
	factory.RegisterProvider("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})

	s := &struct {
		Ping int `factory:"Ping"`
	}{}

	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid faker function 'Ping' for type 'int'", err.Error())
}

func TestNotEmptyValueBuild(t *testing.T) {
	factory.SetSeed(603)
	s := &struct {
		Field1 int
		Field2 int
	}{Field2: 21}

	err := factory.Build(&s)
	assert.Nil(t, err)
	assert.Equal(t, -236043479, s.Field1)
	assert.Equal(t, 21, s.Field2)
}

func TestSkipTag(t *testing.T) {
	s := &struct {
		Field    int  `factory:"-"`
		PtrField *int `factory:"-"`
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Empty(t, s.Field)
	assert.Nil(t, s.PtrField)
}

func TestUniqueTag(t *testing.T) {
	factory.SetSeed(602)
	s1 := &struct {
		Field int `factory:"IntInRange(0,1);unique"`
	}{}
	s2 := &struct {
		Field int `factory:"IntInRange(0,1);unique"`
	}{}
	s3 := &struct {
		Field int `factory:"IntInRange(0,1);unique"`
	}{}

	err := factory.Build(&s1)
	assert.Nil(t, err)
	t.Log(s1)
	assert.Equal(t, 1, s1.Field)

	err = factory.Build(&s2)
	assert.Nil(t, err)
	t.Log(s2)
	assert.Equal(t, 0, s2.Field)

	err = factory.Build(&s3)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to generate a unique value", err.Error())
}

func TestInterfaceNotAllowed(t *testing.T) {
	var i interface{}
	err := factory.Build(i)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input interface{} not allowed", err.Error())
}

func TestNoPtrNotAllowed(t *testing.T) {
	var i int
	err := factory.Build(i)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input is not a pointer", err.Error())
}

func TestNilNotAllowed(t *testing.T) {
	err := factory.Build(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input interface{} not allowed", err.Error())
}

func TestStructBuild(t *testing.T) {
	// Generic quite complete test on struct
}
