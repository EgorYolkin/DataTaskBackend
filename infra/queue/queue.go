package queue

import (
	"DataTask/internal/config"
	"DataTask/pkg/logger"
	"fmt"
)

func SetupPublisher(cfg config.Config) *RabbitMQ {
	ampqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Pass,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	logger.Log.Info(fmt.Sprintf("AMQP URL: %s", ampqUrl))
	rabbitmq, err := NewRabbitMQ(ampqUrl)

	if err != nil {
		logger.Log.Errorf("Error creating rabbitmq %v", err)
		return &RabbitMQ{}
	}
	return rabbitmq
}

func InitConsumers(publisher *RabbitMQ, config RabbitMQConfig) {
	err := initPublisher(publisher, config)

	logger.Log.Info("initialize eda consumers")

	//usersEventHandler := users_event_handler.NewUsersEventHandler(nil)

	if err != nil {
		logger.Log.Fatal(err)
	}

	msgs, _ := publisher.Consume(config.QueueName)
	go func() {
		for msg := range msgs {
			//err = usersEventHandler.CreateUserEventHandler(context.Background(), msg.Body)

			if err != nil {
				logger.Log.Info("error processing event: %v", err)
			}

			_ = msg.Ack(false)
		}
	}()

	select {}
}

func initPublisher(publisher *RabbitMQ, config RabbitMQConfig) error {
	err := publisher.DeclareExchange(config.ExchangeName, config.ExchangeType)
	if err != nil {
		return err
	}
	err = publisher.DeclareQueue(config.QueueName)
	if err != nil {
		return err
	}
	err = publisher.BindQueue(config.QueueName, config.ExchangeName, config.RoutingKey)
	if err != nil {
		return err
	}

	return nil
}
