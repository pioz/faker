package faker_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestPtrBuild(t *testing.T) {
	faker.SetSeed(601)
	s := &struct {
		IntField    int
		PtrIntField *int
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, 1222834422, s.IntField)
	assert.Equal(t, 1743740582, *s.PtrIntField)
}

func TestArgumentSliceBuild(t *testing.T) {
	faker.SetSeed(605)
	slice := make([]string, 2)
	err := faker.Build(&slice)
	assert.Nil(t, err)
	assert.Equal(t, "BlRzrRA", slice[0])
	assert.Equal(t, "SHevwfBd", slice[1])
}

func TestArgumentSliceNotAllZeroBuild(t *testing.T) {
	faker.SetSeed(606)
	slice := make([]string, 2)
	slice[0] = "test"
	err := faker.Build(&slice)
	assert.Nil(t, err)
	assert.Equal(t, "test", slice[0])
	assert.Equal(t, "typkI", slice[1])
}

func TestSliceBuild(t *testing.T) {
	faker.SetSeed(600)
	s := &struct {
		SliceIntField    []int  `faker:"len=1"`
		PtrSliceIntField *[]int `faker:"len=2"`
		SlicePtrIntField []*int `faker:"len=3"`
		SliceIntField2   []int
	}{}

	err := faker.Build(&s)
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

	assert.Equal(t, 5, len(s.SliceIntField2))
}

func TestMapBuild(t *testing.T) {
	faker.SetSeed(604)
	s := &struct {
		Field1 map[string]int         `faker:"len=1"`
		Field2 *map[int]time.Duration `faker:"len=2"`
		Field3 map[int]*int           `faker:"len=2"`
		Field4 map[*int]int           `faker:"len=2"`
		Field5 map[*int]*int          `faker:"len=2"`
		Field6 map[int]int
	}{}

	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	t.Log(s.Field2)

	assert.Equal(t, 1, len(s.Field1))
	assert.Equal(t, 414796793, s.Field1["pBrjY"])

	assert.Equal(t, 2, len(*s.Field2))
	assert.Equal(t, "-1670489h52m2.887881426s", (*s.Field2)[-244204264].String())
	assert.Equal(t, "-13678h42m55.366897902s", (*s.Field2)[1577669889].String())

	assert.Equal(t, 2, len(s.Field3))
	assert.Equal(t, 2027576377, *s.Field3[-1215017291])
	assert.Equal(t, -1688325745, *s.Field3[2062137016])

	assert.Equal(t, 2, len(s.Field4))
	assert.Equal(t, 2, len(s.Field5))

	assert.Equal(t, 4, len(s.Field6))
	assert.Equal(t, -1762085720, s.Field6[-1847156268])
}

func TestTagFuncCallBuild(t *testing.T) {
	err := faker.RegisterBuilder("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})
	assert.Nil(t, err)

	s := &struct {
		Ping string `faker:"Ping"`
	}{}

	err = faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "pong", s.Ping)
	err = faker.UnregisterBuilder("Ping", "string")
	assert.Nil(t, err)
}

func TestTagFuncCallCaseSensitiveBuild(t *testing.T) {
	err := faker.RegisterBuilder("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})
	assert.Nil(t, err)

	s := &struct {
		Ping1 string `faker:"Ping"`
		Ping2 string `faker:"PING"`
		Ping3 string `faker:"ping"`
		Ping4 string `faker:"pInG"`
		Ping5 string `faker:"ping(a,b,c)"`
	}{}

	err = faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "pong", s.Ping1)
	assert.Equal(t, "pong", s.Ping2)
	assert.Equal(t, "pong", s.Ping3)
	assert.Equal(t, "pong", s.Ping4)
	assert.Equal(t, "pong", s.Ping5)
	err = faker.UnregisterBuilder("Ping", "string")
	assert.Nil(t, err)
}

