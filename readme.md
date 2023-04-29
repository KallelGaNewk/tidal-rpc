# TIDAL Rich Presence for Discord

This application provides TIDAL Rich Presence integration for Discord, allowing you to display your currently playing TIDAL song as your Discord status. It uses the TIDAL desktop client's window title to retrieve the currently playing song information and updates your Discord status accordingly.

<small>Consider this as a humble project, as I'm still learning the language.</small>

## Usage

1. Run the built executable file (`TIDAL-RPC.exe`).
2. The application will appear in your system tray.
3. Right-click the application icon in the system tray to access the menu.
4. By default, the Rich Presence integration is enabled.
5. The menu options allow you to control the application and the Rich Presence feature:

   - **Enable RPC**: Toggle the Rich Presence integration on or off.
   - **Artist**: Shows the currently playing artist.
   - **Song name**: Shows the name of the currently playing song.
   - **Quit**: Exit the application.

## Notes

- **Currently, if the TIDAL app is minimized to the system tray, there is no active window available, which prevents the functionality from working at the moment.**
- If no song is playing, the application sets your Discord status to "Idling."

## Build Instructions

To build the application, follow these steps:

1. Ensure you have Go installed on your system.
2. Open a terminal or command prompt.
3. Navigate to the project directory.
4. Run the following command to build the executable:

```
go build -o TIDAL-RPC.exe -ldflags -H=windowsgui
```

This command builds the application as a GUI and disables the console window.

## Dependencies

This application relies on the following dependencies:

- [github.com/getlantern/systray](https://github.com/getlantern/systray): A cross-platform Go library for placing an icon and menu in the notification area.
- [github.com/hugolgst/rich-go](https://github.com/hugolgst/rich-go): A Go library for Discord Rich Presence integration.

## Contributing

Contributions to this project are welcome. If you encounter any issues or have suggestions for improvements, please create an issue or a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
