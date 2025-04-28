package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(amqpUrl string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(amqpUrl)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn: conn, ch: ch}, nil
}

type RabbitMQConfig struct {
	ExchangeName string
	ExchangeType string
	QueueName    string
	RoutingKey   string
}

func NewRabbitMQConfig(exchangeName, exchangeType, queueName, routingKey string) *RabbitMQConfig {
	return &RabbitMQConfig{
		ExchangeName: exchangeName, // users
		ExchangeType: exchangeType, // direct
		QueueName:    queueName,    // user_created_queue
		RoutingKey:   routingKey,   // user.created
	}
}

func (r *RabbitMQ) Publish(exchange, key string, body []byte) error {
	return r.ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	_, err := r.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := r.ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	return msgs, err
}

func (r *RabbitMQ) DeclareExchange(name, kind string) error {
	return r.ch.ExchangeDeclare(
		name,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (r *RabbitMQ) BindQueue(queue, exchange, key string) error {
	return r.ch.QueueBind(
		queue,
		key,
		exchange,
		false,
		nil,
	)
}

func (r *RabbitMQ) DeclareQueue(name string) error {
	_, err := r.ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (r *RabbitMQ) HandleReconnect(url string) {
	for {
		select {
		case amqpErr := <-r.conn.NotifyClose(make(chan *amqp.Error)):
			log.Printf("Connection lost: %v. Reconnecting...", amqpErr)
			newConn, err := NewRabbitMQ(url)
			if err == nil {
				*r = *newConn
				return
			}
			time.Sleep(5 * time.Second)
		}
	}
}

func (r *RabbitMQ) Close() {
	_ = r.conn.Close()
	_ = r.ch.Close()
}
