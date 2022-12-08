package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func smallestDirToFreeEnoughSpace(root File) int {
	total := 70000000
	needed := 30000000
	used := root.CalculateSize()
	unused := total - used
	toDelete := needed - unused
	fmt.Printf("To Delete: \n%d\n", toDelete)
	min := 100000000000000000

	return minimalDirSize(root, toDelete, min)
}

func minimalDirSize(root File, toDelete, min int) int {
	for _, f := range root.Items {
		if !f.Dir {
			continue
		}
		size := f.CalculateSize()
		if size >= toDelete && size <= min {
			min = size
		}
		min = minimalDirSize(*f, toDelete, min)
	}

	return min
}

func sum10kOrLessDirectories(root File) (sum int) {
	for _, file := range root.Items {
		if !file.Dir {
			continue
		}
		if file.CalculateSize() <= 100_000 {
			sum += file.CalculateSize()
		}
		sum += sum10kOrLessDirectories(*file)
	}

	return
}

type File struct {
	Name   string
	Items  []*File
	Size   int
	Parent *File
	Dir    bool
}

func (this *File) CalculateSize() (size int) {
	for _, f := range this.Items {
		if f.Dir {
			size += f.CalculateSize()
		} else {
			size += f.Size
		}
	}
	return
}

func parse(r io.Reader) (root File) {
	scanner := bufio.NewScanner(r)
	root.Name = "/"
	root.Dir = true
	var curDir = &root
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			continue
		}
		if strings.Contains(line, "$ cd ") {
			var newDir File
			fmt.Sscanf(line, "$ cd %s", &newDir.Name)
			if newDir.Name == ".." {
				curDir = curDir.Parent
				continue
			}
			for _, dir := range curDir.Items {
				// Find the cd #dir in the current directory's items.
				if dir.Name == newDir.Name {
					curDir = dir
					break
				}
			}

			continue
		}
		if strings.Contains(line, "dir ") {
			var dir File
			dir.Parent = curDir
			dir.Dir = true
			fmt.Sscanf(line, "dir %s", &dir.Name)
			curDir.Items = append(curDir.Items, &dir)
			continue
		}

		if !strings.Contains(line, "$ ls") {
			var file File
			file.Parent = curDir
			fmt.Sscanf(line, "%d %s", &file.Size, &file.Name)
			curDir.Items = append(curDir.Items, &file)
		}
	}

	return root
}
