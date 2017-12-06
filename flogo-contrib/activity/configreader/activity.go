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
	readEachTime = "readEachTime"
)

type ConfigReader struct {
	sync.Mutex
	metadata *activity.Metadata
	gonfigConf gonfig.Gonfig
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ConfigReader{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ConfigReader) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *ConfigReader) getConfig(configFile string, reachEachTime bool, configName string, configType string) {
	a.Lock()
	defer a.Unlock()

	if reachEachTime || a.gonfigConf == nil {
		f, err := os.Open(configFile)
		if err != nil {
			log.Error("Error while opening file ! ", err)
		}
		defer f.Close();
		a.gonfigConf, err := gonfig.FromJson(f)
		if err != nil {
	    	log.Error("Error while reading configuration file ! ", err)
	    }
	}

    confValue, err := a.gonfigConf.GetString(configName, "default")
	if err != nil {
		log.Error("Error while getting configuration value ! ", err)
	}

	log.Info("Final value: ",confValue)

	return confValue
}

// Eval implements activity.Activity.Eval
func (a *ConfigReader) Eval(context activity.Context) (done bool, err error)  {

	configFile := context.GetInput(configFile).(string)
	log.Debugf("Config [%s]", configFile)

	if context.GetInput(readEachTime) != nil {
		readEachTimeB := context.GetInput(readEachTime).(bool)
	}

	configValue := a.getConfig(configFile, readEachTimeB, "test_config", "string")
	log.Debugf("Final value returned [%s]", configValue)

	return true, nil
}
