package day01

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

var input string

func TestMain(m *testing.M) {
	dataPath := filepath.Join("..", "testdata", "day01.txt")
	dataBytes, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Test data could not be read: %v", err)
	}
	input = string(dataBytes)
	m.Run()
}

func TestPart1(t *testing.T) {
	const want = -1
	got := Part1(input)

	if got != want {
		t.Errorf("Part1(input) = %v; want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	const want = -1
	got := Part2(input)

	if got != want {
		t.Errorf("Part2(input) = %v; want %v", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
