package lo

import (
	"errors"
	"fmt"
	"reflect"
)

// LoPanic is user custom panic. example wrap error stack.
var LoPanic = func(e any) { panic(e) }

// LoErrorF is user custom error. example wrap error stack.
var LoErrorF = fmt.Errorf

// Validate is a helper that creates an error when a condition is not met.
// Play: https://go.dev/play/p/vPyh51XpCBt
func Validate(ok bool, format string, args ...any) error {
	if !ok {
		return LoErrorF(fmt.Sprintf(format, args...))
	}
	return nil
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 1 {
		if msgAsStr, ok := msgAndArgs[0].(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msgAndArgs[0])
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// must panics if err is error or false.
func must(err any, messageArgs ...interface{}) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case bool:
		if !e {
			message := messageFromMsgAndArgs(messageArgs...)
			if message == "" {
				message = "not ok"
			}

			LoPanic(message)
		}

	case error:
		message := messageFromMsgAndArgs(messageArgs...)
		if message != "" {
			LoPanic(message + ": " + e.Error())
		} else {
			LoPanic(e.Error())
		}

	default:
		LoPanic("must: invalid err type '" + reflect.TypeOf(err).Name() + "', should either be a bool or an error")
	}
}

// Must is a helper that wraps a call to a function returning a value and an error
// and panics if err is error or false.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must[T any](val T, err any, messageArgs ...interface{}) T {
	must(err, messageArgs...)
	return val
}

// Must0 has the same behavior as Must, but callback returns no variable.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must0(err any, messageArgs ...interface{}) {
	must(err, messageArgs...)
}

// Must1 is an alias to Must
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must1[T any](val T, err any, messageArgs ...interface{}) T {
	return Must(val, err, messageArgs...)
}

// Must2 has the same behavior as Must, but callback returns 2 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must2[T1 any, T2 any](val1 T1, val2 T2, err any, messageArgs ...interface{}) (T1, T2) {
	must(err, messageArgs...)
	return val1, val2
}

// Must3 has the same behavior as Must, but callback returns 3 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must3[T1 any, T2 any, T3 any](val1 T1, val2 T2, val3 T3, err any, messageArgs ...interface{}) (T1, T2, T3) {
	must(err, messageArgs...)
	return val1, val2, val3
}

// Must4 has the same behavior as Must, but callback returns 4 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must4[T1 any, T2 any, T3 any, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err any, messageArgs ...interface{}) (T1, T2, T3, T4) {
	must(err, messageArgs...)
	return val1, val2, val3, val4
}

// Must5 has the same behavior as Must, but callback returns 5 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err any, messageArgs ...interface{}) (T1, T2, T3, T4, T5) {
	must(err, messageArgs...)
	return val1, val2, val3, val4, val5
}

// Must6 has the same behavior as Must, but callback returns 6 variables.
// Play: https://go.dev/play/p/TMoWrRp3DyC
func Must6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err any, messageArgs ...interface{}) (T1, T2, T3, T4, T5, T6) {
	must(err, messageArgs...)
	return val1, val2, val3, val4, val5, val6
}

type wrapErrPrefixMsg struct {
	Base   error
	Attach string
}

func (err wrapErrPrefixMsg) Error() string {
	if err.Attach != "" {
		return err.Attach + ": " + err.Base.Error()
	}
	return err.Base.Error()
}
func (err wrapErrPrefixMsg) String() string {
	return err.Error()
}
func (err wrapErrPrefixMsg) Unwrap() error {
	return err.Base
}

func mustE(err error, messageArgs ...any) {
	if err == nil {
		return
	}
	message := messageFromMsgAndArgs(messageArgs...)
	LoPanic(wrapErrPrefixMsg{Base: err, Attach: message})
}

func MustE[T any](val T, err error, messageArgs ...any) T {
	mustE(err, messageArgs...)
	return val
}

// MustE0 has the same behavior as Must, but callback returns no variable.
func MustE0(err error, messageArgs ...any) {
	mustE(err, messageArgs...)
}

// MustE1 is an alias to MustE
func MustE1[T1 any](val1 T1, err error, messageArgs ...any) T1 {
	mustE(err, messageArgs...)
	return val1
}

// MustE2 has the same behavior as MustE, but callback returns 2 variables.
func MustE2[T1, T2 any](val1 T1, val2 T2, err error, messageArgs ...any) (T1, T2) {
	mustE(err, messageArgs...)
	return val1, val2
}

// MustE3 has the same behavior as MustE, but callback returns 3 variables.
func MustE3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err error, messageArgs ...any) (T1, T2, T3) {
	mustE(err, messageArgs...)
	return val1, val2, val3
}

// MustE4 has the same behavior as MustE, but callback returns 4 variables.
func MustE4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err error, messageArgs ...any) (T1, T2, T3, T4) {
	mustE(err, messageArgs...)
	return val1, val2, val3, val4
}

// MustE5 has the same behavior as MustE, but callback returns 5 variables.
func MustE5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err error, messageArgs ...any) (T1, T2, T3, T4, T5) {
	mustE(err, messageArgs...)
	return val1, val2, val3, val4, val5
}

// MustE6 has the same behavior as MustE, but callback returns 6 variables.
func MustE6[T1, T2, T3, T4, T5, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err error, messageArgs ...any) (T1, T2, T3, T4, T5, T6) {
	mustE(err, messageArgs...)
	return val1, val2, val3, val4, val5, val6
}

