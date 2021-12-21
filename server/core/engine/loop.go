package engine

import (
	"time"
)

type Loop struct {
	TickRate uint16
	ticker   *time.Ticker
	quit     chan struct{}
}

func (gl *Loop) Start(callback func()) {
	gl.ticker = time.NewTicker(time.Second / time.Duration(gl.TickRate))
	gl.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-gl.ticker.C:
				callback()
			case <-gl.quit:
				gl.ticker.Stop()
				return
			}
		}
	}()
}

func (gl *Loop) Stop() {
	close(gl.quit)
}

func NewLoop() *Loop {
	return &Loop{TickRate: 64}
}
