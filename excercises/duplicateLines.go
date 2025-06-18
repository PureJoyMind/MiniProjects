package excercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// These are small methods each teaching me something about the Go language

func Dup1() {
	// Dup1 prints the text of each line that appears more than
	// once in the standard input, preceded by its count.

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup2() {
	// Dup2 prints the count and text of lines that appear more than once
	// in the input. It reads from stdin or from a list of named files.
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) >= 1 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			CountLines(f, counts)
			f.Close()
		}
	} else {
		CountLines(os.Stdin, counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func CountLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func Dup3() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) >= 1 {
		for _, f := range files {
			d, e := os.ReadFile(f)
			if e != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", e)
				continue
			}
			for line := range strings.SplitSeq(string(d), "\n") {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
