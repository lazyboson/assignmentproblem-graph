package tasks

import (
	"github.com/segmentio/ksuid"
	"math/rand"
)

func GetRandomSkill() string {
	skills := [10]string{
		"sales",
		"support",
		"l3",
		"hindi",
		"english",
		"french",
		"spanish",
		"technical",
		"billing",
		"default",
	}

	return skills[rand.Intn(9-0)-0]
}

type Task struct {
	Tid            string `json:"tid"`
	TaskExpression string `json:"task_expression"`
}

func GenerateTask() *Task {
	t := &Task{
		Tid:            ksuid.New().String(),
		TaskExpression: GetRandomSkill(),
	}

	return t
}
