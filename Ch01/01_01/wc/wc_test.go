package wc

import (
	"strings"
	"testing"
	"testing/iotest"
)

var testData = `
One for my master,
One for my dame,
And one for the little boy
Who lives down the lane.
`

func TestLineCountError(t *testing.T) {
	r := iotest.TimeoutReader(strings.NewReader(testData))
	_, err := LineCount(r)
	if err == nil {
		t.Fatal("no error on read")
	}
}


// OUT:
// Running tool: /usr/local/go/bin/go test -timeout 30s -run ^TestLineCountError$ example.com/m
//
// ok  	example.com/m	(cached)

// Running tool: /usr/local/go/bin/go test -timeout 30s -run ^TestLineCountError$ example.com/m
//
// === RUN   TestLineCountError
// --- PASS: TestLineCountError (0.00s)
// PASS
// ok      example.com/m   0.002s
