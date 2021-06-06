package main

import (
    "testing"
)

func BenchmarkStack(b *testing.B) {
    for N := 0; N < b.N; N++ {
        s := make([]int, 5)
        for i := range s {
            s[i] = 42
        }
    }
}

func BenchmarkHeap(b *testing.B) {
    var s []int
    for N := 0; N < b.N; N++ {
        s = make([]int, 5)
        for i := range s {
            s[i] = 42
        }
    }
}

func BenchmarkHygull(b *testing.B) {
    for N := 0; N < b.N; N++ {
        var s []int
        for i := 0; i < 5; i++ {
            s = append(s, 42)
        }
    }
}
