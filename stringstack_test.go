package aocutils

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewStringStack(t *testing.T) {
	st := NewStringStack()

	if st == nil {
		t.Fatal("Initialized Stack should not be nil")
	}

	if st.Len() != 0 {
		t.Errorf("New Stack length should be zero got: %v want %v", st.Len(), 0)
	}
}

func TestNewPopulatedStringStack(t *testing.T) {
	elements := []string{"1", "2", "3", "4", "5"}
	st := NewPopulatedStringStack(elements)

	if st == nil {
		t.Fatal("Expected stack to not be nil")
	}
	want := len(elements)
	if st.Len() != want {
		t.Errorf("Populated Stack does not have same number of elements as the number passed to it got: %v want %v", st.Len(), want)
	}
}

func TestPush(t *testing.T) {
	st := NewStringStack()

	st.Push("1")
	st.Push("2")
	st.Push("3")

	want := 3
	if st.Len() != want {
		t.Errorf("Push operation on stack did not grow its length correctly got: %v want: %v", st.Len(), want)
	}
}

func TestPop(t *testing.T) {
	elements := []string{"1", "2", "3", "4", "5"}
	st := NewPopulatedStringStack(elements)

	for i := len(elements) - 1; i >= 0; i-- {
		popped, err := st.Pop()
		if err != nil {
			t.Error("Encountered error while popping elements off a valid stack")
		}

		if popped != elements[i] {
			t.Errorf("Got incorrect value after popping the stack got: %v want %v", popped, elements[i])
		}
	}

	_, err := st.Pop()

	if err == nil {
		t.Errorf("Did not error out on a stack underflow condition")
	}
}

func TestPeek(t *testing.T) {
	elements := []string{"1", "2", "3", "4", "5"}

	st := NewPopulatedStringStack(elements)

	for i := len(elements) - 1; i >= 0; i-- {
		want := elements[i]

		v, err := st.Peek()
		if err != nil {
			t.Errorf("Got underflow error from a valid stack peek")
		}
		if v != want {
			t.Errorf("Got wrong value of top element while peeking got: %v, want %v", v, want)
		}

		st.Pop()
	}

	_, err := st.Peek()
	if err == nil {
		t.Errorf("Expected error while peeking on empty stack, got nil")
	}
}

func TestEmpty(t *testing.T) {
	populated := NewPopulatedStringStack([]string{"1", "2", "3", "4", "5"})
	empty := NewStringStack()

	if populated.IsEmpty() {
		t.Errorf("Got populated stack to be said to be empty")
	}

	if !empty.IsEmpty() {
		t.Errorf("Got an empty stack to be said to be not empty")
	}
}

func TestString(t *testing.T) {
	populated := NewPopulatedStringStack([]string{"1", "2", "3", "4", "5"})
	empty := NewStringStack()

	populatedPrintOutput := fmt.Sprint(populated)
	emptyPrintOutput := fmt.Sprint(empty)

	want := "StringStack: empty"
	if emptyPrintOutput != want {
		t.Errorf("Wrong string output for empty StringStack got %v, want %v", emptyPrintOutput, want)
	}

	wantSubstr := "StringStack: "
	if !strings.HasPrefix(populatedPrintOutput, wantSubstr) {
		t.Log(populatedPrintOutput)
		t.Errorf("Wrong string output for populated StringStack. Does not start with prefix '%v'", wantSubstr)
	}
}
