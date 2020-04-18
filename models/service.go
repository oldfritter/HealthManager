package models

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var AllServices = make(map[string]Service)

type Service interface {
	Symbol() string
}

var (
	mysqlDBs  = make(map[string]MysqlDB)
	redisDBs  = make(map[string]RedisDB)
	rabbitmqs = make(map[string]Rabbitmq)
)

func InitializeMysqlDBs() {
	pathStr, _ := filepath.Abs("config/mysql.yml")
	content, err := ioutil.ReadFile(pathStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &mysqlDBs)
	if err != nil {
		log.Fatal(err)
		return
	}
	for k, m := range mysqlDBs {
		m.Name = k
		AllServices[m.Symbol()] = m
	}
}

type MysqlDB struct {
	Name     string `json:"-"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbargs   string `json:"dbargs"`
	Pool     string `json:"pool"`
}

func (m MysqlDB) Symbol() string {
	return "mysql-" + m.Name
}

func InitializeRedisDBs() {
	pathStr, _ := filepath.Abs("config/redis.yml")
	content, err := ioutil.ReadFile(pathStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &redisDBs)
	if err != nil {
		log.Fatal(err)
		return
	}
	for k, m := range redisDBs {
		m.Name = k
		AllServices[m.Symbol()] = m
	}
}

type RedisDB struct {
	Name string `json:"-"`
	Host string `json:"host"`
	Port string `json:"port"`
	Db   string `json:"db"`
	Pool string `json:"pool"`
}

func (r RedisDB) Symbol() string {
	return "redis-" + r.Name
}

func InitializeRabbitmqs() {
	pathStr, _ := filepath.Abs("config/amqp.yml")
	content, err := ioutil.ReadFile(pathStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &rabbitmqs)
	if err != nil {
		log.Fatal(err)
		return
	}
	for k, m := range rabbitmqs {
		m.Name = k
		AllServices[m.Symbol()] = m
	}
}

type Rabbitmq struct {
	Name    string `json:"-"`
	Connect struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Vhost    string `yaml:"vhost"`
	} `yaml:"connect"`
}

func (r Rabbitmq) Symbol() string {
	return "rabbitmq-" + r.Name
}
