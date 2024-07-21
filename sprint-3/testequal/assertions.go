//go:build !solution

package testequal

import "reflect"

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if areEqual(expected, actual) {
		return true
	}
	recordMsg(t, msgAndArgs...)
	return false
}

func areEqual(expected, actual interface{}) bool {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return false
	}
	switch expected.(type) {
	case int:
		return expected.(int) == actual.(int)
	case int8:
		return expected.(int8) == actual.(int8)
	case int16:
		return expected.(int16) == actual.(int16)
	case int32:
		return expected.(int32) == actual.(int32)
	case int64:
		return expected.(int64) == actual.(int64)
	case uint:
		return expected.(uint) == actual.(uint)
	case uint8:
		return expected.(uint8) == actual.(uint8)
	case uint16:
		return expected.(uint16) == actual.(uint16)
	case uint32:
		return expected.(uint32) == actual.(uint32)
	case uint64:
		return expected.(uint64) == actual.(uint64)
	case string:
		return expected.(string) == actual.(string)
	case map[string]string:
		return mapsAreEqual(expected.(map[string]string), actual.(map[string]string))
	case []int:
		return intSlicesAreEqual(expected.([]int), actual.([]int))
	case []byte:
		return byteSlicesAreEqual(expected.([]byte), actual.([]byte))
	}
	return false
}

func mapsAreEqual(m1, m2 map[string]string) bool {
	if m1 == nil && m2 != nil || m1 != nil && m2 == nil {
		return false
	}
	if len(m1) != len(m2) {
		return false
	}
	for i, str := range m1 {
		if v, ok := m2[i]; !ok || str != v {
			return false
		}
	}
	return true
}

func intSlicesAreEqual(s1, s2 []int) bool {
	if s1 == nil && s2 != nil || s1 != nil && s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i, num := range s1 {
		if num != s2[i] {
			return false
		}
	}
	return true
}

func byteSlicesAreEqual(s1, s2 []byte) bool {
	if s1 == nil && s2 != nil || s1 != nil && s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i, b := range s1 {
		if b != s2[i] {
			return false
		}
	}
	return true
}

func recordMsg(t T, msgAndArgs ...interface{}) {
	t.Helper()
	if len(msgAndArgs) == 0 {
		t.Errorf("")
		return
	}
	msg, ok := msgAndArgs[0].(string)
	if ok {
		t.Errorf(msg, msgAndArgs[1:]...)
	} else {
		panic("first argument must be a format string")
	}
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !areEqual(expected, actual) {
		return true
	}
	recordMsg(t, msgAndArgs...)
	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if areEqual(expected, actual) {
		return
	}
	recordMsg(t, msgAndArgs...)
	t.FailNow()
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !areEqual(expected, actual) {
		return
	}
	recordMsg(t, msgAndArgs...)
	t.FailNow()
}
