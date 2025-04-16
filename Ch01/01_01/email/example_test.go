package email

import "fmt"

func ExampleEmail_String() {
	e := Email{"Frodo Baggins", "frodo@shire.org"}
	fmt.Println(e)

	// Output:
	// Frodo Baggins <frodo@shire.org>
}


// OUT:
// Running tool: /usr/local/go/bin/go test -timeout 30s -run ^ExampleEmail_String$ example.com/m
//
// ok  	example.com/m	0.002s
//
// Running tool: /usr/local/go/bin/go test -timeout 30s -run ^ExampleEmail_String$ example.com/m
//
// === RUN   ExampleEmail_String
// --- PASS: ExampleEmail_String (0.00s)
// PASS
// ok      example.com/m   0.002s