// Try calls the function and return false in case of error.
func Try(callback func() error) (ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	err := callback()
	if err != nil {
		ok = false
	}

	return
}

// Try0 has the same behavior as Try, but callback returns no variable.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try0(callback func()) bool {
	return Try(func() error {
		callback()
		return nil
	})
}

// Try1 is an alias to Try.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try1(callback func() error) bool {
	return Try(callback)
}

// Try2 has the same behavior as Try, but callback returns 2 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try2[T any](callback func() (T, error)) bool {
	return Try(func() error {
		_, err := callback()
		return err
	})
}

// Try3 has the same behavior as Try, but callback returns 3 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try3[T, R any](callback func() (T, R, error)) bool {
	return Try(func() error {
		_, _, err := callback()
		return err
	})
}

// Try4 has the same behavior as Try, but callback returns 4 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try4[T, R, S any](callback func() (T, R, S, error)) bool {
	return Try(func() error {
		_, _, _, err := callback()
		return err
	})
}

// Try5 has the same behavior as Try, but callback returns 5 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool {
	return Try(func() error {
		_, _, _, _, err := callback()
		return err
	})
}

// Try6 has the same behavior as Try, but callback returns 6 variables.
// Play: https://go.dev/play/p/mTyyWUvn9u4
func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool {
	return Try(func() error {
		_, _, _, _, _, err := callback()
		return err
	})
}

// TryOr has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr[A any](callback func() (A, error), fallbackA A) (A, bool) {
	return TryOr1(callback, fallbackA)
}

// TryOr1 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr1[A any](callback func() (A, error), fallbackA A) (A, bool) {
	ok := false

	Try0(func() {
		a, err := callback()
		if err == nil {
			fallbackA = a
			ok = true
		}
	})

	return fallbackA, ok
}

// TryOr2 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr2[A any, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (A, B, bool) {
	ok := false

	Try0(func() {
		a, b, err := callback()
		if err == nil {
			fallbackA = a
			fallbackB = b
			ok = true
		}
	})

	return fallbackA, fallbackB, ok
}

// TryOr3 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr3[A any, B any, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (A, B, C, bool) {
	ok := false

	Try0(func() {
		a, b, c, err := callback()
		if err == nil {
			fallbackA = a
			fallbackB = b
			fallbackC = c
			ok = true
		}
	})

	return fallbackA, fallbackB, fallbackC, ok
}

// TryOr4 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr4[A any, B any, C any, D any](callback func() (A, B, C, D, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D) (A, B, C, D, bool) {
	ok := false

	Try0(func() {
		a, b, c, d, err := callback()
		if err == nil {
			fallbackA = a
			fallbackB = b
			fallbackC = c
			fallbackD = d
			ok = true
		}
	})

	return fallbackA, fallbackB, fallbackC, fallbackD, ok
}

// TryOr5 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr5[A any, B any, C any, D any, E any](callback func() (A, B, C, D, E, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E) (A, B, C, D, E, bool) {
	ok := false

	Try0(func() {
		a, b, c, d, e, err := callback()
		if err == nil {
			fallbackA = a
			fallbackB = b
			fallbackC = c
			fallbackD = d
			fallbackE = e
			ok = true
		}
	})

	return fallbackA, fallbackB, fallbackC, fallbackD, fallbackE, ok
}

// TryOr6 has the same behavior as Must, but returns a default value in case of error.
// Play: https://go.dev/play/p/B4F7Wg2Zh9X
func TryOr6[A any, B any, C any, D any, E any, F any](callback func() (A, B, C, D, E, F, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E, fallbackF F) (A, B, C, D, E, F, bool) {
	ok := false

	Try0(func() {
		a, b, c, d, e, f, err := callback()
		if err == nil {
			fallbackA = a
			fallbackB = b
			fallbackC = c
			fallbackD = d
			fallbackE = e
			fallbackF = f
			ok = true
		}
	})

	return fallbackA, fallbackB, fallbackC, fallbackD, fallbackE, fallbackF, ok
}

// TryWithErrorValue has the same behavior as Try, but also returns value passed to panic.
// Play: https://go.dev/play/p/Kc7afQIT2Fs
func TryWithErrorValue(callback func() error) (errorValue any, ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
			errorValue = r
		}
	}()

	err := callback()
	if err != nil {
		ok = false
		errorValue = err
	}

	return
}

// TryCatch has the same behavior as Try, but calls the catch function in case of error.
// Play: https://go.dev/play/p/PnOON-EqBiU
func TryCatch(callback func() error, catch func()) {
	if !Try(callback) {
		catch()
	}
}

// TryCatchWithErrorValue has the same behavior as TryWithErrorValue, but calls the catch function in case of error.
// Play: https://go.dev/play/p/8Pc9gwX_GZO
func TryCatchWithErrorValue(callback func() error, catch func(any)) {
	if err, ok := TryWithErrorValue(callback); !ok {
		catch(err)
	}
}

// ErrorsAs is a shortcut for errors.As(err, &&T).
// Play: https://go.dev/play/p/8wk5rH8UfrE
func ErrorsAs[T error](err error) (T, bool) {
	var t T
	ok := errors.As(err, &t)
	return t, ok
}
