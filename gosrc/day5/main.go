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
	supply := NewSupply()
	cmds := NewCommands()

	part := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")

		if len(line) == 0 {
			part++
			continue
		}

		switch part {
		case 0:
			supply.ParseAndAddLine(line)
		case 1:
			cmds.ParseAndAddCommand(line)
		default:
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	slog.Info("parsed supply", "supply", supply)
	slog.Info("parsed commands", "cmds", cmds)

	for _, cmd := range cmds.GetCommands() {
		slog.Info("run command", "cmd", cmd)
		supply.ApplyCommand(cmd)
	}

	fmt.Println(supply.Result())

	slog.Info("done")
}

type Supply struct {
	Stacks [][]string
}

func NewSupply() *Supply {
	return &Supply{
		Stacks: make([][]string, 0),
	}
}

func (s Supply) Result() string {
	b := &strings.Builder{}

	for _, stack := range s.Stacks {
		fmt.Fprint(b, stack[0])
	}

	return b.String()
}

func (s *Supply) ParseAndAddLine(line string) {
	idx := 0
	cur := line
	for {
		item, ok := parseItem(cur)
		if !ok {
			return
		}

		if item != EmptyItem {
			for len(s.Stacks) < idx+1 {
				s.Stacks = append(s.Stacks, []string{})
			}
			s.Stacks[idx] = append(s.Stacks[idx], item)
		}

		idx++
		cur = cur[3:]

		// if it has more than one item, skip space separator
		if len(cur) > 3 {
			cur = cur[1:]
		}
	}
}

func (s *Supply) ApplyCommand(cmd Command) {
	for idx := 0; idx < cmd.Repeats; idx++ {
		item := s.popStack(cmd.Source)
		s.pushStack(cmd.Target, item)
	}
}

func (s *Supply) popStack(stack int) string {
	slog.Info("pop", "stack", stack, "stacks", s.Stacks)

	item := s.Stacks[stack][0]
	s.Stacks[stack] = s.Stacks[stack][1:]
	return item
}

func (s *Supply) pushStack(stack int, item string) {
	// if stack >= len(s.Stacks) {
	// 	s.Stacks = append(s.Stacks, []string{})
	// }
	s.Stacks[stack] = append([]string{item}, s.Stacks[stack]...)
}

const EmptyItem = ""

func parseItem(in string) (string, bool) {
	if len(in) < 3 {
		return "", false
	}

	if in[0] != '[' {
		return EmptyItem, true
	}

	slog.Info("got item", "item", string(in[1]), "in", in)
	return string(in[1]), true
}

func (s Supply) String() string {
	b := &strings.Builder{}

	for idx, stack := range s.Stacks {
		fmt.Fprintf(b, "[%d: ", idx)
		for jdx, item := range stack {
			if jdx > 0 {
				fmt.Fprint(b, ",")
			}
			fmt.Fprintf(b, "%s", item)
		}
		fmt.Fprintf(b, "]")
	}

	return b.String()
}

type Command struct {
	Repeats int
	Source  int
	Target  int
}

type Commands struct {
	cmds []Command
}

func NewCommands() *Commands {
	c := &Commands{
		cmds: make([]Command, 0),
	}
	return c
}

func (c *Commands) ParseAndAddCommand(line string) bool {
	parts := strings.Split(line, " ")

	if len(parts) < 6 {
		return false
	}

	repeat, err := strconv.Atoi(parts[1])
	if err != nil {
		slog.Error("bad command repeat", err)
		return false
	}

	src, err := strconv.Atoi(parts[3])
	if err != nil {
		slog.Error("bad command source", err)
		return false
	}

	tgt, err := strconv.Atoi(parts[5])
	if err != nil {
		slog.Error("bad command target", err)
		return false
	}

	c.cmds = append(c.cmds, Command{
		Repeats: repeat,
		Source:  src - 1,
		Target:  tgt - 1,
	})

	return true
}

func (c *Commands) GetCommands() []Command {
	return c.cmds
}

func (c *Commands) String() string {
	b := &strings.Builder{}

	for idx, cmd := range c.cmds {
		if idx > 0 {
			fmt.Fprint(b, " ")
		}

		fmt.Fprintf(b, "%d*(%d->%d)", cmd.Repeats, cmd.Source, cmd.Target)
	}

	return b.String()
}
