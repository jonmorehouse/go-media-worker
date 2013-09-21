/*
	Bootstrap package

	1.) Responsible for booting up configuration
	2.) Responsible for dumb callers. We don't want to call the bootstrap function more than once etc
*/

package bootstrap

import (
	
	"fmt"
	"log"
	"path"
	"os"
	"io/ioutil"
	"sync" 
	//"github.com/bitly/go-simplejson"
)

// manage global once element for controlling the runs of bootstrap function
var once sync.Once

// initialize a global config string to be exported 
var Config = map[string]string {

	"configDir": "config",
	"configFile": "config.json", 
}

// now lets initialize the correct json cnofiguration etc
// we are going to initialize a struct to wrap the function we are calling to ensure that we don't call and boot the application twice
// we want to export this etc
//var Once sync.Once
// bootstrap is responsible for initializing application and setting up server elements etc
func bootstrap () {

	// grab the gopath environment variable
	goPath := os.Getenv("GOPATH")

	// make sure we have a go element
	if goPath == "" {

		log.Fatal("Invalid GOPATH")
	}

	// now lets make sure the config file exists etc
	configPath := path.Join(goPath, Config["configDir"], Config["configFile"])

	// now call a fatal error if config file doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
	
		log.Fatalf("Config file doesn't exist at %s", configPath)
	}
	
	// now lets read the file and work a little magic!
	rawJson, err := ioutil.ReadFile(configPath)

	// check to make sure there wasn't a read error
	if err != nil {

		log.Fatalf("Could not read config file properly %s", configPath)
	}

	// create a new json structure from the configuration
	//jsonData, err := simplejson.NewJson(rawJson)
	
	// now that we have created our config, lets link it to a global variable in this element
	// note that if this function is called multiple times, then we will have extra memory overhead ... we could have a basic run tim

	fmt.Println(string(rawJson))

}


func Bootstrap() {	
	
	once.Do(bootstrap)

}


