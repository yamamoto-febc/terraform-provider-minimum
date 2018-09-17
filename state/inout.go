package state

import "time"

type InOutValue struct {
	Required         string
	Optional         string
	OptionalDefault  string
	OptionalComputed string
	InputOnly        string
}

func (b *InOutValue) Computed() string {
	return time.Now().Format(time.RFC3339)
}

func (b *InOutValue) SetOptionalComputed(v string) {
	if v == "" {
		b.OptionalComputed = time.Now().Format(time.RFC3339)
	}
}

func NewInOutValue(required string, optional string, optionalComputed string, inputOnly string) *InOutValue {
	value := &InOutValue{
		Required:  required,
		Optional:  optional,
		InputOnly: inputOnly,
	}
	value.SetOptionalComputed(optionalComputed)
	return value
}
