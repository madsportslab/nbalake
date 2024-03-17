package nbalake

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


func GetServerAddr() string {
	return fmt.Sprintf("%s:%s", viper.GetString(BLOB_HOST),
    viper.GetString(BLOB_PORT))
} // GetServerAddr


func parseConfig() {

	viper.SetConfigName(CONF_DEFAULT_FILE)
	viper.SetConfigType(CONF_DEFAULT_FILE_TYPE)
	viper.AddConfigPath(CONF_DEFAULT_DIR)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

} // parseConfig
