package main

import (
	"encoding/json"
	"fmt"
)

const (
	BaseType    = "base"
	MessageType = "message"
	ErrorType   = "error"
)

type Base struct {
	Mtype string `json:"type"`
}

func (b *Base) mtype() string  { return BaseType }
func (b *Base) render() string { return "" }

type Message struct {
	Mtype   string `json:"type"`
	Message string `json:"message"`
}

func (m *Message) mtype() string  { return MessageType }
func (m *Message) render() string { return m.Message }

type ErrorMessage struct {
	Mtype  string `json:"type"`
	Merror string `json:"error"`
}

func (e *ErrorMessage) mtype() string  { return ErrorType }
func (e *ErrorMessage) render() string { return e.Merror }

type jsonMessage interface {
	mtype() string
	render() string
}

func unpack(m []byte) jsonMessage {
	var b Base

	if err := json.Unmarshal(m, &b); err != nil {
		panic(err)
	}

	if b.Mtype == "base" {
		return &b
	} else if b.Mtype == "message" {
		out := &Message{}
		json.Unmarshal(m, &out)
		return out
	} else if b.Mtype == "error" {
		out := &ErrorMessage{}
		json.Unmarshal(m, &out)
		return out
	} else {
	}

	return nil
}

func show(j jsonMessage) {
	if j != nil {
		fmt.Println("show:", j.mtype(), j.render())
	} else {
		fmt.Println("show: nil")
	}
}

func main() {
	inboundB := []byte(`{"type":"base"}`)
	inboundM := []byte(`{"type":"message", "message":"hello world"}`)
	inboundE := []byte(`{"type":"error", "error":"goodbye world"}`)
	inboundO := []byte(`{"type":"other", "something": {"a": 1}, "else":[1,2,3]}`)

	var m jsonMessage

	m = unpack(inboundB)
	show(m) // show: base

	m = unpack(inboundM)
	show(m) // show: message hello world

	m = unpack(inboundE)
	show(m) // show: error goodbye world

	m = unpack(inboundO)
	show(m) // show: nil
}
