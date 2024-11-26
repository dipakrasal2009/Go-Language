package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
)

const maxMenuItems = 15 // Maximum number of items to show in history
var historyItems []*systray.MenuItem

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("📋")
	systray.SetTooltip("Clipboard History")

	// Initialize history items slice
	historyItems = make([]*systray.MenuItem, 0, maxMenuItems)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Exit application")

	// Handle quit
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	// Start monitoring clipboard
	go monitorClipboard()
}

func monitorClipboard() {
	lastContent := ""
	for {
		content, _ := clipboard.ReadAll()
		if content != lastContent && content != "" {
			// Truncate content if too long
			displayContent := content
			if len(displayContent) > 50 {
				displayContent = displayContent[:47] + "..."
			}

			// Add new item to menu
			addToMenu(content, displayContent)
			lastContent = content
		}
		time.Sleep(time.Second)
	}
}

func addToMenu(content, displayContent string) {
	// Remove oldest item if we're at capacity
	if len(historyItems) >= maxMenuItems {
		oldest := historyItems[len(historyItems)-1]
		oldest.Hide()
		historyItems = historyItems[:len(historyItems)-1]
	}

	// Add new item at the top
	timestamp := time.Now().Format("15:04:05")
	menuItem := systray.AddMenuItem(fmt.Sprintf("%s: %s", timestamp, displayContent), content)

	// Handle click on menu item
	go func() {
		for range menuItem.ClickedCh {
			clipboard.WriteAll(content)
		}
	}()

	// Insert at beginning of slice
	historyItems = append([]*systray.MenuItem{menuItem}, historyItems...)
}

func onExit() {
	// Cleanup code here
}
