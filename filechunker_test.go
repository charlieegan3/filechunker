package filechunker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	expected := FileChunker{maxLength: 3, indentationString: "\t"}
	assert.Equal(t, expected, filechunker, "they should be equal")
}

func TestBasicFile(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1\nline2\nline3"
	expected := []string{"line1", "line2", "line3"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestBlankLinesAreIgnored(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1\n\nline2"
	expected := []string{"line1", "line2"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestIndentedBlockGrouping(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "var = 1\ndef method\n\treturn value\nend\nvar = 1"
	expected := []string{"var = 1", "def method\n\treturn value\nend", "var = 1"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestLinesAreStripped(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1   \nline2"
	expected := []string{"line1", "line2"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestWhitespaceIsNormalized(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1\n\tline2\nline3"
	file2 := "line1\n line2\nline3"

	assert.Equal(t, filechunker.Chunk(file), filechunker.Chunk(file2), "they should be equal")

	file = "line1\n line2\nline3"
	expected := []string{"line1\n\tline2\nline3"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestIncompleteBlocksAreIncluded(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1\n\tline2"
	expected := []string{"line1\n\tline2"}

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}

func TestMaxLengthIsEnforced(t *testing.T) {
	filechunker := NewFileChunker(3, "\t")
	file := "line1\n\tline2\n\tline3\nline4"
	expected := []string(nil)

	assert.Equal(t, expected, filechunker.Chunk(file), "they should be equal")
}
