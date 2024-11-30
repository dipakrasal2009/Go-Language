## Clipboard History Manager

This is a Clipboard History Manager written in Go, which allows users to monitor their clipboard activity and access a history of copied content via a system tray icon. The application keeps track of up to 15 recent clipboard entries, enabling easy recall and reuse.

<b>Features</b>

    Clipboard Monitoring: Automatically tracks clipboard changes.
    System Tray Integration: Displays a tray icon with clipboard history.
    History Management: Stores up to 15 clipboard entries, removing the oldest ones when the limit is reached.
    Quick Copy: Clicking on a menu item copies its content back to the clipboard.
    Time-stamped Entries: Each entry is labeled with the time it was copied.
    Lightweight: Operates seamlessly in the background.

<b>ow It Works</b>

    The application starts by creating a system tray icon.
    It monitors the clipboard continuously for new content.
    When new content is detected:
        The content is truncated to 50 characters for display in the tray menu.
        A new menu item is added to the top of the clipboard history.
    Clicking on a history item copies the content back to the clipboard.
    Exiting the application can be done via the Quit menu item.

<b>Installation Prerequisites</b>

    Go installed on your system.
    A compatible operating system (Windows, macOS, Linux).

<b>Steps</b>

    Clone this repository:

git clone https://github.com/your-username/clipboard-history-manager.git
cd clipboard-history-manager

<b>Install dependencies:</b>

go mod tidy

<b>Build the application:</b>

go build -o clipboard-manager

<b>Run the application:</b>

    ./clipboard-manager

<b>Usage</b>

    Start the application.
    The tray icon will appear in your system's notification area.
    Copy text to the clipboard as usual. The text will be added to the clipboard history.
    Click on the tray icon to view the history.
    Select an item from the history to copy it back to the clipboard.
    Use the Quit option to exit the application.

<b>Code Overview Key Libraries</b>

    systray: Manages the system tray icon and menu.
    clipboard: Provides clipboard read and write functionality.

<b>File Explanation</b>

    main(): Initializes the application and starts the system tray.
    onReady(): Sets up the tray icon, tooltip, and menu structure. Starts clipboard monitoring.
    monitorClipboard(): Continuously checks the clipboard for changes and updates the menu.
    addToMenu(content, displayContent):
        Adds a new clipboard entry to the menu.
        Removes the oldest entry if the history exceeds the maximum limit.
    onExit(): Placeholder for cleanup code when the application exits.

<b>Future Enhancements</b>

    Add configuration options for:
        Maximum history size.
        Clipboard monitoring frequency.
    Support for rich clipboard content (e.g., images).
    Cross-platform testing and improvements.

<b>License</b>

This project is licensed under the MIT License. See the LICENSE file for details.
Contributing

Contributions are welcome! Feel free to submit issues or pull requests.
Acknowledgments

    Lantern's systray library: For the robust system tray integration.
    Atotto's clipboard library: For easy clipboard manipulation.
