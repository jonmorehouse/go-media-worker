package bootstrap

import (

	"github.com/mikespook/gearman-go/worker"	
	"fmt"
)

func Worker() {

	gearmanWorker := worker.New(worker.Unlimited)

	gearmanWorker.Work()	
}

func Client() {

	fmt.Println("HELLO WORLD")	
	//Bootstrap()

}


