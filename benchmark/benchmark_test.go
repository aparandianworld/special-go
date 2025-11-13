package main

import (
	"strings"
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "A" + "B" + "C"
	}
}

func BenchmarkJoin(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString("A")
		sb.WriteString("B")
		sb.WriteString("C")
		_ = sb.String()
	}
}
