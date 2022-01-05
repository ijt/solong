// Program solong prints out the paths of all files in the tree rooted at .
// whose average line lengths are greater than some long threshold
// like 1000. The purpose of solong is to detect files that would
// add noise to search results, so you can add them to your .ignore
// or .rgignore file.
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
		if nlc == 0 || n/nlc > threshold {
			fmt.Println(path)
		}
		return nil
	})
}
