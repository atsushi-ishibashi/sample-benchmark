package main

type MemoryFunction struct {
	text string
}

func valueMemoryFunction() MemoryFunction {
	return MemoryFunction{"hoge"}
}

func pointerMemoryFunction() *MemoryFunction {
	return &MemoryFunction{"hoge"}
}

func valueFunction(mf MemoryFunction) string {
	return mf.text
}

func pointerFunction(mf *MemoryFunction) string {
	return mf.text
}
