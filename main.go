package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	_cmdpath := "/home/street/Development/go/src/file_exec/echo.go"

	output, err := exec.Command("go", "run", _cmdpath).CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else {
		log.Printf("Executed successfully")
	}

}
