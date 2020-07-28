package queue

//WorkerGroup is exported as struct
type WorkerGroup struct {
	primary   chan bool
	completed chan func()
}

//New to invoke woker group
func New(n int) WorkerGroup {

	done := make(chan bool)
	operations := WorkerGroup{
		primary:   make(chan bool),
		completed: make(chan func()),
	}

	for i := 0; i < n; i++ {
		go func() {
			for f := range operations.completed {
				f()
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <-done
		}
		operations.primary <- true

	}()
	return operations

}

func (wg WorkerGroup) Add(f func()) {
	wg.completed <- f
}

func (wg WorkerGroup) Wait() {
	close(wg.completed)
	_ = <-wg.primary

}
