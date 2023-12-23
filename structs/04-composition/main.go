package main

import "fmt"

type CPU struct {
	Brand string
	Core  int
	Speed float64
}
type Memory struct {
	Capacity       int
	FrequencyInMHz int
}
type Srorage struct {
	Capacity int
	Type     string
}
type Computer struct {
	cpu     CPU
	memory  Memory
	storage Srorage
}

func main() {
	c := CPU{Brand: "Intel", Core: 4, Speed: 2.5}
	m := Memory{Capacity: 16, FrequencyInMHz: 2666}
	s := Srorage{Capacity: 512, Type: "NVMe"}

	comp := Computer{
		cpu:     c,
		memory:  m,
		storage: s,
	}

	fmt.Printf("%+v\n", comp)
	fmt.Println("Core : ", comp.cpu.Core)
	fmt.Println("Memory Capacity : ", comp.memory.Capacity)
}
