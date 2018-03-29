package service


import (
	"io/ioutil"
	"log"

	"go.uber.org/zap"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
)

// CommonConfig Common
type CommonConfig struct {
	Version  string
	IsDebug  bool
	LogLevel string
	LogPath  string
}


// EchoConf echo config struct
type IrisConf struct {
	Addr string
	Fromaddr string
	Toaddr string
	Serviceaddr string
	Port int
	Sub string
	Bodytitle string
	Pwd string
}



// Config ...
type Config struct {
	Common *CommonConfig
	IrisC  *IrisConf
}

// Conf ...
var Conf = &Config{}

// LoadConfig ...
func LoadConfig() {
	// init the new config params
	initConf()

	contents, err := ioutil.ReadFile("service.toml")
	if err != nil {
		log.Fatal("[FATAL] load service.toml: ", err)
	}
	tbl, err := toml.Parse(contents)
	if err != nil {
		log.Fatal("[FATAL] parse service.toml: ", err)
	}
	// parse common config
	parseCommon(tbl)
	// init log
	InitLogger()

	// parse Echo config
	parseIris(tbl)

	Logger.Info("LoadConfig", zap.Any("Config", Conf))
}

func initConf() {
	Conf = &Config{
		Common: &CommonConfig{},
		IrisC:  &IrisConf{},
	}
}

func parseCommon(tbl *ast.Table) {
	if val, ok := tbl.Fields["common"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Common)
		if err != nil {
			log.Fatalln("[FATAL] parseCommon: ", err, subTbl)
		}
	}
}


func parseIris(tbl *ast.Table) {
	if val, ok := tbl.Fields["ech"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.IrisC)
		if err != nil {
			log.Fatalln("[FATAL] parseEcho: ", err, subTbl)
		}
	}
}

