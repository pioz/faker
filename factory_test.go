package factory_test

import (
	"errors"
	"testing"
	"time"

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

func TestMapBuild(t *testing.T) {
	factory.SetSeed(604)
	s := &struct {
		Field1 map[string]int         `factory:"len=1"`
		Field2 *map[int]time.Duration `factory:"len=2"`
		Field3 map[int]*int           `factory:"len=2"`
		Field4 map[*int]int           `factory:"len=2"`
		Field5 map[*int]*int          `factory:"len=2"`
		Field6 map[int]int
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, 1, len(s.Field1))
	assert.Equal(t, -408064131, s.Field1["pBrjYVHE4Letot5yTTsmIopXpnBff32wlwbLiH0EPQ7KR9upBp6zb2HqCkmQkzR3x71NOyPItdyK9AIqn0F9ROrYnDURrUNLMiDHSgI0ACoyFkguWYZTfOYjb"])

	assert.Equal(t, 2, len(*s.Field2))
	assert.Equal(t, "-567614h16m7.375390531s", (*s.Field2)[-1984597982].String())
	assert.Equal(t, "-1393246h51m23.927637982s", (*s.Field2)[-473279944].String())

	assert.Equal(t, 2, len(s.Field3))
	assert.Equal(t, 140005533, *s.Field3[-664005745])
	assert.Equal(t, -1314422524, *s.Field3[-107123621])

	assert.Equal(t, 2, len(s.Field4))
	assert.Equal(t, 2, len(s.Field5))

	assert.Equal(t, 7, len(s.Field6))
	assert.Equal(t, 1242048708, s.Field6[-503311227])
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

func TestNoErrorOnNotSupportedType(t *testing.T) {
	s := &struct {
		Channel chan int
	}{}

	err := factory.Build(&s)
	assert.Nil(t, err)
	assert.Equal(t, chan int(nil), s.Channel)
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

type testParent struct {
	Name   string
	Parent *testParent
}

func TestRecursiveStructBuild(t *testing.T) {
	factory.SetSeed(621)
	s := testParent{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "rpCLQtCMxnMcaT6CChHFohkI58zXOnTNQz5c2J25iD7VSU0SrNVxoADmVRx66L1tQx6ljCDaetBPRNqfLC13hczryEpD3YXCV8nGOsHUIpYGdcTjmiH0XTn", s.Name)
	assert.Equal(t, (*testParent)(nil), s.Parent)
}

type testNode1 struct {
	Name string
	Node *testNode2
}

type testNode2 struct {
	Name string
	Node *testNode1 `factory:"-"` // if we remove the ignore tag "-" we will have an infinite loop
}

func TestMutualRecursiveStructBuild(t *testing.T) {
	factory.SetSeed(622)
	s := testNode1{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "E", s.Name)
	assert.Equal(t, "0LX6dzNLnPogQSisq70w3F6YmFBQVk1CHiOoV9IUMNE2YfI4AvyupkIwiGVOTwpaHat1LOP3IVbz1X81M3ReM", s.Node.Name)
	assert.Nil(t, nil, s.Node.Node)
}

type userTest struct {
	Username      string         `factory:"username"`
	Email         string         `factory:"email"`
	Comments      []string       `factory:"sentence;len=3"`
	Feedbacks     map[string]int `factory:"feedbacks"`
	FakeFeedbacks map[string]int
}

// Generic quite complete test on struct
func TestStructBuild(t *testing.T) {
	factory.SetSeed(620)
	factory.RegisterProvider("coin", "string", func(...string) (interface{}, error) {
		if factory.Bool() {
			return "head", nil
		} else {
			return "tail", nil
		}
	})

	factory.RegisterProvider("feedbacks", "map[string]int", func(...string) (interface{}, error) {
		f := make(map[string]int)
		f["power"] = 2
		f["speed"] = 3
		f["intellect"] = 4
		return f, nil
	})

	s := &struct {
		Number1  int `factory:"intinrange(0,5)"`
		Number2  uint
		Coin     string `factory:"coin"`
		List     []int  `factory:"len=4"`
		NotEmpty string
		Unknown  chan int
		User1    userTest
		User2    *userTest
	}{NotEmpty: "not changed"}

	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, 3, s.Number1)
	assert.Equal(t, uint(1593375208), s.Number2)
	assert.Equal(t, "tail", s.Coin)
	assert.Equal(t, []int{-1641723469, -1455267700, -2127403521, -434130014}, s.List)
	assert.Equal(t, "not changed", s.NotEmpty)
	assert.Equal(t, chan int(nil), s.Unknown)

	assert.Equal(t, "strawboard", s.User1.Username)
	assert.Equal(t, "palecek@stat.name", s.User1.Email)
	assert.Equal(t, 3, len(s.User1.Comments))
	assert.Equal(t, "We can assume that any instance of an owl can be construed as a patient lemon.", s.User1.Comments[0])
	assert.Equal(t, map[string]int{"power": 2, "speed": 3, "intellect": 4}, s.User1.Feedbacks)
	assert.Equal(t, 7, len(s.User1.FakeFeedbacks))
	assert.Equal(t, -421206271, s.User1.FakeFeedbacks["8YV0qeIEnt3ixbFlO2sSFxDMdogn9GZlBD6mV1kgJtcFAvhx0D7C1Mafs5GJecb4z8G9MVWAX2NOjdxsbU2PYig9yZHR6Jc6jXxJnwc5"])

	assert.Equal(t, "roede", s.User2.Username)
	assert.Equal(t, "monarski@unceremonious.name", s.User2.Email)
	assert.Equal(t, 3, len(s.User2.Comments))
	assert.Equal(t, "The first warmhearted alligator is, in its own way, a tiger.", s.User2.Comments[0])
	assert.Equal(t, map[string]int{"power": 2, "speed": 3, "intellect": 4}, s.User2.Feedbacks)
	assert.Equal(t, 6, len(s.User2.FakeFeedbacks))
	assert.Equal(t, 982485615, s.User2.FakeFeedbacks["D8N3TcJSnbQNXxD4MHI7QQO4VGNj0E3Oat4OvZdRGfrZVrtf55iQ"])
}
