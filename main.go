package main

import (
	// "fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("gop: Provide a Go command.")
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("gop: Unable to get working directory.")
	}
	for {
		wd = path.Dir(wd)
		b := path.Base(wd)
		if b == "src" {
			wd = path.Dir(wd)
			// fmt.Println("GOPATH is now: " + wd )
			break
		} else if b == "/" {
			wd = "$GOPATH"
			// fmt.Println("Using default GOPATH.")
			break
		}
	}
	cmd := exec.Command("go", args...)
	env := os.Environ()
	for i := range env {
		if strings.HasPrefix(env[i], "GOPATH=") {
			env[i] = "GOPATH=" + wd
			break
		}
	}
	cmd.Env = env
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if cmd.Start() != nil {
		log.Fatalln("gop: Unable to start command.")
	}
	_ = cmd.Wait()
}
