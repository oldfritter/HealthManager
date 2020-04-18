package publishes

import (
	"github.com/robfig/cron/v3"

	"github.com/oldfritter/HealthManager/publishes/tasks"
)

func InitSchedule() {
	c := cron.New(cron.WithSeconds())

	// 每10秒广播一次
	c.AddFunc("*/10 * * * * *", tasks.ComponentList)
	c.AddFunc("*/10 * * * * *", tasks.ServiceList)

	c.Start()
}
