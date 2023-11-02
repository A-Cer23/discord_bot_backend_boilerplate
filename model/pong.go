package model

type Pong struct {
	Text string
}

func (pong *Pong) ResponsePong() (*Pong, error) {
	return pong, nil
}
