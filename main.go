package main

import (
	"github.com/sudheej/gochi/intro"
	"github.com/sudheej/gochi/queue"
)

func main() {
	intro.OpeningArt()
	queue.Process()
	//fmt.Println("Memory ", debugger.GetMemory())
	//start := time.Now()
	//engine.MavenRunner()
	//elapsed := time.Since(start)
	//fmt.Printf("Total time taken %d sec", elapsed.Milliseconds()/1000)
}
