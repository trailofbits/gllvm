package test

import (
	"fmt"
	"testing"

	"github.com/SRI-CSL/gllvm/shared"
)

// gclang -c data/helloworld.c -o helloworld.o
// gclang -c data/testobj.c -o testobj.o
// gclang helloworld.o testobj.o -o helloworld2

func Test_individual_objects(t *testing.T) {
	argss := [][]string{
		{"-c", "../data/helloworld.c", "-o", "helloworld.o"},
		{"-c", "../data/testobj.c", "-o", "testobj.o"},
		{"helloworld.o", "testobj.o", "-o", "../data/helloworld"},
	}

	for _, args := range argss {
		exitCode := shared.Compile(args, "clang")

		if exitCode != 0 {
			t.Errorf("Compile of %v returned %v\n", args, exitCode)
		} else {
			fmt.Println("Compiled OK")
		}

	}

	args := []string{"get-bc", "-v", "../data/helloworld"}

	exitCode := shared.Extract(args)

	if exitCode != 0 {
		t.Errorf("Extraction of %v returned %v\n", args, exitCode)
	} else {
		fmt.Println("Extraction OK")
	}

}
