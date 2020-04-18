package tasks

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"

	"github.com/oldfritter/HealthManager/initializers"
	. "github.com/oldfritter/HealthManager/models"
)

func ComponentList() {
	main := initializers.RabbitMqConnects["main"]
	main.DeclareExchange(initializers.AmqpGlobalConfigs["main"].Exchange["components"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["components"]["kind"], true, false, false, false, amqp.Table{})

	var allComponents []Component
	t := time.Now().Add(-time.Second * 30).Unix()
	for _, v := range AllComponents {
		if v.Timestamp > t {
			allComponents = append(allComponents, v)
		}
	}
	var b []byte
	var err error
	if len(allComponents) == 0 {
		b = []byte("[]")
	} else {
		b, err = json.Marshal(allComponents)
		if err != nil {
			log.Println(err)
		}
	}
	main.PublishMessageWithRouteKey(initializers.AmqpGlobalConfigs["main"].Exchange["components"]["name"], "#", "text/plain", &b, amqp.Table{}, amqp.Persistent)

}
