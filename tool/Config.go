package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Drive     string `json:"drive"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	DbName    string `json:"db_name"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parse_time"`
	Loc       string `json:"loc"`
	Debug     bool   `json:"debug"`
	Singular  bool   `json:"singular"`
}

var _cfg *Config

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		return nil, err
	}
	return _cfg, err
}
