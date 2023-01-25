package configs

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)

type Database struct {
	Driver string `json:"driver"     yaml:"driver"`
	DSN    string `json:"dsn"        yaml:"dsn"`
}

type Redis struct {
	Addr     string `json:"addr"     yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db"       yaml:"db"`
}

// Config s
type Config struct {
	Secret   string   `json:"secret"   yaml:"secret"`
	Host     string   `json:"host"     yaml:"host"`
	Port     string   `json:"port"     yaml:"port"`
	Database Database `json:"database" yaml:"database"`
	Redis    Redis    `json:"redis"    yaml:"redis"`
}

var conf Config

func fromEnvfallbackString(value string, zero string, some string) string {
	val := os.Getenv(zero)
	if val != "" {
		return val
	}
	if value != "" {
		return value
	}
	return some
}

func fromEnvfallbackInt(value int, zero string, some int) int {
	val := os.Getenv(zero)
	if val != "" {
		v, err := strconv.Atoi(val)
		if err == nil {
			return v
		}
	}
	if value != 0 {
		return value
	}
	return some
}

var configPath string = ""

func getConfigArgs() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configPathDefault := path.Join(dir, "config.yml")
	flag.StringVar(&configPath, "c", configPathDefault, "Archivo de configuracion")
	flag.Parse()
}

// Init s
func Init() {
	getConfigArgs()
	f, err := os.Open(configPath)
	if err == nil {
		err = yaml.NewDecoder(f).Decode(&conf)
		if err == nil {
			return
		}
	}

	_ = godotenv.Load()

	conf.Host = fromEnvfallbackString(conf.Host, "HOST", "")
	conf.Port = fromEnvfallbackString(conf.Port, "PORT", "1323")
	conf.Secret = fromEnvfallbackString(conf.Secret, "SECRET", "123456")

	conf.Redis.Addr = fromEnvfallbackString(conf.Redis.Addr, "REDIS_ADDR", "localhost:6379")
	conf.Redis.DB = fromEnvfallbackInt(conf.Redis.DB, "REDIS_DB", 0)
	conf.Redis.Password = fromEnvfallbackString(conf.Redis.Password, "REDIS_PWD", "")

	conf.Database.DSN = fromEnvfallbackString(conf.Database.DSN, "DATABASE_DSN", "")
	conf.Database.Driver = fromEnvfallbackString(conf.Database.Driver, "DATABASE_DRIVER", "")
}

// GetConfig s
func GetConfig() Config {
	return conf
}
