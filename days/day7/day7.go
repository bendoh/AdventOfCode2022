package day7

import (
	"fmt"
	"math"
)

type Dir struct {
	name   string
	dirs   []*Dir
	parent *Dir
	files  map[string]int
}

func makeDir(name string, parent *Dir) *Dir {
	dir := Dir{name: name}
	dir.files = make(map[string]int)
	dir.dirs = make([]*Dir, 0)
	dir.parent = parent

	return &dir
}

var dirSizes map[string]int

func findDirSizes(dir *Dir, path string) int {
	dirSize := 0

	for _, size := range dir.files {
		dirSize += size
	}

	for _, subdir := range dir.dirs {
		dirSize += findDirSizes(subdir, path+subdir.name+"/")
	}

	dirSizes[path] = dirSize

	return dirSize

}
func Day7(input []string) []string {
	listing := false

	dirSizes = make(map[string]int)
	fs := makeDir("/", nil)
	curdir := fs

	for _, line := range input[1:] {
		if len(line) == 0 {
			continue
		}
		if line[0:4] == "$ cd" {
			listing = false
			dest := line[5:]

			if dest == "/" {
				curdir = fs
			} else if dest == ".." {
				curdir = curdir.parent
			} else {
				found := -1
				for idx, subdir := range curdir.dirs {
					if subdir.name == dest {
						found = idx
						break
					}
				}
				if found >= 0 {
					curdir = curdir.dirs[found]
				}
			}
			continue
		} else if line[0:4] == "$ ls" {
			listing = true
			continue
		}

		if listing {
			if line[0:3] == "dir" {
				(*curdir).dirs = append(curdir.dirs, makeDir(line[4:], curdir))
			} else {
				var size int
				var name string
				fmt.Sscanf(line, "%d %s", &size, &name)

				curdir.files[name] = size
			}
		}
	}

	findDirSizes(fs, "/")

	diskSize := 70000000
	need := 30000000
	smallest := math.MaxInt

	unusedSpace := diskSize - dirSizes["/"]
	sum := 0
	winner := 0

	for _, size := range dirSizes {
		if size <= 100000 {
			sum += size
		}

		if unusedSpace+size > need {
			if size < smallest {
				smallest = size
				winner = size
			}
		}
	}

	return []string{fmt.Sprintf("Sum: %d; Winner: %d", sum, winner)}
}
