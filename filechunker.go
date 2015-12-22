package filechunker

import (
	"bytes"
	"strings"
)

type FileChunker struct {
	maxLength         int
	indentationString string
}

func (f *FileChunker) Chunk(file string) []string {
	lines := strings.Split(file, "\n")
	var chunks []string
	var chunkBuffer bytes.Buffer
	for i, line := range lines {
		line = formatLine(line, f.indentationString)
		if i < len(lines)-1 && isIndented(lines[i+1]) || isIndented(line) {
			chunkBuffer.WriteString(line + "\n")
			continue
		}
		chunk := line
		if chunkBuffer.Len() > 0 {
			chunk = chunkBuffer.String() + line
			chunkBuffer.Reset()
		}
		if isValid(chunk, f.maxLength) {
			chunks = append(chunks, chunk)
		}
	}

	if chunkBuffer.Len() > 0 {
		chunks = append(chunks, chunkBuffer.String()[0:chunkBuffer.Len()-1])
	}
	return chunks
}

func isIndented(line string) bool {
	if len(line) == 0 {
		return false
	} else if line[0] == '\t' || line[0] == ' ' {
		return true
	} else {
		return false
	}
}

func isValid(chunk string, maxLength int) bool {
	if chunk == "" {
		return false
	} else if len(strings.Split(chunk, "\n")) > maxLength {
		return false
	} else {
		return true
	}
}

func formatLine(line string, indentationString string) string {
	if isIndented(line) {
		unindentedLine := strings.TrimLeft(line, " ")
		unindentedLine = strings.TrimLeft(unindentedLine, "\t")
		lengthDifference := len(line) - len(unindentedLine)

		var newLine bytes.Buffer
		for i := 0; i < lengthDifference; i++ {
			newLine.WriteString(indentationString)
		}
		newLine.WriteString(unindentedLine)
		line = newLine.String()
	}
	line = strings.TrimRight(line, " ")
	line = strings.TrimRight(line, "\t")
	return line
}

func NewFileChunker(maxLength int, indentationString string) FileChunker {
	return FileChunker{maxLength: maxLength, indentationString: indentationString}
}
