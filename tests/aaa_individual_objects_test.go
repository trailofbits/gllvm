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
		{"-c", "../data/helloworld.c", "-o", "helloworld2.o"},
		{"-c", "../data/testobj.c", "-o", "testobj2.o"},
		{"helloworld2.o", "testobj2.o", "-o", "../data/hello2"},
	}

	for _, args := range argss {
		exitCode := shared.Compile(args, "clang")

		if exitCode != 0 {
			t.Errorf("Compile of %v returned %v\n", args, exitCode)
		} else {
			fmt.Println("Compiled OK")
		}

	}

	args := []string{"get-bc", "-v", "../data/hello2"}

	exitCode := shared.Extract(args)

	if exitCode != 0 {
		t.Errorf("Extraction of %v returned %v\n", args, exitCode)
	} else {
		fmt.Println("Extraction OK")
	}

}
