Solong
======

The solong program prints out all the files in the tree rooted at . whose average line lengths are greater than some long threshold like 1000. The purpose of solong is to detect files that would add noise to search results, so you can add those noisy files to your .ignore or .rgignore file. Having done that, `ripgrep` will ignore these files and you can have a greater signal-to-noise ratio in your search results.

## Installation
```sh
$ go get github.com/ijt/solong
```

## Example usage
```sh
$ solong >> .rgignore
$ ripgrep 'function'
```
