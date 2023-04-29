package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
)

const (
	LoopTimeout = 16 * time.Second
	ClientID    = "1098099931006390325"
)

type Song struct {
	Name   string
	Artist string
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	var RPCEnabled = true

	err := client.Login(ClientID)
	if err != nil {
		panic(err)
	}

	iconBytes := getIconBytes()
	systray.SetIcon(iconBytes)
	titleMenu := systray.AddMenuItem("TIDAL Rich Presence for Discord", "TIDAL")
	rpcToggle := systray.AddMenuItemCheckbox("Enable RPC", "RPC", RPCEnabled)
	titleMenu.SetIcon(iconBytes)
	systray.AddSeparator()
	artistMenu := systray.AddMenuItem("Artist: None", "Artist")
	songMenu := systray.AddMenuItem("Song name: None", "Name")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "Quit the app, duh.")

	go func() {
		for {
			tidalWindowName := getProcessWindowName("TIDAL")
			song := formatSongName(tidalWindowName)

			if song != nil {
				songMenu.Show()
				artistMenu.SetTitle(fmt.Sprintf("Artist: %s", song.Artist))
				songMenu.SetTitle(fmt.Sprintf("Song name: %s", song.Name))

				client.SetActivity(client.Activity{
					Details:    song.Name,
					State:      fmt.Sprintf("by %s", song.Artist),
					LargeImage: "tidal",
					LargeText:  "Listen to music the way it’s meant to sound.",
				})
			} else {
				artistMenu.SetTitle("Not playing nothing")
				songMenu.Hide()

				client.SetActivity(client.Activity{
					Details:    "Idling",
					LargeImage: "tidal",
					LargeText:  "Listen to music the way it’s meant to sound.",
				})
			}

			time.Sleep(LoopTimeout)
		}
	}()

	go func() {
		for {
			select {
			case <-rpcToggle.ClickedCh:
				if rpcToggle.Checked() {
					rpcToggle.Uncheck()
					client.Logout()
				} else {
					rpcToggle.Check()
					err := client.Login(ClientID)
					if err != nil {
						panic(err)
					}
				}

				RPCEnabled = rpcToggle.Checked()
			case <-quitItem.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	fmt.Println("Quitting")
}

func getIconBytes() []byte {
	file, err := os.Open("icon.ico")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	iconBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return iconBytes
}

func formatSongName(inputString string) *Song {
	parts := strings.Split(inputString, " - ")
	if len(parts) != 2 {
		return nil
	}

	songName := strings.TrimSpace(parts[0])
	songArtist := strings.TrimSpace(parts[1])

	return &Song{
		Name:   songName,
		Artist: songArtist,
	}
}

func getProcessWindowName(processName string) string {
	pwsh_instance := exec.Command("powershell", fmt.Sprintf("(Get-Process %s).MainWindowTitle", processName))

	// Hide PowerShell window from showing
	// https://stackoverflow.com/a/48365926
	pwsh_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := pwsh_instance.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return strings.TrimSpace(string(output))
}
