package main

import (
	"bufio"
	"fmt"
	"os"
)


// Exercise 1.4
func main() {
        counts := make(map[string]int)
        files := os.Args[1:]
        if len(files) == 0 {
             countLines(os.Stdin, counts)
        } else {
            for _, arg := range files {
                f, err := os.Open(arg)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}
                countLines(f, counts)
				f.Close()
				count := 0
		        for _, n := range counts {
		            if n > 1 {
		                count++
					} 
				}
				if count > 0 {
					fmt.Printf("File %s has %d dumplicate lines!\n", arg, count)
				}
				counts = make(map[string]int)
			}
        }
}	
     
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}