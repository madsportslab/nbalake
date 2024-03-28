package nbalake

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


func GetServerAddr() string {
	return fmt.Sprintf("%s:%s", NBALAKE_SERVICE,
    viper.GetString(NBALAKE_PORT))
} // GetServerAddr


func parseConfig() {

	viper.AutomaticEnv()
	viper.SetEnvPrefix(NBALAKE)
	
  c := viper.Get(NBALAKE_CONF_DIR)

	viper.SetConfigName(CONF_DEFAULT_FILE)
	viper.SetConfigType(CONF_DEFAULT_FILE_TYPE)
	
	if c == nil {
		viper.AddConfigPath(CONF_DEFAULT_DIR)
	} else {
		viper.AddConfigPath(c.(string))
	}
	
	viper.SetDefault(NBALAKE_HOST, NBALAKE_API_DEFAULT_HOST)
	viper.SetDefault(NBALAKE_PORT, NBALAKE_API_DEFAULT_PORT)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

} // parseConfig
