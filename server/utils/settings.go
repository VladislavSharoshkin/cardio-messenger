package utils

import (
	"awesomeProject/crypto"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const fileName = "config.json"

type settings struct {
	Version             int64
	UserStaffPermission bool
	Debug               bool
	Registration        bool
	Database            string
	Ip                  string
	PushServer          string
	Fingerprint         string
}

var Settings settings

func setDefaults() {
	viper.SetDefault("Debug", true)
	viper.SetDefault("User_Staff_Permission", false)
	viper.SetDefault("Registration", true)
	viper.SetDefault("Database", "host=127.0.0.1 user=postgres password=postgres port=5432 dbname=postgres sslmode=disable")
	viper.SetDefault("Ip", "62.183.103.210")
	viper.SetDefault("Push_Server", "62.183.103.210:27991")
}

func setStruct() {
	fp, _ := crypto.CertFingerprint("cert.pem")

	Settings = settings{
		Version:             1,
		Debug:               viper.GetBool("Debug"),
		UserStaffPermission: viper.GetBool("User_Staff_Permission"),
		Registration:        viper.GetBool("Registration"),
		Database:            viper.GetString("Database"),
		Ip:                  viper.GetString("Ip"),
		PushServer:          viper.GetString("Push_Server"),
		Fingerprint:         fp,
	}
}

func createFile() {
	f, err := os.Create(fileName)
	if err != nil {
		panic("error create config file")
	}
	if _, err := f.WriteString("{}"); err != nil {
		panic("error write {} in config file")
	}
}

func (thisObject settings) GetPublic() settings {
	return settings{
		Version:             thisObject.Version,
		UserStaffPermission: thisObject.UserStaffPermission,
		Registration:        thisObject.Registration,
	}
}

func init() {
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createFile()
			if err := viper.ReadInConfig(); err != nil {
				panic("error read config file")
			}
		} else {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}

	}
	viper.WriteConfig()
	setStruct()

	if Settings.Debug {
		Print("debug mode", PrintTypeNormal)
	} else {
		Print("publish mode", PrintTypeWarning)
	}
	//fmt.Printf("test %s", base58.Encode(crypto.Sha(crypto.RandomBytes256())))
	//fmt.Printf("server fingerprint %s", Settings.Fingerprint)
}
