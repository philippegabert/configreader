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

func (a *ConfigReader) getConfig(configFile string, reachEachTime bool, configName string, configType string) string{
	a.Lock()
	defer a.Unlock()

	log.Debug("Variable readEachTime = ", reachEachTime)
	if reachEachTime || a.gonfigConf == nil {
		log.Debug("Need to read the configuration file...")
		f, err := os.Open(configFile)
		if err != nil {
			log.Error("Error while opening file ! ", err)
		}
		defer f.Close();
		a.gonfigConf, err = gonfig.FromJson(f)
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
	log.Debugf("Config file [%s]", configFile)

	var readEachTimeB bool

	if context.GetInput(readEachTime) != nil {
		log.Debug("Variable readEachTime is not null. (Value = "+context.GetInput(readEachTime)+")")
		readEachTimeB = context.GetInput(readEachTime).(bool)
		log.Debug("Variable readEachTimeB = ", readEachTimeB)
	}

	log.Debug("Getting config value...")
	configValue := a.getConfig(configFile, readEachTimeB, "test_config", "string")
	log.Debugf("Final value returned [%s]", configValue)

	return true, nil
}
