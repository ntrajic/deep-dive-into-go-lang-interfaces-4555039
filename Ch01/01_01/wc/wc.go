package wc

import (
	"bufio"
	"io"
)

// LineCount return how many lines in r.
func LineCount(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		count++
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// OUT:
// go mod init <enter>   ---> help
// go mod init example.com/m
// go mod tidy
// You can specify which packages to build by providing their import paths as arguments to go build. For example:
// go build . (builds the package in the current directory)
// go build ./mypackage (builds the package in the mypackage subdirectory)
// go build example.com/m/mypackage (builds a package within your module)
//
// Output Name and Location:
//
// The -o flag allows you to specify the name and location of the output executable:
// go build -o myapp (creates an executable named myapp in the current directory)
// go build -o bin/myapp (creates an executable named myapp in the bin subdirectory)
//
// Building for Different Platforms:
//
// You can use environment variables like GOOS (operating system) and GOARCH (architecture) to cross-compile your Go code for different platforms:
// GOOS=linux GOARCH=amd64 go build -o myapp_linux
// Build Tags:
//
// Build tags are special comments in your Go source files that allow you to include or exclude code during the build process based on certain conditions. You can specify build tags using the -tags flag:
// go build -tags debug (includes code marked with the // +build debug tag)