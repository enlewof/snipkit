package ui

import (
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
)

func runScreenTest(t *testing.T, procedure func(s tcell.Screen), test func(s tcell.SimulationScreen)) {
	t.Helper()
	screen := mkTestScreen(t)

	donec := make(chan struct{})
	go func() {
		defer close(donec)
		time.Sleep(time.Millisecond * 50)
		test(screen)
	}()

	procedure(screen)
	<-donec
}

func mkTestScreen(t *testing.T) tcell.SimulationScreen {
	t.Helper()
	s := tcell.NewSimulationScreen("")

	if s == nil {
		t.Fatalf("Failed to get simulation screen")
	}
	if e := s.Init(); e != nil {
		t.Fatalf("Failed to initialize screen: %v", e)
	}
	return s
}
