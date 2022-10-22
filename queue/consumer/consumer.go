package consumer

import (
	"encoding/json"
	"github.com/lazyboson/assignmentproblem/tasks"
	"github.com/streadway/amqp"
	"log"
)

type MQConsumer struct {
	amqpURL    string
	exchange   string
	tag        string
	queueName  string
	bindingKey string
	conn       *amqp.Connection
	TaskList   []*tasks.Task
}

func NewMQConsumer(amqpURI, exchange, tag, queueName, bindingKey string) *MQConsumer {
	c := &MQConsumer{
		amqpURL:    amqpURI,
		exchange:   exchange,
		tag:        tag,
		queueName:  queueName,
		bindingKey: bindingKey,
	}
	c.TaskList = make([]*tasks.Task, 0)
	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Fatalf("error:: amqp dial: %+v", err)
	}

	c.conn = connection
	return c
}

func (c *MQConsumer) Start() {
	channel, err := c.conn.Channel()
	if err != nil {
		log.Fatalf("error:: getting channel: %+v \n", err)
	}

	if err = channel.QueueBind(
		c.queueName,  // name of the queue
		c.bindingKey, // bindingKey
		c.exchange,   // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		log.Fatalf("error:: creating binding: %+v \n", err)
	}

	// set prefect = 1
	err = channel.Qos(1, 0, false)
	if err != nil {
		log.Fatalf("error:: set basic.qos: %v", err)
	}

	deliveries, err := channel.Consume(
		c.queueName, // name
		c.tag,       // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("error:: consumer start: %v", err)
	}

	go c.handleMessages(deliveries)
}

func (c *MQConsumer) handleMessages(deliveries <-chan amqp.Delivery) {
	for message := range deliveries {
		task := &tasks.Task{}
		err := json.Unmarshal(message.Body, task)
		if err != nil {
			log.Println("error:: ", err)
		}
		c.TaskList = append(c.TaskList, task)
		message.Ack(false)
	}

}
