package wc

import (
	"fmt"
	"os"
	"strings"
)

var data = `Baa, baa, black sheep,
Have you any wool?
Yes, sir, yes, sir,
Three bags full.
`

func ExampleLineCount() {
	r := strings.NewReader(data)
	fmt.Println(LineCount(r))

	file, err := os.Open("testdata/poem.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer file.Close()
	fmt.Println(LineCount(file))

	// Output:
	// 4 <nil>
	// 4 <nil>
}


// OUT:
// Running tool: /usr/local/go/bin/go test -timeout 30s -run ^ExampleLineCount$ example.com/m
//
// === RUN   ExampleLineCount
// --- PASS: ExampleLineCount (0.00s)
// PASS
// ok      example.com/m   0.003s