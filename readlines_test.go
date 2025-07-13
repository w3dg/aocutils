package aocutils

import "testing"

func TestReadFileLines(t *testing.T) {
	t.Run("empty file name", func(t *testing.T) {
		f := ""
		_, err := ReadFileLines(f)
		if err == nil {
			t.Errorf("Empty file name passed to ReadFileLines should not pass filename: %v", f)
		}
	})

	t.Run("non existent file name", func(t *testing.T) {
		f := "nonexistent.txt"
		_, err := ReadFileLines(f)
		if err == nil {
			t.Errorf("Empty file name passed to ReadFileLines should not pass filename: %v", f)
		}
	})

	t.Run("existent file", func(t *testing.T) {
		f := "testdata/empty.txt"
		_, err := ReadFileLines(f)
		if err != nil {
			t.Errorf("Erorred for a valid present filename: %v", f)
		}
	})

	t.Run("empty file", func(t *testing.T) {
		f := "testdata/empty.txt"
		lines, _ := ReadFileLines(f)

		if len(lines) != 0 {
			t.Errorf("Expected empty file to return an empty slice of lines got %v want %v", len(lines), 0)
		}
	})

	t.Run("a file with 10 lines no blocks", func(t *testing.T) {
		f := "testdata/without-blocks.txt"
		lines, _ := ReadFileLines(f)

		want := 10
		if len(lines) != want {
			t.Errorf("Wrong number of lines got %v want %v", len(lines), want)
		}
	})

	t.Run("a file with 11 lines two blocks but parsed with ReadFileLines", func(t *testing.T) {
		f := "testdata/with-blocks.txt"
		lines, _ := ReadFileLines(f)
		want := 11
		if len(lines) != want {
			t.Errorf("Wrong number of lines got %v want %v", len(lines), want)
		}
	})

	t.Run("a file with 11 lines with two blocks parsed by ReadFileBlocks", func(t *testing.T) {
		f := "testdata/with-blocks.txt"
		blocks, _ := ReadFileBlocks(f)
		want := 2
		if len(blocks) != want {
			t.Errorf("Wrong number of blocks got %v want %v", len(blocks), want)
		}
	})

}
