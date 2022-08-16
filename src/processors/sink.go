// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package processors

// Sink 用于输出数据的
type Sink struct {
	BaseProcessor
}

func NewSink(name string) IProcessor {
	return &Sink{
		BaseProcessor: NewBaseProcessor(name),
	}
}

func (p *Sink) In() *InPort {
	return p.BaseProcessor.In()
}
