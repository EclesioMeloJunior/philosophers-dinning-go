package philosophers

import (
	"log"
	"os"
	"sync"
	"time"
)

const thinking = time.Second * 3
const eating = time.Second * 3

var (
	ph  = []string{"philosopher 00", "philosopher 01", "philosopher 02", "philosopher 03"}
	fmt = log.New(os.Stdout, "", 0)
	w   sync.WaitGroup
)

func resolution(ph string, right, left *sync.Mutex) {
	fmt.Println(ph, "seated")

	left.Lock()
	right.Lock()

	fmt.Println(ph, "started eating")

	time.Sleep(eating)

	fmt.Println(ph, "started thinking")
	right.Unlock()
	left.Unlock()

	time.Sleep(thinking)

	w.Done()
	fmt.Println(ph, "left the table")
}

func Problem() {
	w.Add(len(ph))

	fmt.Println("table empty")

	firstFork := &sync.Mutex{}
	lefthand := firstFork

	for _, p := range ph[1:] {
		rigthhand := &sync.Mutex{}
		go resolution(p, rigthhand, lefthand)
		lefthand = rigthhand
	}

	go resolution(ph[0], firstFork, lefthand)
	w.Wait()
	fmt.Println("table empty")
}
