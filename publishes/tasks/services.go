package tasks

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/oldfritter/HealthManager/initializers"
	. "github.com/oldfritter/HealthManager/models"
)

func ServiceList() {
	main := initializers.RabbitMqConnects["main"]
	main.DeclareExchange(initializers.AmqpGlobalConfigs["main"].Exchange["services"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["services"]["kind"], true, false, false, false, amqp.Table{})

	b, err := json.Marshal(AllServices)
	if err != nil {
		log.Println(err)
	}
	main.PublishMessageWithRouteKey(initializers.AmqpGlobalConfigs["main"].Exchange["services"]["name"], "#", "text/plain", &b, amqp.Table{}, amqp.Persistent)

}
