package configs

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)

type MongoDB struct {
	DSN          string `json:"dsn"          yaml:"dsn"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
}

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
	PoolSize int    `json:"poolSize" yaml:"poolSize"`
}

type Mongo struct {
	DSN      string `json:"dsn"      yaml:"dsn"`
	Database string `json:"database" yaml:"database"`
}

type ChatGPT3 struct {
	ApiKey string `json:"api_key" yaml:"api_key"`
}

type Minio struct {
	Endpoint        string `json:"endpoint"          yaml:"endpoint"`
	AccessKeyID     string `json:"access_key_id"     yaml:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl"           yaml:"use_ssl"`
}

// Config s
type Config struct {
	Env             string    `json:"env"               yaml:"env"`
	Secret          string    `json:"secret"            yaml:"secret"`
	Host            string    `json:"host"              yaml:"host"`
	Port            string    `json:"port"              yaml:"port"`
	ReadBufferSize  int       `json:"read_buffer_size"  yaml:"read_buffer_size"`
	WriteBufferSize int       `json:"write_buffer_size" yaml:"write_buffer_size"`
	Database        *Database `json:"database"          yaml:"database"`
	Redis           *Redis    `json:"redis"             yaml:"redis"`
	Mongo           *Mongo    `json:"mongo"             yaml:"mongo"`
	ChatGPT3        *ChatGPT3 `json:"chat_gpt"          yaml:"chat_gpt"`
	Minio           *Minio    `json:"minio"             yaml:"minio"`
}

var conf Config = Config{
	Database:        &Database{},
	Redis:           &Redis{},
	Mongo:           &Mongo{},
	ChatGPT3:        &ChatGPT3{},
	Minio:           &Minio{},
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

func fromEnvfBool(zero string, some bool) bool {
	val := os.Getenv(zero)
	if val != "" {
		return strings.ToLower(val) == "true"
	}
	return some
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
		decoder := yaml.NewDecoder(f)
		if err = decoder.Decode(&conf); err == nil {
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
	conf.Redis.PoolSize = fromEnvfInt(conf.Redis.PoolSize, "REDIS_POOL_SIZE", 50)

	conf.Database.DSN = fromEnvfString(conf.Database.DSN, "DATABASE_DSN", "")
	conf.Database.Driver = fromEnvfString(conf.Database.Driver, "DATABASE_DRIVER", "")
	conf.Database.MaxOpenConns = fromEnvfInt(conf.Redis.DB, "DATABASE_MAX_OPEN_CONNS", 50)
	conf.Database.MaxIdleConns = fromEnvfInt(conf.Redis.DB, "DATABASE_MAX_IDLE_CONNS", 5)

	conf.Mongo.DSN = fromEnvfString(conf.Mongo.DSN, "MONGO_DSN", "")
	conf.Mongo.Database = fromEnvfString(conf.Mongo.Database, "MONGO_DATABASE", "")

	conf.ChatGPT3.ApiKey = fromEnvfString(conf.ChatGPT3.ApiKey, "CHATGPT3_API_KEY", "")

	conf.Minio.Endpoint = fromEnvfString(conf.ChatGPT3.ApiKey, "MINIO_ENDPOINT", "")
	conf.Minio.AccessKeyID = fromEnvfString(conf.ChatGPT3.ApiKey, "MINIO_ACCESS_KEY_ID", "")
	conf.Minio.SecretAccessKey = fromEnvfString(conf.ChatGPT3.ApiKey, "MINIO_SECRET_ACCESS_KEY", "")
	conf.Minio.UseSSL = fromEnvfBool("MINIO_USE_SSL", false)
}

// GetConfig s
func GetConfig() Config {
	return conf
}
