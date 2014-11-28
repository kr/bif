package interp

import (
	"bufio"
	"io"
	"os"
)

func prompt(s string, scanner *bufio.Scanner) bool {
	io.WriteString(os.Stdout, s)
	return scanner.Scan()
}
