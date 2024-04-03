package component

import (
	"fmt"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

func GetRabbitMQConnection(cnf *config.Config) *amqp091.Channel {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s%s", cnf.RabbitMQ.Username, cnf.RabbitMQ.Password, cnf.RabbitMQ.Host, cnf.RabbitMQ.Port, cnf.RabbitMQ.User)
	connection, err := amqp091.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}

	rmqChannel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return rmqChannel
}
