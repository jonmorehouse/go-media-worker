/*
	Bootstrap package

	@note assumes that the project path is always the last path on the $GOPATH
	@note will look in the project path to determine the actual config path

	1.) Responsible for booting up configuration
	2.) Responsible for dumb callers. We don't want to call the bootstrap function more than once etc
*/
package bootstrap

import (
	
	"log"
	"path"
	"os"
	"io/ioutil"
	"sync" 
	"strings"
	"github.com/bitly/go-simplejson"
)

// manage global once element for controlling the runs of bootstrap function
var once sync.Once

// initialize a global config string to be exported 
// initialize internal configuraetion for this module
var config = map[string]string {

	"configDir": "config",
	"configFile": "config.json", 
}

// initialize a config element for storing our json
var Config *simplejson.Json


// now lets initialize the correct json cnofiguration etc
// we are going to initialize a struct to wrap the function we are calling to ensure that we don't call and boot the application twice
// we want to export this etc
//var Once sync.Once
// bootstrap is responsible for initializing application and setting up server elements etc
func bootstrap () {

	// grab the gopath environment variable
	goPath := strings.Split(os.Getenv("GOPATH"), ":")[1:][0]

	// make sure we have a go element
	if goPath == "" {

		log.Fatal("Invalid GOPATH")
	}

	// now lets make sure the config file exists etc
	configPath := path.Join(goPath, config["configDir"], config["configFile"])

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
	jsonData, err := simplejson.NewJson(rawJson)
	
	// create a simple error handler
	if err != nil {
		
		log.Fatalf("Configuration invalid.")
	}

	// now lets copy over some of the values from our configuration element
	jsonData.Set("configPath", configPath)

	// if all is well, go ahead and save a pointer to the jsonData element to the configuration element
	Config = jsonData

}


// wrap our bootstrap and create an interface for outside callers to call 
// ensure that we only call the function once
func Bootstrap() {	
	
	once.Do(bootstrap)
}


