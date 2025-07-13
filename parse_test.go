package aocutils

import "testing"

func TestParseNum(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		number, err := ParseNum("42")
		if err != nil {
			t.Errorf("Got error for valid string for conversion")
		}
		if number != 42 {
			t.Errorf("Got wrong converted number got: %v, want %v", number, 42)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := ParseNum("abcd")

		if err == nil {
			t.Errorf("Should have errored for invalid input, instead did not")
		}
	})
}

func TestParseNumOrPanic(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		result := ParseNumOrPanic("123")
		if result != 123 {
			t.Errorf("ParseNumOrPanic(\"123\") = %d; want 123", result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("ParseNumOrPanic(\"abc\") did not panic")
			}
		}()

		// This should cause a panic
		ParseNumOrPanic("abc")

		// If we reach here, no panic occurred
		t.Errorf("Expected panic, but function returned normally")
	})
}
