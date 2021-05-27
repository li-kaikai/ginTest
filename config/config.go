package config

import (
	"flag"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	Dft dfter
)

func init() {
	Dft = &dft{}
}

type dfter interface {
	Get() cfg
}

type dft struct {
	sync.Once
	dftCfg cfg
}

type cfg struct {
	Runmode string `yaml:"runmode"`
	Http    struct {
		Port int `yaml:"port"`
	} `yaml:"http"`
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Pass string `yaml:"password"`
	} `yaml:"redis"`
	Mysql struct {
		Order struct {
			Addr         string `yaml:"addr"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
		} `yaml:"order"`
	} `yaml:"mysql"`
}

func (this *dft) Get() cfg {
	this.Do(func() {
		path := ""
		flag.StringVar(&path, "conf", "./config/config.yml", "help")
		flag.Parse()
		bytes, err := ioutil.ReadFile(path)
		if nil != err {
			panic(err)
		}
		err = yaml.Unmarshal(bytes, &this.dftCfg)
		if nil != err {
			panic(err)
		}
	})
	return this.dftCfg
}
