package pkg

import "fmt"

type Server struct {
	Type   string
	Cpu    int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Cpu:    196,
		Memory: 256,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	fmt.Printf("%s, Cpu:[%v], Memory:[%v]\n", pc.Type, pc.Cpu, pc.Memory)
}
