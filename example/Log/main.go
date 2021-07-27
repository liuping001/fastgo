// Author: coolliu
// Date: 2021/7/27

package main

import (
	"fmt"
	"github.com/liuping001/fastgo/config"
	"github.com/liuping001/fastgo/log"
)

func main() {
	err := config.Init("./config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.Close()

	log.Infof("log start")
	log.Infof("log end")
}
