package initializers

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v2"

	"github.com/oldfritter/HealthManager/utils"
)

type Amqp struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`

	Exchange map[string]map[string]string `yaml:"exchange"`
}

var (
	AmqpGlobalConfigs = make(map[string]Amqp)
	RabbitMqConnects  = make(map[string]utils.RabbitMqConnect)
)

func InitializeAmqpConfig() {
	pathStr, _ := filepath.Abs("config/amqp.yml")
	content, err := ioutil.ReadFile(pathStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &AmqpGlobalConfigs)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func InitializeAmqpConnection(name string) {
	var err error
	conn, err := amqp.Dial("amqp://" + AmqpGlobalConfigs[name].Username + ":" + AmqpGlobalConfigs[name].Password + "@" + AmqpGlobalConfigs[name].Host + ":" + AmqpGlobalConfigs[name].Port + "/" + AmqpGlobalConfigs[name].Vhost)
	RabbitMqConnects[name] = utils.RabbitMqConnect{conn}
	if err != nil {
		log.Fatal("rabbimq connect error: %v", err)
		time.Sleep(5000)
		InitializeAmqpConnection(name)
		return
	}
	go func() {
		<-RabbitMqConnects[name].NotifyClose(make(chan *amqp.Error))
		InitializeAmqpConnection(name)
	}()
}

func CloseAmqpConnection(name string) {
	RabbitMqConnects[name].Close()
}

func GetRabbitMqConnect(name string) utils.RabbitMqConnect {
	return RabbitMqConnects[name]
}
