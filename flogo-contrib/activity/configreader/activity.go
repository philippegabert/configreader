package configreader

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/creamdog/gonfig"
	"os"
	"sync"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-configreader")


const (
	configFile = "configFile"
)

type ConfigReader struct {
	sync.Mutex
	metadata *activity.Metadata
	config map[string]string
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ConfigReader{metadata: metadata, config: make(map[string]string)}
}

// Metadata implements activity.Activity.Metadata
func (a *ConfigReader) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ConfigReader) Eval(context activity.Context) (done bool, err error)  {

	configFile := context.GetInput(configFile).(string)
	log.Debugf("Config [%s]", configFile)

	f, err := os.Open(configFile)
	if err != nil {
		log.Error("Error while opening file ! ", err)
	}
	defer f.Close();
	config, err := gonfig.FromJson(f)
	if err != nil {
    	log.Error("Error while reading configuration file ! ", err)
    }

    confValue, err := config.GetString("test_config", "default")
	if err != nil {
		log.Error("Error while getting configuration value ! ", err)
	}

	log.Info("Final value: ",confValue)
	return true, nil
}
