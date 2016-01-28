package rest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/couchbaselabs/logg"
)

//Config is the serialized config.json file
type Config struct {
	SyncEndpoint string `json:"syncEndpoint"`
	Port         int    `json:"port"`
	Bucket       string `json:"bucket"`
	TTL          int    `json:"ttl"`
}

var (
	configFileDescription = "The name of the config file.  Defaults to 'config.json'"
	configFileName        = kingpin.Arg("config file name", configFileDescription).Default("config.json").String()
	config                Config
	logTag                = "SYNC_AUTH"
)

func init() {
	// parse config file
	kingpin.Parse()
	if *configFileName == "" {
		kingpin.Errorf("Config file name missing")
		return
	}
	configFile, err := os.Open(*configFileName)
	if err != nil {
		logg.LogPanic("Unable to open file: %v.  Err: %v", *configFileName, err.Error())
		return
	}
	defer configFile.Close()

	configReader := bufio.NewReader(configFile)
	parseConfigFile(configReader)

	//set logging
	logg.LogKeys[logTag] = true
}

func parseConfigFile(r io.Reader) {
	config = Config{}

	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&config); err != nil {
		logg.LogPanic("Error parsing config file: %v", err)
	}
}

//StartServer ...
func StartServer() {
	logg.LogTo(logTag, "Listening on port :%d", config.Port)
	router := createRouter()

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
