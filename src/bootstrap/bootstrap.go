package bootstrap

import (
	
	"fmt"
	"log"
	"path"
	"os"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"reflect"
)

// initialize a global config string to be exported 
var Config = map[string]string {

	"configDir": "config",
	"configFile": "config.json", 
}

// bootstrap is responsible for initializing application and setting up server elements etc
func Bootstrap() {
	
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
	jsonData, err := simplejson.NewJson(rawJson)
	
	test := jsonData.Get("gearmanHosts").MustArray()
		
	fmt.Println(jsonData.Get("gearmanHosts"))
	
	fmt.Println(reflect.TypeOf(test))
}


