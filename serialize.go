package main

import (
	"encoding/json"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type Element struct {
	BICD  int    `json:"bicd"`
	Code  string `json:"code"`
	Value int64  `json:"value"`
}

func MsgpackSerialize(elements []Element) []byte {
	payload, _ := msgpack.Marshal(elements)
	return payload
}

func MsgpackDeserialize(payload []byte) []Element {
	var elements []Element
	msgpack.Unmarshal(payload, &elements)
	return elements
}

func JSONSerialize(elements []Element) []byte {
	payload, _ := json.Marshal(elements)
	return payload
}

func JSONDeserialize(payload []byte) []Element {
	var elements []Element
	json.Unmarshal(payload, &elements)
	return elements
}

func MsgpackPSerialize(elements []Element) []byte {
	payload, _ := msgpack.Marshal(&elements)
	return payload
}

func MsgpackPDeserialize(payload []byte) []Element {
	var elements []Element
	msgpack.Unmarshal(payload, &elements)
	return elements
}

func JSONPSerialize(elements []Element) []byte {
	payload, _ := json.Marshal(&elements)
	return payload
}

func JSONPDeserialize(payload []byte) []Element {
	var elements []Element
	json.Unmarshal(payload, &elements)
	return elements
}
