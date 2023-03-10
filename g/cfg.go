package g

import (
	"log"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/toolkits/file"
)


type Http struct {
	Enabled  bool
	HOST     string
	PORT     int64
	Timeout  int
}

type Auth struct {
	Username string
        Password string
}

type Global struct {
	Debug    bool
	Http     *Http
	Auth     *Auth
}

var (
	ConfigFile string
	config     *Global
	lock       = new(sync.RWMutex)
)

func Config() *Global {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.toml cfg.toml`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var g Global
	_, err = toml.Decode(configContent, &g)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()
	config = &g
	log.Println("read config file:", cfg, "successfully")
}
