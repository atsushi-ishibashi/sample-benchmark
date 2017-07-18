package main

import "testing"

func BenchmarkValueFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := valueMemoryFunction()
		_ = valueFunction(mf)
	}
}

func Benchmark_pointer_ValueFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := pointerMemoryFunction()
		_ = valueFunction(*mf)
	}
}

func Benchmark_value_PointerFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := valueMemoryFunction()
		_ = pointerFunction(&mf)
	}
}

func BenchmarkPointerFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := pointerMemoryFunction()
		_ = pointerFunction(mf)
	}
}
