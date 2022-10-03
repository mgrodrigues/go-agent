package prm

import "os"

type Agent struct {
	Name string
	PID  int
}

func NewAgent(name string) *Agent {
	return &Agent{
		Name: name,
		PID:  os.Getegid(),
	}
}
