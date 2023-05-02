package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("ERROR: expected exactly one arg - [PORT]\n")
		return
	}

	port := os.Args[1]
	port = strings.TrimPrefix(port, ":")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("ERROR: invalid port %q\n", port)
		os.Exit(1)
	}

	if err := pid(portInt); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func pid(port int) error {

	// -t specifies a terse output that can be piped to the next command
	cmd := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port))
	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return err
	}

	pid := buf.String()
	fmt.Println(pid)

	return nil
}
