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
	Url             string `json:"url"               yaml:"url"`
	Endpoint        string `json:"endpoint"          yaml:"endpoint"`
	AccessKeyID     string `json:"access_key_id"     yaml:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl"           yaml:"use_ssl"`
	Bucket          string `json:"bucket"            yaml:"bucket"`
}

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

type Parser struct {
	prefix string
}

func (p *Parser) SetPrefix(prefix string) {
	p.prefix = prefix
}

func (p *Parser) Bool(value *bool, zero string, some bool) {
	val := os.Getenv(p.prefix + zero)
	if val != "" {
		*value = strings.ToLower(val) == "true"
		return
	}
	*value = some
}

func (p *Parser) String(value *string, zero string, some string) {
	val := os.Getenv(p.prefix + zero)
	if val != "" {
		*value = val
		return
	}
	if *value == "" {
		*value = some
	}
}

func (p *Parser) Int(value *int, zero string, some int) {
	val := os.Getenv(p.prefix + zero)
	if val != "" {
		v, err := strconv.Atoi(val)
		if err == nil {
			*value = v
			return
		}
	}
	if *value == 0 {
		*value = some
	}
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

	p := &Parser{}

	p.String(&conf.Host, "HOST", "")
	p.String(&conf.Port, "PORT", "1323")
	p.String(&conf.Secret, "SECRET", "123456")
	p.String(&conf.Env, "ENV", "development")
	p.Int(&conf.ReadBufferSize, "READ_BUFFER_SIZE", 2048)
	p.Int(&conf.WriteBufferSize, "WRITE_BUFFER_SIZE", 2048)

	p.SetPrefix("REDIS_")
	re := conf.Redis
	p.String(&re.Addr, "ADDR", "localhost:6379")
	p.Int(&re.DB, "DB", 0)
	p.String(&re.Password, "PWD", "")
	p.Int(&re.PoolSize, "POOL_SIZE", 50)

	p.SetPrefix("DATABASE_")
	db := conf.Database
	p.String(&db.DSN, "DSN", "")
	p.String(&db.Driver, "DRIVER", "")
	p.Int(&db.MaxOpenConns, "MAX_OPEN_CONNS", 50)
	p.Int(&db.MaxIdleConns, "MAX_IDLE_CONNS", 5)

	p.SetPrefix("MONGO_")
	mo := conf.Mongo
	p.String(&mo.DSN, "DSN", "")
	p.String(&mo.Database, "DATABASE", "")

	p.SetPrefix("CHATGPT3_")
	ch := conf.ChatGPT3
	p.String(&ch.ApiKey, "API_KEY", "")

	p.SetPrefix("MINIO_")
	mi := conf.Minio
	p.String(&mi.Url, "URL", "")
	p.String(&mi.Endpoint, "ENDPOINT", "")
	p.String(&mi.AccessKeyID, "ACCESS_KEY_ID", "")
	p.String(&mi.SecretAccessKey, "SECRET_ACCESS_KEY", "")
	p.Bool(&mi.UseSSL, "USE_SSL", false)
	p.String(&mi.Bucket, "BUCKET", "")
}

func GetConfig() Config {
	return conf
}
