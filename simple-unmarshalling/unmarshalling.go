package reddit

import (
	"encoding/json"
	"fmt"
)

type Type int

type Foo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Bar struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Message struct {
	Type Type `json:"type"`
	Data any  `json:"data"`
}

const FooType Type = 1
const BarType Type = 2

func (m *Message) UnmarshalJSON(b []byte) error {

	var msg map[string]any

	_ = json.Unmarshal(b, &msg)

	t, ok := msg["type"].(Type)

	if !ok {
		return fmt.Errorf("message needs a type field")
	}

	m.Type = t

	switch t {
	case FooType:
		m.Data = msg["data"].(Foo)
		return nil
	case BarType:
		m.Data = msg["data"].(Bar)
		return nil
	}

	return fmt.Errorf("unknown type %d", t)
}
