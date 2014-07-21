Gop
===

Gop is a GOPATH resolution tool. It allows you to execute Go commands with local GOPATH resolution.

For example, if your directory structure looks like this:

```sh
project
	src
		github.com
			username
				project
					subpackage
						subpackage.go
					main.go
```

The following scenarios would work:
- gop install (inside of project/github.com/username/project)
- gop install project/ (inside of project/github.com/username)