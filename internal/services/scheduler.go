package services

import (
	"github.com/madflojo/tasks"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Scheduler = "scheduler"

var SchedulerDef = dingo.Def{
	Name:  Scheduler,
	Scope: di.App,
	Build: func() (*tasks.Scheduler, error) {
		return tasks.New(), nil
	},
	Close: func(s *tasks.Scheduler) error {
		s.Stop()
		return nil
	},
}
