package pkg

import (
	"github.com/robfig/cron/v3"
)

func NewCron() *cron.Cron {
	return cron.New(cron.WithSeconds())
}
