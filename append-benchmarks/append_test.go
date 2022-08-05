package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/scristobal/go-experiments/random"
)

func AppendConcat(s, t string) string {
	return s + t
}

func AppendBuilder(s, t string) string {

	var builder strings.Builder

	builder.WriteString(s)
	builder.WriteString(t)

	return builder.String()

}

func AppendCopy(s, t string) string {
	r := make([]byte, len(s)+len(t))
	copy(r[:len(s)], s)
	copy(r[len(s):], t)
	return string(r)
}

func AppendJoin(s, t string) string {

	return strings.Join([]string{s, t}, "")

}

func AppendFprint(s, t string) string {

	buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, s, t)

	return buffer.String()
}

func AppendSprintf(s, t string) string {
	return fmt.Sprintf("%s%s", s, t)
}

const SIZE = 1000
const REPS = 10

var s = random.RandSeq(SIZE)

func BenchmarkConcat(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendConcat(s, r)
		}

	}
}
func BenchmarkBuilder(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendBuilder(s, r)
		}

	}
}

func BenchmarkCopy(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendCopy(s, r)
		}
	}
}

func BenchmarkJoin(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendJoin(s, r)
		}
	}
}

func BenchmarkFprint(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendFprint(s, r)
		}
	}
}

func BenchmarkSprintf(b *testing.B) {

	for i := 0; i < b.N; i++ {

		var r string

		for i := 0; i < REPS; i++ {
			r = AppendSprintf(s, r)
		}
	}
}
