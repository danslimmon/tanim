package tanim

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/sirupsen/logrus"
)

type stackEntry struct {
	LastOrigin Dim
	LastExtent Dim
	Fig        Figure
}

type Animation struct {
	done   chan struct{}
	screen tcell.Screen
	stack  []*stackEntry
	ticker chan int
}

func (a Animation) TickEvery(dur time.Duration) {
	tticker := time.Tick(dur)
	go func() {
		i := 0
		for {
			select {
			case <-tticker:
				a.ticker <- i
			case <-a.done:
				return
			}
			i++
		}
	}()
}

func (a Animation) mainLoop() {
	for {
		select {
		case t := <-a.ticker:
			a.onTick(t)
		case <-a.done:
			return
		}
		a.screen.Show()
	}
}

func (a Animation) handleEvents() {
	for {
		switch ev := a.screen.PollEvent().(type) {
		case *tcell.EventResize:
			a.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				a.screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func (a Animation) onTick(t int) {
	for _, se := range a.stack {
		se.Fig.OnTick(t)
	}
	for _, se := range a.stack {
		origin, extent := se.Fig.Origin(), se.Fig.Extent()

		// Erase any cells that used to be occupied by the Figure but are no longer
		lastRange := dimRange{se.LastOrigin, se.LastExtent}
		newRange := dimRange{origin, origin.Add(extent)}
		for _, cell := range lastRange.Sub(newRange) {
			logger.WithField("cell", cell).Info("erasing cell")
			a.screen.SetContent(cell.X, cell.Y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))
		}

		// uh, replace this with a dimRange and get rid of Dim.Each
		extent.Each(func(cell Dim) {
			draw, char, style := se.Fig.DrawCell(cell)
			if draw {
				a.screen.SetContent(origin.X+cell.X, origin.Y+cell.Y, char, nil, style)
			}
		})

		se.LastOrigin = origin
		se.LastExtent = origin.Add(extent)
	}
}

func (a Animation) Start() {
	a.screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))
	a.screen.Clear()
	go a.handleEvents()
	go a.mainLoop()
}

func (a Animation) Stop() {
	a.screen.Clear()
	close(a.done)
}

func (a Animation) Wait() {
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewAnimation(figs []Figure) (*Animation, error) {
	stack := make([]*stackEntry, len(figs))
	for i, fig := range figs {
		stack[i] = &stackEntry{Fig: fig}
	}

	// Init tcell screen
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	err = screen.Init()
	if err != nil {
		return nil, err
	}

	// Init logger
	logger = logrus.New()
	f, err := os.Create("/tmp/tanim.log")
	if err != nil {
		return nil, fmt.Errorf("error creating log file: %w", err)
	}
	logger.SetOutput(f)

	return &Animation{
		done:   make(chan struct{}),
		stack:  stack,
		screen: screen,
		ticker: make(chan int, 1),
	}, nil
}
