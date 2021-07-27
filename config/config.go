// Author: coolliu
// Date: 2021/7/27

package config

import (
	"fmt"
	"github.com/liuping001/fastgo/log"
	"github.com/liuping001/fastgo/util"
	"go.uber.org/zap"
)

type Config struct {
	ZapLog *log.ZapLog `json:"zap_log" yaml:"zap_log"`
}

var gConfig Config
var zapLog *zap.SugaredLogger

func Init(file string) error {
	err := util.FileToYaml(file, &gConfig)
	if err != nil {
		return err
	}
	fmt.Printf("config:%s\n", util.ToJson(gConfig))
	if gConfig.ZapLog != nil {
		zapLog = log.NewZapLogger(gConfig.ZapLog)
		log.SetGlobalLog(zapLog)
	}
	return nil
}

func Close() {
	if zapLog != nil {
		zapLog.Sync()
	}
}
