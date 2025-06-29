// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files_name := make(map[string]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, files_name)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, files_name)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            files_repeted := files_name[line]
            fmt.Printf("%d\t%s\tfiles: %s\n", n, line, files_repeted)
        }
    }
}

func countLines(f *os.File, counts map[string]int, files_name map[string]string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        files_name[input.Text()] = files_name[input.Text()] + f.Name() + " "
    }
    // NOTE: ignoring potential errors from input.Err()
}