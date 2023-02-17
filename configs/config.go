package configs

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)

type MongoDB struct {
	DSN          string `json:"dsn"          yaml:"dsn"          valid:"required"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns" valid:"required"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns" valid:"required"`
}

type Database struct {
	Driver       string `json:"driver"       yaml:"driver"       valid:"required"`
	DSN          string `json:"dsn"          yaml:"dsn"          valid:"required"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns" valid:"required"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns" valid:"required"`
}

type Redis struct {
	Addr     string `json:"addr"     yaml:"addr"     valid:"required"`
	Password string `json:"password" yaml:"password" valid:""`
	DB       int    `json:"db"       yaml:"db"       valid:""`
	PoolSize int    `json:"poolSize" yaml:"poolSize" valid:"required"`
}

type Mongo struct {
	DSN            string `json:"dsn"             yaml:"dsn"             valid:"required"`
	StorageDB      string `json:"storage_db"      yaml:"storage_db"      valid:"required"`
	ConversationDB string `json:"conversation_db" yaml:"conversation_db" valid:"required"`
}

type ChatGPT3 struct {
	ApiKey string `json:"api_key" yaml:"api_key" valid:"required"`
}

type Minio struct {
	Url             string `json:"url"               yaml:"url"               valid:"required"`
	Endpoint        string `json:"endpoint"          yaml:"endpoint"          valid:"required"`
	AccessKeyID     string `json:"access_key_id"     yaml:"access_key_id"     valid:"required"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key" valid:"required"`
	UseSSL          bool   `json:"use_ssl"           yaml:"use_ssl"           valid:""`
	Bucket          string `json:"bucket"            yaml:"bucket"            valid:"required"`
}

type Config struct {
	Env             string    `json:"env"               yaml:"env"               valid:"required"`
	Secret          string    `json:"secret"            yaml:"secret"            valid:"required"`
	Host            string    `json:"host"              yaml:"host"              valid:"required"`
	Port            string    `json:"port"              yaml:"port"              valid:"required"`
	ReadBufferSize  int       `json:"read_buffer_size"  yaml:"read_buffer_size"  valid:"required"`
	WriteBufferSize int       `json:"write_buffer_size" yaml:"write_buffer_size" valid:"required"`
	Database        *Database `json:"database"          yaml:"database"          valid:"required"`
	Redis           *Redis    `json:"redis"             yaml:"redis"             valid:"required"`
	Mongo           *Mongo    `json:"mongo"             yaml:"mongo"             valid:"required"`
	ChatGPT3        *ChatGPT3 `json:"chat_gpt"          yaml:"chat_gpt"          valid:"required"`
	Minio           *Minio    `json:"minio"             yaml:"minio"             valid:"required"`
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

func loadFromYaml() error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		return err
	}
	return nil
}

func loadFromEnv() error {
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
	p.String(&mo.StorageDB, "STORAGE_DB", "")
	p.String(&mo.ConversationDB, "CONVERSATION_DB", "")

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

	return nil
}

func init() {
	getConfigArgs()
	err := loadFromYaml()
	if err != nil {
		loadFromEnv()
	}
	_, err = govalidator.ValidateStruct(conf)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() Config {
	return conf
}
