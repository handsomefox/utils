package res_test

import (
	"errors"
	"testing"

	"adt/res"
)

func returnErrRes() res.Result[int] {
	return res.WithError[int](errors.New("some error"))
}

func returnValRes(value int) res.Result[int] {
	return res.WithValue(value)
}

func Test_Error(t *testing.T) {
	res1 := returnErrRes()
	res2 := returnValRes(1)

	if res1.Error() == nil {
		t.Fatal("result was expected to have an error")
	}

	if res2.Error() != nil {
		t.Fatal("result was expected to have no errors")
	}
}

func Test_Value(t *testing.T) {
	res1 := returnErrRes()
	res2 := returnValRes(1)

	if res1.IsError() != true {
		t.Fatal("result should have an error")
	}

	if res2.IsError() != false {
		t.Fatal("result should have no errors")
	}

	value := 3
	res3 := res.WithValue(value)

	if res3.Value() != value {
		t.Fatal("result should have a value")
	}
}

func Test_Unwrap_NoPanic(t *testing.T) {
	result := returnValRes(1)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not have panicked")
		}
	}()

	_ = result.Unwrap()
}

func Test_Unwrap_Panic(t *testing.T) {
	result := returnErrRes()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code should have panicked")
		}
	}()

	_ = result.Unwrap()
}

func Test_ValueOr(t *testing.T) {
	withValue := 1
	res1 := res.WithValue(withValue)

	res2 := returnErrRes()

	if res1.ValueOr(3) != withValue {
		t.Errorf("unexpected value in ValueOr")
	}

	noValue := 3
	if res2.ValueOr(noValue) != noValue {
		t.Errorf("unexpected value in ValueOr")
	}
}

func Test_Expect_Panic(t *testing.T) {
	result := returnErrRes()
	message := "Test_Expect"
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code should have panicked")
		}

		s, ok := r.(string)
		if !ok {
			t.Fatalf("no error string found")
		}

		if s != message {
			t.Errorf("expected different message")
		}
	}()

	_ = result.Expect(message)
}

func Test_Expect_NoPanic(t *testing.T) {
	result := returnValRes(1)
	message := "Test_Expect"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not have panicked")
		}
	}()

	_ = result.Expect(message)
}

func Test_IsError(t *testing.T) {
	res1 := returnErrRes()
	res2 := returnValRes(1)

	if res1.IsError() != true {
		t.Fatal("result should have an error")
	}

	if res2.IsError() != false {
		t.Fatal("result should have no errors")
	}
}
