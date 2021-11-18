package main

import "time"

type Pool struct {
	Heap chan Item
	Capacity int
}

func (p *Pool) Init() {
	p.Heap = make(chan Item, p.Capacity)
	p.Heap <- Item{
		Duration: time.Second,
		Name:     "Magic Item #1",
	}
	p.Heap <- Item{
		Duration: time.Second,
		Name:     "Magic Item #2",
	}
}

type Item struct {
	Duration time.Duration
	Name string
}

func (i Item) WorkIt() {
	time.Sleep(i.Duration)
}

func main() {
	p := Pool{
		Capacity: 5,
	}
	p.Init() // inits the pool. makes the buffered channel and fills it.
	item := <- p.Heap
	item.WorkIt()
	p.Heap <- item
}