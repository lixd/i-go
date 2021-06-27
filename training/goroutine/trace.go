package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 我们使用服务端埋点来跟踪记录一些事件。

type Tracker struct {
}

func (t *Tracker) Event(data string) {
	time.Sleep(time.Millisecond)
	log.Println(data)
}

// 在 Handler 方法中调用 Event 进行记录

type App struct {
	track Tracker
}

func (a *App) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// BUG 无法管控该 goroutine 的生命周期
	// 服务关闭时无法保证 goroutine 能执行完
	go a.track.Event("this event")
}

// 优化后

// Tracker2 knows how to track events for the application .
type Tracker2 struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker2() *Tracker2 {
	return &Tracker2{
		ch: make(chan string, 10),
	}
}

func (t *Tracker2) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (t *Tracker2) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}
func (t *Tracker2) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}

func main() {
	tr := NewTracker2()
	go tr.Run()
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}
