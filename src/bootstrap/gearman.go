package bootstrap

import (

	"github.com/mikespook/gearman-go/worker"	
)

func Worker() {

	gearmanWorker := worker.New(worker.Unlimited)

	gearmanWorker.Work()	
}

func Client() {

	fmt.Println("HELLO WORLD")	
	//Bootstrap()

}


