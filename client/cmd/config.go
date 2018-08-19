package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var (
	Host string
	Port int
)

type Config struct {
	host string
	port int
}

func NewConfig() (*Config, error) {

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 5555)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	cfg := &Config{
		host: viper.Get("host").(string),
		port: viper.Get("port").(int),
	}

	err := viper.ReadInConfig()
	if err != nil {
		cfg.FlushToFile()
	}

	return cfg, nil
}

func (cfg *Config) SetHost(host string) {
	cfg.host = host
}

func (cfg *Config) SetPort(port int) {
	cfg.port = port
}

func (cfg *Config) FlushToFile() {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		panic(errors.New("baam!"))
	}
	ioutil.WriteFile("./config.yml", bs, 0644)
}

func init() {
	configCmd.Flags().StringVarP(&Host, "host", "H", "localhost", "the host to connect to")
	configCmd.Flags().IntVarP(&Port, "port", "p", 5555, "the port to connect to")
	configCmd.MarkFlagRequired("host")
	configCmd.MarkFlagRequired("port")
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("host", Host)
		viper.Set("port", Port)

		cfg, _ := NewConfig()
		cfg.SetHost(Host)
		cfg.SetPort(Port)
		cfg.FlushToFile()
	},
}
