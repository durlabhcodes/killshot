package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	procName := flag.String("n", "", "Process name or substring of name you want to kill")

	flag.Parse()

	fmt.Printf("Searching process with name containing %s\n", *procName)

	//TODO - Will be using later
	//	operatingSystem := runtime.GOOS

	//output, cmdError := exec.Command("ps aux | grep " + *procName).Output()
	psCommand := exec.Command("ps", "aux")
	grepCommand := exec.Command("grep", *procName)

	grepCommand.Stdin, _ = psCommand.StdoutPipe()
	grepCommand.Stdout = os.Stdout

	psCommand.Start()
	grepCommand.Start()

	psCommand.Wait()
	grepCommand.Wait()
	//	fmt.Println(string(grepCommandPipe))
}
