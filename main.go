// Program solong prints out all the files in the tree rooted at .
// that contain no newlines.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	threshold := flag.Int("threshold", 1000, "files with average line length greater than this get identified")
	flag.Parse()

	if err := solong(*threshold); err != nil {
		log.Fatalf("solong: %v", err)
	}
}

func solong(threshold int) error {
	return filepath.Walk(".", func(path string, info fs.FileInfo, e error) error {
		if info.IsDir() {
			if path == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
		nlc := bytes.Count(bs, []byte("\n"))
		n := len(bs)
		if nlc == 0 || n / nlc > threshold {
			fmt.Println(path)
		}
		return nil
	})
}
