package main

import (
	"fmt"
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
	src := srcDir(wd)
	cmd := exec.Command("go", args...)
	if src != "" {
		fmt.Printf("Running with custom GOPATH: %s\n", src)
		env := os.Environ()
		cmd.Env = envWithGopath(src, env)
	} else {
		fmt.Println("Running with default GOPATH.")
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if cmd.Start() != nil {
		log.Fatalln("gop: Unable to start command.")
	}
	_ = cmd.Wait()
}

func srcDir(wd string) string {
	src := wd
	for {
		src = path.Dir(src)
		b := path.Base(src)
		if b == "src" {
			src = path.Dir(src)
			break
		} else if b == "/" {
			src = ""
			break
		}
	}
	return src
}

func envWithGopath(src string, env []string) []string {
	v := "GOPATH=" + src
	found := false
	for i := range env {
		if strings.HasPrefix(env[i], "GOPATH=") {
			env[i] = v
			found = true
			break
		}
	}
	if !found {
		env = append(env, v)
	}
	return env
}
