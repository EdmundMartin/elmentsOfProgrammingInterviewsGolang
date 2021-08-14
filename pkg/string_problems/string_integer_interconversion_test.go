package string_problems

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	result := IntToString(312)
	if result != "312" {
		t.Errorf("expected '312', got: %s", result)
	}
	result = IntToString(-123)
	if result != "-123" {
		t.Errorf("expected -123, got: %s", result)
	}
}

func TestStrToInt(t *testing.T) {
	result := StrToInt("312")
	if result != 312 {
		t.Errorf("expected '312', got: %d", result)
	}

	result = StrToInt("-123")
	if result != -123 {
		t.Errorf("expected -123, got: %d", result)
	}
}
