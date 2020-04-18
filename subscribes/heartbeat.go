package subscribes

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"

	"github.com/oldfritter/HealthManager/initializers"
	. "github.com/oldfritter/HealthManager/models"
)

func HeartBeat() {
	main := initializers.RabbitMqConnects["main"]
	main.DeclareQueue(initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], amqp.Table{})
	main.DeclareExchange(initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["kind"], true, false, false, false, amqp.Table{})
	main.QueueBind(initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], false, amqp.Table{})
	go func() {
		channel, err := main.Channel()
		defer channel.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		msgs, err := channel.Consume(
			initializers.AmqpGlobalConfigs["main"].Exchange["heartbeat"]["name"], // queue
			"Health Manager", // consumer
			false,            // auto-ack
			false,            // exclusive
			false,            // no-local
			false,            // no-wait
			nil,              // args
		)

		for d := range msgs {
			err = d.Ack(treat(&d.Body))
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}()
	return
}

func treat(data *[]byte) (success bool) {
	var current Component
	err := json.Unmarshal(*data, &current)
	if err != nil {
		log.Println("err: ", err)
		return
	} else {
		log.Println("received heartbeat: ", current)
	}
	current.Timestamp = time.Now().Unix()
	AllComponents[current.Name] = current
	success = true
	return
}