func TestTagFuncCallWithParamsBuild(t *testing.T) {
	err := faker.RegisterBuilder("Temperature", "string", func(params ...string) (interface{}, error) {
		if len(params) == 1 {
			return params[0], nil
		}
		return "nil", nil
	})
	assert.Nil(t, err)

	s := &struct {
		Temp1 string `faker:"Temperature"`
		Temp2 string `faker:"Temperature(hot)"`
	}{}

	err = faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "nil", s.Temp1)
	assert.Equal(t, "hot", s.Temp2)
}

func TestTagFuncCallReturnErrorBuild(t *testing.T) {
	err := faker.RegisterBuilder("Error", "string", func(params ...string) (interface{}, error) {
		return nil, errors.New("this is an error")
	})
	assert.Nil(t, err)

	s := &struct {
		Err string `faker:"Error"`
	}{}

	err = faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "this is an error", err.Error())
}

func TestTagFuncCallNotSupportedTypeBuild(t *testing.T) {
	err := faker.RegisterBuilder("Ping", "string", func(params ...string) (interface{}, error) {
		return "pong", nil
	})
	assert.Nil(t, err)

	s := &struct {
		Ping int `faker:"Ping"`
	}{}

	err = faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid faker function 'Ping' for type 'int'", err.Error())
	err = faker.UnregisterBuilder("Ping", "string")
	assert.Nil(t, err)
}

func TestNoErrorOnNotSupportedType(t *testing.T) {
	s := &struct {
		Channel chan int
	}{}

	err := faker.Build(&s)
	assert.Nil(t, err)
	assert.Equal(t, chan int(nil), s.Channel)
}

func TestNotEmptyValueBuild(t *testing.T) {
	faker.SetSeed(603)
	s := &struct {
		Field1 int
		Field2 int
	}{Field2: 21}

	err := faker.Build(&s)
	assert.Nil(t, err)
	assert.Equal(t, -236043479, s.Field1)
	assert.Equal(t, 21, s.Field2)
}

func TestSkipTag(t *testing.T) {
	s := &struct {
		Field    int  `faker:"-"`
		PtrField *int `faker:"-"`
	}{}

	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Empty(t, s.Field)
	assert.Nil(t, s.PtrField)
}

func TestUniqueTag(t *testing.T) {
	faker.SetSeed(602)
	type CustomType1 struct {
		Field1 int `faker:"IntInRange(0,1);unique"`
		Field2 int `faker:"IntInRange(0,1);unique"`
	}
	type CustomType2 struct {
		Field1 int `faker:"IntInRange(0,1);unique"`
		Field2 int `faker:"IntInRange(0,1);unique"`
	}

	ct11 := CustomType1{}
	err := faker.Build(&ct11)
	assert.Nil(t, err)
	t.Log(ct11)
	assert.Equal(t, 1, ct11.Field1)
	assert.Equal(t, 1, ct11.Field2)

	ct12 := CustomType1{}
	err = faker.Build(&ct12)
	assert.Nil(t, err)
	t.Log(ct12)
	assert.Equal(t, 0, ct12.Field1)
	assert.Equal(t, 0, ct12.Field2)

	ct13 := CustomType1{}
	err = faker.Build(&ct13)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to generate a unique value for group 'CustomType1-Field1'", err.Error())

	ct21 := CustomType2{}
	err = faker.Build(&ct21)
	assert.Nil(t, err)
	t.Log(ct21)
	assert.Equal(t, 0, ct21.Field1)
	assert.Equal(t, 1, ct21.Field2)
}

func TestInterfaceNotAllowed(t *testing.T) {
	var i interface{}
	err := faker.Build(i)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input interface{} not allowed", err.Error())
}

func TestNoPtrNotAllowed(t *testing.T) {
	var i int
	err := faker.Build(i)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input is not a pointer", err.Error())
}

func TestNilNotAllowed(t *testing.T) {
	err := faker.Build(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "faker.Build input interface{} not allowed", err.Error())
}

