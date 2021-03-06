package mono

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	data_files []map[string]string
}

func Init() error {
	viper.SetEnvPrefix("monorest")
	viper.SetConfigName("monorest")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Printf("error parsing config file: %v", e)
			return e
		}
		log.Printf("No config file used")
	} else {
		log.Printf("Using config file: %v", viper.ConfigFileUsed())
	}
	return nil
}

func NewConfig() *Config {
	config := Config{data_files: make([]map[string]string, 0)}

	data_files := viper.Get("monorest.data_files")
	if data_files != nil {
		for _, d := range data_files.([]interface{}) {
			data_file_i := d.(map[interface{}]interface{})
			data_file := make(map[string]string)
			for k, v := range data_file_i {
				data_file[k.(string)] = v.(string)
			}
			config.data_files = append(config.data_files, data_file)
		}
	}
	return &config
}

type configError struct {
	err string
}

func (e configError) Error() string {
	return e.err
}

func (c Config) FindDataFile(name string) (string, string, string, error) {
	for _, v := range c.data_files {
		if v["name"] == name {
			return name, v["type"], v["path"], nil
		}
	}
	e := fmt.Sprintf("Data file %s not found", name)
	return "", "", "", configError{err: e}
}

func (c Config) GetDataFiles() []string {
	data_files := make([]string, 0)
	for _, v := range c.data_files {
		data_files = append(data_files, v["name"])
	}
	return data_files
}
