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
	for i := 0; i < 4; i++ {
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

	// get auxiliary constructed agents -- this should be driven by presence service
	agents := Agents.GenerateAgents()

	// sleeping here to get consumer data here as it is in separate goroutine -- not the best way to do -- for testing purpose ok
	time.Sleep(2 * time.Second)
	taskList := c.TaskList

	g := graph.InitGraph(len(agents), len(taskList))

	//building graph -- O(agents*tasks) -- as agents dictionary can be unordered if we are not interested in use case of longest, waiting agent -- we can apply binary search to reduce
	// graph building from O(agents*tasks) ~ O(tasks *log(agents))
	for i := 0; i < len(taskList); i++ {
		for j := 0; j < len(agents); j++ {
			// there is an edge if skill matches -- there can be separate method for comparing complex string matching
			if taskList[i].TaskExpression == agents[j].AgentExpression {
				g.AddEdges(j, i)
			}
		}
	}

	// getting maximum cardinality matching = 1 as one agent can take one task -- Time complexity -- O(tasks*sqrt(agents))
	matchingEdges := g.HopcroftKarp()
	for key, val := range matchingEdges {
		fmt.Printf("agent: %s assigned the task: %s", agents[key].AgentId, taskList[val].Tid)
		fmt.Println()
	}
}
