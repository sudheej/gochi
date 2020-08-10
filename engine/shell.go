package engine

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uilive"
	"github.com/sudheej/gochi/stdout"
)

//MavenRunner is the function is to run Maven commands
func MavenRunner(goals string) {
	defer timeTrack(time.Now(), "Maven "+goals)
	params := []string{"-f", filepath.FromSlash(".temp/pom.xml"), goals}
	executor("Maven "+goals, "mvn", params...)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %d sec ...\n", name, elapsed.Milliseconds()/1000)
}

func executor(invocation string, cmdx string, params ...string) {
	writer := uilive.New()
	app := cmdx

	cmd := exec.Command(app, params...)

	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	s := spinner.New(spinner.CharSets[9], 200*time.Millisecond)
	//s.Start()

	//fmt.Println(invocation)

	writer.Start()
	go func() {
		for scanner.Scan() {
			if len(scanner.Text()) > 150 {
				//fmt.Fprintf(writer, "\t > %s\n", scanner.Text()[:100])
				//stdout.ConsoleOut(true, invocation, scanner.Text()[:100])
				//stdout.ConsoleUI(scanner.Text()[:100])
				stdout.ConsoleTreat(scanner.Text()[:100])
				writer.Flush()
			}

		}
	}()

	err = cmd.Start()
	if err != nil {

		//fmt.Fprintln(writer, "Error starting Cmd", err.Error())
		stdout.ConsoleOut(true, invocation, err.Error())
		s.Stop()
		writer.Stop()
		return
	}

	err = cmd.Wait()
	if err != nil {
		//fmt.Fprintln(writer, "Error waiting for Cmd", err.Error())
		stdout.ConsoleOut(false, invocation, err.Error())
		s.Stop()
		writer.Stop()
		return
	}
	s.Stop()
	//writer.Stop()
	stdout.ConsoleOut(false, invocation, " Done !!")

}
