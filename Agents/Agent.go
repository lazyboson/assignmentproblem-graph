package Agents

import (
	"github.com/lazyboson/assignmentproblem/tasks"
	"github.com/segmentio/ksuid"
)

type Agent struct {
	AgentId         string `json:"agent_id"`
	AgentExpression string `json:"agent_expression"`
}

func GenerateAgents() []*Agent {
	agents := make([]*Agent, 0)

	for i := 1; i <= 20; i++ {
		pAgent := &Agent{
			AgentId:         ksuid.New().String(),
			AgentExpression: tasks.GetRandomSkill(),
		}
		agents = append(agents, pAgent)
	}

	return agents
}
