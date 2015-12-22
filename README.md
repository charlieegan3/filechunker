#FileChunker

Divide source code files into 'chunks' of related lines.

#Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/charlieegan3/filechunker"
)

func main() {
	// create a new filechunker, max lines per chunk: 3, using tab as indent
	filechunker := filechunker.NewFileChunker(3, "\t")

	fileString := "var thing string\nfunc main() {\n\tfmt.Println(\"string\")\n}"
	chunks := filechunker.Chunk(fileString)
	fmt.Println(strings.Join(chunks, "\n--\n"))
}
```

Output:

```
var thing string
--
func main() {
	fmt.Println("string")
}
```
