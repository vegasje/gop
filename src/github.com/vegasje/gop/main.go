package main

import (
	// "fmt"
	"log"
	"os"
	"os/exec"
	"path"
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
			break
		} else if b == "/" {
			wd = "$GOPATH"
			break
		}
	}
	// for i := range env {
	// 	fmt.Println(env[i])
	// }
	env := append(os.Environ(), "GOPATH="+wd)
	cmd := exec.Command("go", args...)
	cmd.Env = env
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if cmd.Start() != nil {
		log.Fatalln("gop: Unable to start command.")
	}
	_ = cmd.Wait()
}
