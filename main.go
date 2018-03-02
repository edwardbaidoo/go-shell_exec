package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	_cmdpath := "echo.go" //path to file to be executed

	//if var _cmdpath isnt used, it can be commented out or removed.
	output, err := exec.Command("go", "run", _cmdpath).CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else {
		log.Printf("Executed successfully")
	}

}
