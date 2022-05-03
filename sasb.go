package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spenserblack/sasb/pkg/sasb"
)

func main() {
	argc := len(os.Args)
	if argc != 2 {
		fmt.Fprintf(os.Stderr, "Expected 1 arg, got %d\n", argc)
		os.Exit(1)
	}

	s := os.Args[1]

	b, o, d, h := sasb.StringAsBytes(s)

	output := []struct {
		name   string
		values []string
	}{
		{"binary", b},
		{"octal", o},
		{"decimal", d},
		{"hexadecimal", h},
	}

	for _, out := range output {
		values := strings.Join(out.values, ", ")
		label := out.name + ":"
		fmt.Printf("%-12s [%s]\n", label, values)
	}
}
