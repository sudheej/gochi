package engine

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uilive"
)

//MavenRunner is the function is to run Maven commands
func MavenRunner() {
	defer timeTrack(time.Now(), "Maven Build")
	params := []string{"-f", "C:\\Users\\sudhe\\Documents\\playground\\Spring-Boot-Sample-Project-master\\Spring-Boot-Sample-Project-master\\pom.xml", "dependency-check:check", "install", "-Dskiptests"}
	executor("Maven Build", "mvn", params...)
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
	s := spinner.New(spinner.CharSets[9], 300*time.Millisecond)
	s.Start()

	fmt.Println(invocation)

	writer.Start()
	go func() {
		for scanner.Scan() {
			if len(scanner.Text()) > 150 {
				fmt.Fprintf(writer, "\033[36m"+"\t > %s\n", scanner.Text()[:150])
				writer.Flush()
			}

		}
	}()

	err = cmd.Start()
	if err != nil {

		fmt.Fprintln(writer, "Error starting Cmd", err.Error())
		s.Stop()
		writer.Stop()
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(writer, "Error waiting for Cmd", err.Error())
		s.Stop()
		writer.Stop()
		return
	}
	s.Stop()
	writer.Stop()
	fmt.Println(invocation + " completed..")

}
