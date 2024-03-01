package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	const version = 1

	st := NewState()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")
		slog.Info("process line", "line", line)

		st = ParseCommand(line).Apply(st)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	slog.Info("parsed state", "state", st)

	fmt.Println(st.SumAllDirWithMaxSize(100000))

	unusedSpace := 70000000 - st.root.size
	requiredUnusedSpace := 30000000
	fmt.Println(st.FindSmallestDirSizeBiggerThan(requiredUnusedSpace - unusedSpace))

	slog.Info("done")
}

type State struct {
	root *FileNode
	curr *FileNode
}

func NewState() *State {
	root := NewFileTree()
	return &State{
		root: root,
		curr: root,
	}
}

func (s *State) SumAllDirWithMaxSize(size int) int {
	return s.root.SumAllDirWithMaxSize(size)
}

func (s *State) FindSmallestDirSizeBiggerThan(size int) int {
	return s.root.FindSmallestDirSizeBiggerThan(size)
}

func (s *State) String() string {
	return fmt.Sprintf("State{root: %v}", s.root)
}

type Command interface {
	Apply(s *State) *State
}

func ParseCommand(line string) Command {
	if !strings.HasPrefix(line, "$") {
		firstPart := strings.Split(line, " ")[0]
		firstPartAsInt, _ := strconv.Atoi(firstPart)
		return fileinfoCommand{isDir: firstPart == "dir", size: firstPartAsInt}
	}

	line = strings.TrimPrefix(line, "$ ")
	if strings.HasPrefix(line, "cd") {
		return cdCommand{arg: strings.TrimPrefix(line, "cd ")}
	}
	return new(nullCommand)
}

type cdCommand struct {
	arg string
}

func (c cdCommand) Apply(s *State) *State {
	switch c.arg {
	case "/":
		s.curr = s.root
	case "..":
		s.curr = s.curr.parent
	default:
		s.curr = s.curr.AddChild(c.arg)
	}
	return s
}

type fileinfoCommand struct {
	isDir bool
	size  int
}

func (c fileinfoCommand) Apply(s *State) *State {
	if c.size > 0 {
		s.curr.AddSize(c.size)
	}
	return s
}

type nullCommand struct{}

func (nullCommand) Apply(s *State) *State {
	return s
}

type FileNode struct {
	parent   *FileNode
	children map[string]*FileNode
	size     int
}

func NewFileTree() *FileNode {
	return &FileNode{parent: nil, children: make(map[string]*FileNode)}
}

func (f *FileNode) AddChild(name string) *FileNode {
	child := &FileNode{parent: f, children: make(map[string]*FileNode)}
	f.children[name] = child
	return child
}

func (f *FileNode) AddSize(size int) {
	f.size += size
	for f.parent != nil {
		f.parent.size += size
		f = f.parent
	}
}

func (f *FileNode) IsDir() bool {
	return len(f.children) > 0
}

func (f *FileNode) SumAllDirWithMaxSize(size int) int {
	sum := 0
	if f.size <= size {
		sum = f.size
	}

	for _, child := range f.children {
		sum += child.SumAllDirWithMaxSize(size)
	}
	return sum
}

func (f *FileNode) FindSmallestDirSizeBiggerThan(size int) int {
	if f.size < size {
		return -1
	}

	candidate := f.size
	for _, child := range f.children {
		if s := child.FindSmallestDirSizeBiggerThan(size); s > 0 {
			if s < candidate {
				candidate = s
			}
		}
	}
	return candidate
}

func (f *FileNode) String() string {
	return fmt.Sprintf("FileNode{size: %d, children: %v}", f.size, f.children)
}
