package global

import (
	"goformat-v2/app/global/structs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Config 讀取dev.yaml檔案
var Config *structs.EnvConfig

// Lang 各語系
var Lang = []string{"en", "tw", "cn"}

func getEnv() string {
	if len(os.Getenv("ENV")) <= 0 {
		log.Fatalf("🔔🔔🔔  Can not get ENV value 🔔🔔🔔")
	}
	return os.Getenv("ENV")
}

// Start 執行main.go的第一步驟，載入各環境設定檔
func Start() (err error) {
	env := getEnv()

	envPathList := []string{
		"env/" + env + "/db.yaml",
		"env/" + env + "/api.yaml",
		"env/" + env + "/other.yaml",
	}

	for _, path := range envPathList {
		configFile, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("🔔🔔🔔  Can not find Yaml file %v 🔔🔔🔔", err)
		}

		// 塞值進入struct
		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			panic(err)
		}
	}

	return nil
}