func TestWorkWithPrivateFields(t *testing.T) {
	faker.SetSeed(624)
	s := &struct {
		PublicField  int
		privateField int
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
}

func TestErrOnSliceBuild(t *testing.T) {
	faker.SetSeed(625)
	s := &struct {
		Field []int `faker:"StringWithSize(3)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid faker function 'StringWithSize' for type 'int'", err.Error())
}

func TestErrOnMapKeyBuild(t *testing.T) {
	faker.SetSeed(626)
	s := &struct {
		Field map[int]bool `faker:"StringWithSize(3)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid faker function 'StringWithSize' for type 'int'", err.Error())
}

func TestErrOnMapValueBuild(t *testing.T) {
	faker.SetSeed(627)
	s := &struct {
		Field map[int]bool `faker:"IntInRange(1,3)"`
	}{}
	err := faker.Build(&s)
	t.Log(s)
	assert.NotNil(t, err)
	// assert.Equal(t, "invalid faker function 'IntInRange' for type 'bool'", err.Error())
}

type testParent struct {
	Name   string
	Parent *testParent
}

func TestRecursiveStructBuild(t *testing.T) {
	faker.SetSeed(621)
	s := testParent{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "rpC", s.Name)
	assert.Equal(t, (*testParent)(nil), s.Parent)
}

type testNode1 struct {
	Name string
	Node *testNode2
}

type testNode2 struct {
	Name string
	Node *testNode1 `faker:"-"` // if we remove the ignore tag "-" we will have an infinite loop
}

func TestMutualRecursiveStructBuild(t *testing.T) {
	faker.SetSeed(622)
	s := testNode1{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "Ek0L", s.Name)
	assert.Equal(t, "6dzNLn", s.Node.Name)
	assert.Nil(t, nil, s.Node.Node)
}

type userTest struct {
	Username      string         `faker:"username"`
	Email         string         `faker:"email"`
	Comments      []string       `faker:"sentence;len=3"`
	Feedbacks     map[string]int `faker:"feedbacks"`
	FakeFeedbacks map[string]int
}

// Generic quite complete test on struct
func TestStructBuild(t *testing.T) {
	faker.SetSeed(620)
	err := faker.RegisterBuilder("coin", "string", func(...string) (interface{}, error) {
		if faker.Bool() {
			return "head", nil
		} else {
			return "tail", nil
		}
	})
	assert.Nil(t, err)

	err = faker.RegisterBuilder("feedbacks", "map[string]int", func(...string) (interface{}, error) {
		f := make(map[string]int)
		f["power"] = 2
		f["speed"] = 3
		f["intellect"] = 4
		return f, nil
	})
	assert.Nil(t, err)

	s := &struct {
		Number1  int `faker:"intinrange(0,5)"`
		Number2  uint
		Coin     string `faker:"coin"`
		List     []int  `faker:"len=4"`
		NotEmpty string
		Unknown  chan int
		User1    userTest
		User2    *userTest
	}{NotEmpty: "not changed"}

	err = faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	t.Log(s.User2)
	assert.Equal(t, 3, s.Number1)
	assert.Equal(t, uint(1593375208), s.Number2)
	assert.Equal(t, "tail", s.Coin)
	assert.Equal(t, []int{-1641723469, -1455267700, -2127403521, -434130014}, s.List)
	assert.Equal(t, "not changed", s.NotEmpty)
	assert.Equal(t, chan int(nil), s.Unknown)

	assert.Equal(t, "chorizo", s.User1.Username)
	assert.Equal(t, "anguished@retool.name", s.User1.Email)
	assert.Equal(t, 3, len(s.User1.Comments))
	assert.Equal(t, "Washing and polishing the car,a grape can hardly be considered a calm scorpion without also being an octopus.", s.User1.Comments[0])
	assert.Equal(t, map[string]int{"power": 2, "speed": 3, "intellect": 4}, s.User1.Feedbacks)
	assert.Equal(t, 5, len(s.User1.FakeFeedbacks))
	assert.Equal(t, 1279188829, s.User1.FakeFeedbacks["2ATA"])

	assert.Equal(t, "kenakenaf", s.User2.Username)
	assert.Equal(t, "tingly@consubstantial.net", s.User2.Email)
	assert.Equal(t, 3, len(s.User2.Comments))
	assert.Equal(t, "A kitten is a goldfish's peach?", s.User2.Comments[0])
	assert.Equal(t, map[string]int{"power": 2, "speed": 3, "intellect": 4}, s.User2.Feedbacks)
	assert.Equal(t, 7, len(s.User2.FakeFeedbacks))
	assert.Equal(t, 1093938837, s.User2.FakeFeedbacks["5yjaG"])
}

func ExampleBuild() {
	faker.SetSeed(621)
	type Person struct {
		Name   string `faker:"firstName"`
		City   string `faker:"addressCity"`
		Age    int    `faker:"intinrange(0,120)"`
		Code   string
		UUID   string `faker:"uuid;unique"`
		Number int
	}

	p := Person{}
	err := faker.Build(&p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Name)
	fmt.Println(p.City)
	fmt.Println(p.Age)
	fmt.Println(p.Code)
	fmt.Println(p.UUID)
	fmt.Println(p.Number)
	// Output: Elizabeth
	// Thai Nguyen
	// 7
	// QtCMxnMc
	// 566235bc-4211-4ff2-b966-fa2d49d2b167
	// 1267435813
}

func ExampleBuild_second() {
	faker.SetSeed(622)

	// Define a new builder
	colorBuilder := func(params ...string) (interface{}, error) {
		return faker.Pick("Red", "Yellow", "Blue", "Black", "White"), nil
	}

	// Register a new builder named "color" for string type
	err := faker.RegisterBuilder("color", "string", colorBuilder)
	if err != nil {
		panic(err)
	}

	type Animal struct {
		Name  string `faker:"username"`
		Color string `faker:"color"` // Use custom color builder
	}

	type Person struct {
		FirstName string            `faker:"firstName"`         // Any available function case insensitive
		LastName  *string           `faker:"lastName"`          // Pointer are also supported
		Age       int               `faker:"intinrange(0,120)"` // Can call with parameters
		UUID      string            `faker:"uuid;unique"`       // Guarantees a unique value
		Number    int               `faker:"-"`                 // Skip this field
		Code      string            // No tag to use default builder for this field type
		Pet       Animal            // Recursively fill this struct
		Nicknames []string          `faker:"username;len=3"`          // Build an array of size 3 using faker.Username function
		Extra     map[string]string `faker:"stringWithSize(3);len=2"` // map are supported
	}

	p := Person{}
	err = faker.Build(&p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.FirstName)
	fmt.Println(*p.LastName)
	fmt.Println(p.Age)
	fmt.Println(p.UUID)
	fmt.Println(p.Number)
	fmt.Println(p.Code)
	fmt.Println(p.Pet.Name)
	fmt.Println(p.Pet.Color)
	fmt.Println(len(p.Nicknames))
	fmt.Println(p.Nicknames[0])
	fmt.Println(p.Nicknames[1])
	fmt.Println(p.Nicknames[2])
	fmt.Println(p.Extra)
	// Output: Wilber
	// Gutkowski
	// 25
	// ff8d6917-b920-46e6-b1be-dc2d48becfcb
	// 0
	// zN
	// clung
	// Red
	// 3
	// polypeptide
	// chinfest
	// chungchungking
	// map[0w3:F6Y QSi:sq7]
}

func ExampleBuild_third() {
	faker.SetSeed(623)

	type User struct {
		Username string `faker:"username"`
		Email    string `faker:"email"`
		Country  string `faker:"CountryAlpha2"`
	}

	italianUserFactory := func() *User {
		u := &User{Country: "IT"}
		err := faker.Build(u)
		if err != nil {
			panic(err)
		}
		return u
	}

	italianUser := italianUserFactory()
	fmt.Println(italianUser)
	// Output: &{spicule hoag@ornamented.biz IT}
}

func ExampleBuild_fourth() {
	faker.SetSeed(623)

	type User struct {
		Username string `faker:"username"`
		Email    string `faker:"email"`
		Country  string `faker:"CountryAlpha2"`
	}

	// Build 3 users
	users := make([]User, 3)
	err := faker.Build(&users)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", users)
	// Output: [{Username:spicule Email:hoag@ornamented.biz Country:FI} {Username:twelvetone Email:mechanics@ramiform.name Country:TL} {Username:considerate Email:solnit@canonry.biz Country:TG}]
}
