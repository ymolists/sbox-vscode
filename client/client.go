package main

import (
	"fmt"
	"os"

	"github.com/ymolists/sbox-vscode/server"
)

func main() {

	fmt.Fprintf(os.Stderr, "%s\n", server.ServerFunc())
	fmt.Fprintf(os.Stderr, "%s\n", server.ServerNew00())
}
