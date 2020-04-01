package main

import (
	"fmt"
	"time"
)

type Agent struct {
	createAt time.Time
	endAt    time.Time
	name     string
	child    *Agent
}

func NewAgent() *Agent {
	return &Agent{
		name:     "super",
		createAt: time.Now(),
		child:    nil,
	}
}

func (a *Agent) Close() {
	if !a.endAt.IsZero() {
		return
	}
	a.endAt = time.Now()
	a.child.Close()
}

func (a *Agent) Span(name string) *Agent {
	a.child = &Agent{
		createAt: time.Now(),
		name:     name,
	}
	return a.child
}

func (a *Agent) Collect() {
	if a == nil {
		return
	}

	if a.endAt.IsZero() {
		a.endAt = time.Now()
	}
	fmt.Printf("%s cost time %d\n", a.name, a.endAt.Sub(a.createAt))

	a.child.Collect()
}

func main() {
	a := NewAgent()

	span := a.Span("span")

	subSpan := span.Span("subSpan")

	time.Sleep(1 * time.Second)

	subSpan.Close()

	span.Close()

	a.Collect()

}
