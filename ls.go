package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// 실행할 Linux 명령어와 인자
	cmd := exec.Command("ls", "-l")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Success: %s\n", stderr.String())
		return
	}

	fmt.Printf("Internal Server Error:\n%s\n", out.String())
}
