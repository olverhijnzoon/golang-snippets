package main

import (
	"fmt"
)

type File struct {
	targetDir string
}

func (f *File) GetTargetDir() string {
	return f.targetDir
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Directory List")

	/*
		The intention of this snippet is to deduplicate a list of directories that need to be created. It ensures that only unique directories are added to the "create" map. 
	*/

	fs := []*File{
		&File{targetDir: "/home/user/documents"},
		&File{targetDir: "/usr/local/bin"},
		&File{targetDir: "/var/log"},
	}

	// contains directories that are existing already
	existing := map[string]bool{
		"/usr/local/bin": true,
		"/var/log": true,
	}

	// creating a map called "create" which will store the deduplicated list of directories
	create := map[string]bool{}

	// iterates through each element, "f", in the "fs" list
	for _, f := range fs {
		dir := f.GetTargetDir()
		if existing[dir] {
			continue
		}
		create[dir] = true
	}

	// extracts the keys from the "create" map and adds them to a list called "createList"
	createList := []string{}
	for dir := range create {
		createList = append(createList, dir)
	}

	// deduplicated list of directories
	fmt.Println(createList)
	
}
