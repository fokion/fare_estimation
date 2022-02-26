package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestFlags(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	extractedFile := fmt.Sprintf("output_%s.txt", time.Now().Format("06-01-02_15_04"))
	//pattern that I am following https://mj-go.in/golang/test-the-main-function-in-go
	cases := []struct {
		Name         string
		Args         []string
		ExpectedExit int
	}{
		{"flags set", []string{"-f", "./paths.csv"}, 0},
		{"flag missing", []string{}, 1},
	}
	for _, testCase := range cases {
		flag.CommandLine = flag.NewFlagSet(testCase.Name, flag.ExitOnError)
		os.Args = append([]string{testCase.Name}, testCase.Args...)
		assert.Equal(t, testCase.ExpectedExit, realMain())
	}

	defer func() {
		_, err := os.Stat(extractedFile)
		if err == nil {
			assert.NoError(t, os.Remove(extractedFile))
		} else {
			assert.Fail(t, "Missing the expected file from at one successful call")
		}
	}()

}
