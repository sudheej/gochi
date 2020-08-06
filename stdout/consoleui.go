package stdout

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

//ConsoleUI would be the terminals default ui interface
func ConsoleUI(logmessages string) {

	urlEntry := tui.NewEntry()
	urlEntry.SetText(logmessages)

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
