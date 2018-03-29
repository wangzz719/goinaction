package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprint(&b, "World!\n")
	b.WriteTo(os.Stdout)
}
