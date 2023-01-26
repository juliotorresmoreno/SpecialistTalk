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
	Driver       string `json:"driver"       yaml:"driver"`
	DSN          string `json:"dsn"          yaml:"dsn"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
}

type Redis struct {
	Addr     string `json:"addr"     yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db"       yaml:"db"`
}

// Config s
type Config struct {
	Env             string    `json:"env"             yaml:"env"`
	Secret          string    `json:"secret"          yaml:"secret"`
	Host            string    `json:"host"            yaml:"host"`
	Port            string    `json:"port"            yaml:"port"`
	ReadBufferSize  int       `json:"readBufferSize"  yaml:"readBufferSize"`
	WriteBufferSize int       `json:"writeBufferSize" yaml:"writeBufferSize"`
	Database        *Database `json:"database"        yaml:"database"`
	Redis           *Redis    `json:"redis"           yaml:"redis"`
}

var conf Config = Config{
	Database:        &Database{},
	Redis:           &Redis{},
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

func fromEnvfString(value string, zero string, some string) string {
	val := os.Getenv(zero)
	if val != "" {
		return val
	}
	if value != "" {
		return value
	}
	return some
}

func fromEnvfInt(value int, zero string, some int) int {
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
	_, err = os.Open(configPathDefault)
	if err != nil {
		return
	}
	flag.StringVar(&configPath, "c", configPathDefault, "configuration filepath")
	flag.Parse()
}

func init() {
	getConfigArgs()
	f, err := os.Open(configPath)
	if err == nil {
		err = yaml.NewDecoder(f).Decode(&conf)
		if err == nil {
			return
		}
	}

	_ = godotenv.Load()

	conf.Host = fromEnvfString(conf.Host, "HOST", "")
	conf.Port = fromEnvfString(conf.Port, "PORT", "1323")
	conf.Secret = fromEnvfString(conf.Secret, "SECRET", "123456")
	conf.Env = fromEnvfString(conf.Env, "ENV", "development")
	conf.ReadBufferSize = fromEnvfInt(conf.ReadBufferSize, "READ_BUFFER_SIZE", 2048)
	conf.WriteBufferSize = fromEnvfInt(conf.WriteBufferSize, "WRITE_BUFFER_SIZE", 2048)

	conf.Redis.Addr = fromEnvfString(conf.Redis.Addr, "REDIS_ADDR", "localhost:6379")
	conf.Redis.DB = fromEnvfInt(conf.Redis.DB, "REDIS_DB", 0)
	conf.Redis.Password = fromEnvfString(conf.Redis.Password, "REDIS_PWD", "")

	conf.Database.DSN = fromEnvfString(conf.Database.DSN, "DATABASE_DSN", "")
	conf.Database.Driver = fromEnvfString(conf.Database.Driver, "DATABASE_DRIVER", "")
	conf.Database.MaxOpenConns = fromEnvfInt(conf.Redis.DB, "MAX_OPEN_CONNS", 50)
	conf.Database.MaxIdleConns = fromEnvfInt(conf.Redis.DB, "MAX_IDLE_CONNS", 5)
}

// GetConfig s
func GetConfig() Config {
	return conf
}
