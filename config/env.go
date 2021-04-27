package config

import "os"

var (
	ServerPort         = GetEnv("SERVER_PORT", "9000")
	Job1CronExpression = GetEnv("JOB1_CRON_EXPRESSION", "*/1 * * * *")
	Job2CronExpression = GetEnv("JOB2_CRON_EXPRESSION", "*/2 * * * *")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
