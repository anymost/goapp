package main

type Empty interface {

}
type Semaphore chan Empty

func (s Semaphore) P(n int)  {
	for i := 0; i < n; i++ {
		s <- n
	}
}

func (s Semaphore) V(n int)  {
	for i:= 0; i < n; i++ {
		<- s
	}
}

func (s Semaphore) Lock ()  {
	s.P(1)
}

func (s Semaphore) Unlock()  {
	s.V(1)
}

func (s Semaphore) Wait(n int)  {
	s.P(n)
}

func (s Semaphore) Signal()  {
	s.V(1)
}

func main() {
	
}
