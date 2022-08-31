package opt_test

import (
	"testing"

	"adt/opt"
)

func returnNoneOpt() opt.Option[int] {
	return opt.Empty[int]()
}

func returnSomeOpt() opt.Option[int] {
	return opt.WithValue(1)
}

func Test_Type(t *testing.T) {
	opt1 := returnNoneOpt()
	opt2 := returnSomeOpt()

	if opt1.Type() != opt.None {
		t.Fatal("option type was incorrect")
	}

	if opt2.Type() != opt.Some {
		t.Fatal("option type was incorrect")
	}
}

func Test_Value(t *testing.T) {
	opt1 := returnNoneOpt()
	opt2 := returnSomeOpt()

	if opt1.HasValue() != false {
		t.Fatal("option should not have a value")
	}

	if opt2.HasValue() != true {
		t.Fatal("option should have a value")
	}

	value := 3
	opt3 := opt.WithValue(value)

	if opt3.Value() != value {
		t.Fatal("option should have a value")
	}
}

func Test_Unwrap_NoPanic(t *testing.T) {
	option := returnSomeOpt()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not have paniced")
		}
	}()

	_ = option.Unwrap()
}

func Test_Unwrap_Panic(t *testing.T) {
	option := returnNoneOpt()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code should have paniced")
		}
	}()

	_ = option.Unwrap()
}

func Test_ValueOr(t *testing.T) {
	withValue := 1
	opt1 := opt.WithValue(withValue)

	opt2 := opt.Empty[int]()

	if opt1.ValueOr(3) != withValue {
		t.Errorf("unexpected value in ValueOr")
	}

	noValue := 3
	if opt2.ValueOr(noValue) != noValue {
		t.Errorf("unexpected value in ValueOr")
	}
}

func Test_Expect_Panic(t *testing.T) {
	option := returnNoneOpt()
	message := "Test_Expect"
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code should have paniced")
		}

		s := r.(string)
		if s != message {
			t.Errorf("expected different message")
		}
	}()

	_ = option.Expect(message)
}

func Test_Expect_NoPanic(t *testing.T) {
	option := returnSomeOpt()
	message := "Test_Expect"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not have paniced")
		}
	}()

	_ = option.Expect(message)
}

func Test_HasValue(t *testing.T) {
	opt1 := returnNoneOpt()
	opt2 := returnSomeOpt()

	if opt1.HasValue() != false {
		t.Fatal("option should not have a value")
	}

	if opt2.HasValue() != true {
		t.Fatal("option should have a value")
	}
}
