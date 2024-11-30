package main

import (
	"fmt"  //provide the formated input output function
	"time" //allowe working with time and duration

	"github.com/atotto/clipboard"   //library to interact with the clipboard
	"github.com/getlantern/systray" //create the system application to the notification area
)

const maxMenuItems = 15              //maximum number of item to store in the history
var historyItems []*systray.MenuItem // slice to hold the menu items representing clipboard history.

func main() { //this function is use to start the system tray application
	systray.Run(onReady, onExit)
	//onReady()  : call when the tray icon is ready
	// onExit()   : call when the application is exit
}

func onReady() {
	systray.SetTitle("📋")                   //set the tray icon title as a clipboard
	systray.SetTooltip("Clipboard History") //set the tooltip to describe the app

	// Initialize history items slice
	historyItems = make([]*systray.MenuItem, 0, maxMenuItems) //create the empty sclice to store the clipboard history menu iteam withe the  initial capacityof maxMenuIteam

	systray.AddSeparator()                                   //this function is use to seprate the two things that is history and quit option
	mQuit := systray.AddMenuItem("Quit", "Exit application") //it is use to exit the app

	// Handle quit
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}() //this function is use to quit the application

	// Start to monitering the clipboard
	go monitorClipboard()
}

func monitorClipboard() {
	lastContent := "" // to store the clipboard value
	for {
		content, _ := clipboard.ReadAll()            // to read the current clipboard content
		if content != lastContent && content != "" { // this if condition checks the new content comming in clipboard history
			// Truncate content if too long
			displayContent := content     //content goes to the new veriable that is displayContent
			if len(displayContent) > 50 { // this if function is truncate the clipboard content to 50 characters to display, appending ... if it  exceeds this limit
				displayContent = displayContent[:47] + "..."
			}

			// Add new item to menu
			addToMenu(content, displayContent) //to add the clipboard content into the menu
			lastContent = content              //last content with the new clipboard content
		}
		time.Sleep(time.Second) //Pauses for one second to reduce the frequency of clipboard checks.
	}
}

func addToMenu(content, displayContent string) {
	// Remove oldest item if we're at capacity
	// heare it use the stack data structure
	if len(historyItems) >= maxMenuItems {
		oldest := historyItems[len(historyItems)-1]
		oldest.Hide()
		historyItems = historyItems[:len(historyItems)-1]
	}

	// Add new item at the top
	timestamp := time.Now().Format("15:04:00") //thsi is the format of time
	//display the time and content of the data
	menuItem := systray.AddMenuItem(fmt.Sprintf("%s: %s", timestamp, displayContent), content)

	// Handle click on menu item
	go func() {
		for range menuItem.ClickedCh {
			clipboard.WriteAll(content) //when we click the copid content that is again write in the clipboard
		}
	}()

	// Insert at beginning of slice
	historyItems = append([]*systray.MenuItem{menuItem}, historyItems...)
}

func onExit() {
	// Cleanup code here
}
