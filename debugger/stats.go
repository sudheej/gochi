package debugger

import (
	"bitbucket.org/bertimus9/systemstat"
)

//GetMemory would be used to retreive memory information for debugging
func GetMemory() uint64 {
	mem := systemstat.GetMemSample()
	return mem.MemTotal
}
