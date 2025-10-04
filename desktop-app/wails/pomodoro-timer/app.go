package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx          context.Context
	Timer        *time.Ticker
	TimerDone    chan bool
	Events       []PomodoroEvent
	CurrentEvent *PomodoroEvent
}

const (
	DefaultPomodoroDuration = time.Second * 20 // licenca poetica, pq idealmente sao 30 mins
)

type PomodoroEvent struct {
	StartTime time.Time
	EndTime   time.Time
	Name      string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Events: make([]PomodoroEvent, 0), // persistir num db, fazer o frontend mostrar etc
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// StartTimer starts the pomodoro timer
func (a *App) StartTimer() {
	a.TimerDone = make(chan bool)

	a.Timer = time.NewTicker(1 * time.Second)

	// Cria um stop timer para parar o pomodoro timer em DefaultPomodoroDuration
	stopTimer := time.NewTimer(DefaultPomodoroDuration)
	startTime := time.Now()
	event := new(PomodoroEvent)
	event.StartTime = startTime
	event.Name = "pomodoro-event"
	a.CurrentEvent = event
	a.Events = append(a.Events, *event)

	go func() {
		<-stopTimer.C
		a.Timer.Stop()
		runtime.EventsEmit(a.ctx, "timer:tick", 0)
		runtime.EventsEmit(a.ctx, "timer:complete")
	}()

	go func() {
		for {
			select {
			case <-a.TimerDone:
				fmt.Println("Ticker stopped, resetting remaining to:", 0)
				// por mais que eu possa usar a variavel event, usamos aqui
				// a.CurrentEvent pra garantir que é o evento atual
				a.CurrentEvent.EndTime = time.Now()
				fmt.Println(a.Events)
				runtime.EventsEmit(a.ctx, "timer:tick", 0)
				runtime.EventsEmit(a.ctx, "timer:complete")
				return
			case t := <-a.Timer.C:
				elapsed := time.Since(startTime)
				remaining := DefaultPomodoroDuration - elapsed
				if remaining < 0 {
					remaining = 0
				}
				fmt.Println("Tick at", t, "Remaining:", remaining)
				runtime.EventsEmit(a.ctx, "timer:tick", int(remaining.Seconds()))
			}
		}
	}()
}

// StopTimer stops the pomodoro timer
func (a *App) StopTimer() {
	a.Timer.Stop()
	a.TimerDone <- true
	fmt.Println("Ticker stopped")
}

// GetEvents returns all pomodoro events
func (a *App) GetEvents() []PomodoroEvent {
	return a.Events
}
