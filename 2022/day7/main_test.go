package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestSolutionPart1(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	root := parse(file)
	assert.Equal(t, 1432936, sum10kOrLessDirectories(root))
}

func TestSolutionPart2(t *testing.T) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	root := parse(file)
	assert.Equal(t, 272298, smallestDirToFreeEnoughSpace(root))
}

func TestExamplePart1(t *testing.T) {
	root := parse(strings.NewReader(example))
	assert.Equal(t, 95437, sum10kOrLessDirectories(root))
}

func TestExamplePart2(t *testing.T) {
	root := parse(strings.NewReader(example))
	assert.Equal(t, 24933642, smallestDirToFreeEnoughSpace(root))
}

func TestParse_adds_directory_sizes_to_parent_directories(t *testing.T) {
	root := parse(strings.NewReader(example))
	dirA := root.Items[0]
	assert.Equal(t, "a", dirA.Name)
	assert.Equal(t, 94853, dirA.CalculateSize())
}

func TestParse_adds_sizes_to_parent_directory(t *testing.T) {
	root := parse(strings.NewReader(example))
	dirD := root.Items[3]
	assert.Equal(t, "d", dirD.Name)
	assert.Equal(t, 24933642, dirD.CalculateSize())
}

func TestParse_can_go_up_directories_with_double_period_notation(t *testing.T) {
	root := parse(strings.NewReader(example))
	dirD := root.Items[3]
	assert.Equal(t, "d", dirD.Name)
	assert.Len(t, dirD.Items, 4)
	assert.Equal(t, "j", dirD.Items[0].Name)
	assert.Equal(t, 4060174, dirD.Items[0].Size)
}

func TestParse_parses_sub_directories(t *testing.T) {
	root := parse(strings.NewReader(example))
	dirA := root.Items[0]
	assert.Len(t, dirA.Items, 4)

	dirE := dirA.Items[0]
	assert.Equal(t, "e", dirE.Name)
	assert.Equal(t, "i", dirE.Items[0].Name)
	assert.Equal(t, 584, dirE.Items[0].Size)
}

func TestParse_parses_root_level_files_and_directories(t *testing.T) {
	root := parse(strings.NewReader(example))
	assert.Equal(t, root.Name, "/")
	assert.Equal(t, "a", root.Items[0].Name)
	bTxt := root.Items[1]
	assert.Equal(t, "b.txt", bTxt.Name)
	assert.Equal(t, 14848514, bTxt.Size)
	cDat := root.Items[2]
	assert.Equal(t, "c.dat", cDat.Name)
	assert.Equal(t, 8504156, cDat.Size)
	assert.Equal(t, "d", root.Items[3].Name)
}
