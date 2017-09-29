package def

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"log"
	"runtime"
)

type Config struct {

}
func (c *Config) InitConfig(dir string,confName string,topic string) map[string]string {
	var (
		configFile = flag.String(dir, confName, "General configuration file")
	)
	//var topic string  = "nats"
	//topic list
	var TOPIC = make(map[string]string)
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection(topic) {
		section, err := cfg.SectionOptions(topic)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(topic, v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	fmt.Println(TOPIC)
	//fmt.Println(TOPIC["q_name"])
	return TOPIC
}