package bootstrap

import (

	"github.com/mikespook/gearman-go/worker"	
	"github.com/mikespook/gearman-go/client"	
	"log"
)

// this can be created multiple times 
func CreateWorker() * worker.Worker {

	// make sure application is bootstrapped and config is ready
	Bootstrap()
	
	// now lets grab all of the hosts that we need for this application
	hosts, err := Config.Get("gearmanHosts").Array()
	
	// handler gearman error as it comes about
	if err != nil {

		log.Println("Invalid configuration for gearmanHosts")
	}

	// register a new gearman worker
	w := worker.New(worker.Unlimited)

	// now loop through and add each host
	for _, host := range hosts {
	
		// assert that the host is a string and add it to the gearmanlist 
		if str, ok := host.(string); ok {

			// assert that we pass in a string from the element
			w.AddServer(str)
		}
	}

	return w
}

// can create multiple clients for this particular application
// for now we need to let this be ... come back later on https://github.com/mikespook/gearman-go/blob/master/example/client/client.go
func CreateClient() *client.Client {

	// ensure that application is booted up and ready to go
	Bootstrap()

	// initialize first host
	hosts, _ := Config.Get("gearmanHosts").Array()

	// lets make sure we can safely convert the host to a string 
	if str, ok := hosts[0].(string); ok {

		// now initialize a client
		c, _ := client.New(str)
		
		// return our client
		return c
	}

	// error handler etc for this functionality
	return nil
}


