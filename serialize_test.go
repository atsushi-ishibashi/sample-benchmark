package main

import (
	"fmt"
	"testing"
)

var (
	testElements = make([]Element, 0)
)

func init() {
	for i := 0; i < 100; i++ {
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
	decode := MsgpackDeserialize(payload)
	if len(decode) != 100 {
		t.Errorf("error")
	}
}

func TestJSON(t *testing.T) {
	payload := JSONSerialize(testElements)
	decode := JSONDeserialize(payload)
	if len(decode) != 100 {
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
