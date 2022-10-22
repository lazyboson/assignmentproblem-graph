package main

import (
	"encoding/json"
	"fmt"
	"github.com/lazyboson/assignmentproblem/Agents"
	"github.com/lazyboson/assignmentproblem/graph"
	"github.com/lazyboson/assignmentproblem/queue/consumer"
	"github.com/lazyboson/assignmentproblem/queue/producer"
	"github.com/lazyboson/assignmentproblem/tasks"
	"log"
	"math/rand"
	"time"
)

const (
	queueName  = "QID_MTIzX3NhbGVzX3F1ZXVlNQ=="
	exchange   = "prod-exchange"
	tag        = queueName + "_ConsumerTag"
	routingKey = queueName + "_rKey"
)

func main() {
	q := producer.NewMQProducer("amqp://guest:guest@localhost:5672/", exchange, "direct")
	q.Start()

	err := q.CreateQueue(queueName, 9)
	if err != nil {
		log.Printf("failed to create queue: %+v", err)
		return
	}
	// producer of rabbit mq
	for i := 0; i < 100; i++ {
		taskData := tasks.GenerateTask()
		body, err := json.Marshal(taskData)
		if err != nil {
			log.Printf("error while marshalling json: %+v", err)
		}
		priority := rand.Intn(9-0) + 0
		err = q.PublishMessage(routingKey, body, uint8(priority))
		if err != nil {
			log.Printf("failed to publish message to MQ: %+v", err.Error())
			continue
		}
	}

	// consumer of RM
	c := consumer.NewMQConsumer("amqp://guest:guest@localhost:5672/", exchange, tag, queueName, routingKey)
	c.Start()

	// get agents
	agents := Agents.GenerateAgents()

	//g.HopcroftKarp()
	time.Sleep(2 * time.Second)
	taskList := c.TaskList
	g := graph.InitGraph(len(agents), len(taskList))

	//building graph
	for i := 0; i < len(taskList); i++ {
		for j := 0; j < len(agents); j++ {
			// there is an edge if skill matches
			if taskList[i].TaskExpression == agents[j].AgentExpression {
				g.AddEdges(j, i)
			}
		}
	}

	matchingEdges := g.HopcroftKarp()
	for key, val := range matchingEdges {
		fmt.Printf("agent: %s assigned the task: %s", agents[key].AgentId, taskList[val].Tid)
		fmt.Println()
	}
}
