package main

import (
	"testing"
)

func TestSrcDir(t *testing.T) {
	tests := []struct {
		wd       string
		expected string
	}{
		{wd: "/a/b/c", expected: ""},
		{wd: "/src/b/c", expected: "/"},
		{wd: "/a/b/src/c/d", expected: "/a/b"},
	}

	for _, test := range tests {
		actual := srcDir(test.wd)
		if actual != test.expected {
			t.Fatalf("Actual: %s, Expected: %s\n", actual, test.expected)
		}
	}
}

func TestEnvWithGopath(t *testing.T) {
	tests := []struct {
		src      string
		env      []string
		expected []string
	}{
		{src: "/a/b", env: []string{"ONE=one", "TWO=two"}, expected: []string{"ONE=one", "TWO=two", "GOPATH=/a/b"}},
		{src: "/a/b", env: []string{"ONE=one", "TWO=two", "GOPATH=/c/d"}, expected: []string{"ONE=one", "TWO=two", "GOPATH=/a/b"}},
	}

	for _, test := range tests {
		actual := envWithGopath(test.src, test.env)
		for i, env := range actual {
			if env != test.expected[i] {
				t.Fatalf("Actual: %s, Expected: %s\n", env, test.expected[i])
			}
		}
	}
}
