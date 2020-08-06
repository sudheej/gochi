package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

func main() {

	urlEntry := tui.NewEntry()
	urlEntry.SetText("https://httpbin.org/get")

	urlBox := tui.NewHBox(urlEntry)
	urlBox.SetTitle("URL")
	urlBox.SetBorder(true)

	root := tui.NewVBox(urlBox)

	tui.DefaultFocusChain.Set(urlEntry)

	theme := tui.NewTheme()
	theme.SetStyle("box.focused.border", tui.Style{Fg: tui.ColorYellow, Bg: tui.ColorDefault})

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetTheme(theme)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
