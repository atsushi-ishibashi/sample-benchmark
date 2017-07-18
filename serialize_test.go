package main

import (
	"fmt"
	"log"
	"testing"
	"unsafe"
)

var (
	testElements = make([]Element, 0)
)

func init() {
	for i := 0; i < 1000; i++ {
		e := Element{
			BICD:  i,
			Code:  fmt.Sprintf("code%d", i),
			Value: int64(i) * 10,
		}
		testElements = append(testElements, e)
	}
}

func TestMsgpack(t *testing.T) {
	payload := MsgpackSerialize(testElements)
	log.Printf("msgpack size: %d\n", unsafe.Sizeof(payload))
	decode := MsgpackDeserialize(payload)
	if len(decode) != 1000 {
		t.Errorf("error")
	}
}

func TestJSON(t *testing.T) {
	payload := JSONSerialize(testElements)
	log.Printf("json size: %d\n", unsafe.Sizeof(payload))
	decode := JSONDeserialize(payload)
	if len(decode) != 1000 {
		t.Errorf("error")
	}
}

func BenchmarkMsgpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := MsgpackSerialize(testElements)
		_ = MsgpackDeserialize(payload)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := JSONSerialize(testElements)
		_ = JSONDeserialize(payload)
	}
}

func BenchmarkPMsgpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := MsgpackPSerialize(testElements)
		_ = MsgpackPDeserialize(payload)
	}
}

func BenchmarkPJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		payload := JSONPSerialize(testElements)
		_ = JSONPDeserialize(payload)
	}
}
