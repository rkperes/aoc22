package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	const version = 1
	const startOfPacketSize = 4

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		slog.Error("reading standard input:", err)
		return
	}

	fmt.Println(findDataIndex(line, startOfPacketSize))

	slog.Info("done")
}

func findDataIndex(input string, size int) int {
	tracker := newPacketTracker(size)

	for idx, ch := range input {
		slog.Info("process rune", "rune", string(ch), "idx", idx, "tracked", tracker)

		tracker.put(ch)

		if idx >= size {
			if tracker.pop(rune(input[idx-size])) {
				slog.Info("found data index", "rune", string(ch), "idx", idx, "tracked", tracker, "startOfPacket", input[idx-size:idx])
				return idx + 1
			}
		}
	}

	return 0
}

type packetTracker struct {
	size int
	seen map[rune]int
}

func newPacketTracker(size int) *packetTracker {
	return &packetTracker{size: size, seen: make(map[rune]int, size)}
}

func (p *packetTracker) put(r rune) {
	p.seen[r]++
}

// pop removes a rune from the tracker and returns whether it found a complete start-of-packet.
func (p *packetTracker) pop(r rune) bool {
	if p.seen[r] > 1 {
		p.seen[r]--
	} else {
		delete(p.seen, r)
	}
	return len(p.seen) == p.size
}

func (p *packetTracker) String() string {
	seen := ""
	for r, count := range p.seen {
		for range count {
			seen += string(r)
		}
	}
	return fmt.Sprintf("%s (%d)", seen, len(p.seen))
}
