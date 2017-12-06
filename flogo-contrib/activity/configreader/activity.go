package configreader

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/creamdog/gonfig"
	"os"
	"sync"
	"strconv"
	"fmt"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-configreader")


const (
	configFile = "configFile"
	readEachTime = "readEachTime"
	configName = "configName"
	configValue = "configValue"
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

func (a *ConfigReader) getConfig(configFile string, reachEachTime bool, confName string, configType string) interface{}{
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

    confValue, err := a.gonfigConf.GetString(confName, "default")
	if err != nil {
		log.Error("Error while getting configuration value ! ", err)
	}

	log.Info("Final value: ",confValue)

	return confValue
}

func toBool(val interface{}) (bool, error) {

	b, ok := val.(bool)
	if !ok {
		s, ok := val.(string)

		if !ok {
			return false, fmt.Errorf("Unable to convert to boolean")
		}

		var err error
		b, err = strconv.ParseBool(s)

		if err != nil {
			return false, err
		}
	}

	return b, nil
}

// Eval implements activity.Activity.Eval
func (a *ConfigReader) Eval(context activity.Context) (done bool, err error)  {

	configFile := context.GetInput(configFile).(string)
	log.Debugf("Config file [%s]", configFile)

	var readEachTimeB bool
	var configurationName string

	if context.GetInput(readEachTime) != nil {
		log.Debug("Variable readEachTime is not null.")
		readEachTimeB, _ = toBool(context.GetInput(readEachTime))
	}
	if context.GetInput(configName) != nil {
		log.Debugf("Configuration name [%s]", configName)
		configurationName = context.GetInput(readEachTime).(string)
		log.Debug("Getting config value...")
		confValue := a.getConfig(configFile, readEachTimeB, configurationName, "string")
		log.Debugf("Final value returned [%s]", confValue)

		context.SetOutput(configValue, confValue)

		return true, nil
	} else {
		return false, fmt.Errorf("No configuration name has been set !")
	}


	
}
