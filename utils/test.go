package utils

import "testing"

type TestResult interface {
}

func CheckResult(t *testing.T, name string, result TestResult, expected TestResult) {
	if result == expected {
		t.Logf("\t\tPASSED: %v %v, expected %v", name, result, expected)
	} else {
		t.Fatalf("\t\tFAILED: %v %v, expected %v", name, result, expected)
	}
}
