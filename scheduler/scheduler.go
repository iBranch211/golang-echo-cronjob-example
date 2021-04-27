package scheduler

import (
	"github.com/robfig/cron/v3"
	"golang-echo-cronjob-example/config"
	"golang-echo-cronjob-example/task"
)

func CreateScheduledJobs(c *cron.Cron) {
	c.AddFunc(config.Job1CronExpression, task.RunJob1ScheduledTask)
	c.AddFunc(config.Job2CronExpression, task.RunJob2ScheduledTask)

}
