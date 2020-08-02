package stdout

import (
	"fmt"
	"os"

	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
)

//ConsoleOut would be used to manage console output
func ConsoleOut(processing bool, task string, msg string) {
	writer := goterminal.New(os.Stdout)

	if processing == true {
		ct.Foreground(ct.Yellow, false)
		fmt.Fprintln(writer, "RUNNING : ", task)
		writer.Print()
		ct.ResetColor()
		fmt.Fprintln(writer, "Log: ", msg)
		writer.Print()
	} else {
		// processing done here, after which color should change or text should be over-written.
		writer.Clear()

		ct.Foreground(ct.Green, false)
		fmt.Fprintln(writer, "COMPLETED ", task)
		writer.Print()
		ct.ResetColor()

		fmt.Fprintln(writer, "Log: ", msg)
		writer.Print()

	}

	writer.Reset()

}

//ConsoleError would be used to display error messages
func ConsoleError(processing bool, task string, msg string) {
	writer := goterminal.New(os.Stdout)

	if processing == true {
		ct.Foreground(ct.Red, false)
		fmt.Fprintln(writer, "FAILING : ", task)
		writer.Print()
		ct.ResetColor()
		fmt.Fprintln(writer, "Log: ", msg)
		writer.Print()
	} else {
		// processing done here, after which color should change or text should be over-written.
		writer.Clear()

		ct.Foreground(ct.Red, false)
		fmt.Fprintln(writer, "FAILED : ", task)
		writer.Print()
		ct.ResetColor()

		fmt.Fprintln(writer, "Log: ", msg)
		writer.Print()

	}

	writer.Reset()

}
