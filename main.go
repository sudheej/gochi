package main

import (
	"fmt"

	"github.com/sudheej/gochi/intro"
	"github.com/sudheej/gochi/plugins"
	"github.com/sudheej/gochi/queue"
)

func main() {
	intro.OpeningArt()
	fmt.Println("Testing yaml feature")
	plugins.GitClone()
	queue.Process()
	//fmt.Println("Memory ", debugger.GetMemory())
	//start := time.Now()
	//engine.MavenRunner()
	//elapsed := time.Since(start)
	//fmt.Printf("Total time taken %d sec", elapsed.Milliseconds()/1000)
}
