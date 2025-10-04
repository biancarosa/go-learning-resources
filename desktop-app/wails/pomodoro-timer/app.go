package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// StartTimer starts the pomodoro timer
func (a *App) StartTimer() {
	// TODO: Implement timer start logic
}

// StopTimer stops the pomodoro timer
func (a *App) StopTimer() {
	// TODO: Implement timer stop logic
}
