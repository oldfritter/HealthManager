package tasks

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/oldfritter/HealthManager/initializers"
	. "github.com/oldfritter/HealthManager/models"
)

func ComponentList() {
	main := initializers.RabbitMqConnects["main"]
	main.DeclareExchange(initializers.AmqpGlobalConfigs["main"].Exchange["components"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["components"]["kind"], true, false, false, false, amqp.Table{})

	var allComponents []Component
	for _, v := range AllComponents {
		allComponents = append(allComponents, v)
	}
	b, err := json.Marshal(allComponents)
	if err != nil {
		log.Println(err)
	}
	main.PublishMessageWithRouteKey(initializers.AmqpGlobalConfigs["main"].Exchange["components"]["name"], "#", "text/plain", &b, amqp.Table{}, amqp.Persistent)

}
