package main

import (
	"time"

	"github.com/getlantern/systray"
)

var duration time.Duration

// StartSysTray boots up the system tray icon
func StartSysTray() time.Duration {
	systray.Run(onReady, onExit)

	return duration
}

func onReady() {
	menuQuit := systray.AddMenuItem("Stop tracking", "Stop tracking time")
	start := time.Now()

	go func() {
		for {
			systray.SetTitle(getClockTime(start))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-menuQuit.ClickedCh:
				duration = time.Since(start)
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// ...
}

func getClockTime(start time.Time) string {
	diff := time.Since(start)

	hour, min, sec := int(diff.Hours()), int(diff.Minutes()), int(diff.Seconds())
	return ItoaTwoDigits(hour) + ":" + ItoaTwoDigits(min) + ":" + ItoaTwoDigits(sec)
}
