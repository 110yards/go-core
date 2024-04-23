package postgres

import (
	"fmt"

	"github.com/strivesolutions/go-health-checks/pkg/health"
)

func PostgresCheck() health.Checker {
	return health.CreateHealthCheck("postgres", func(name string, out chan health.CheckResult) {

		err := GetDb().Ping()

		if err != nil {
			out <- health.Unhealthy(name, fmt.Sprintf("Error connecting to database: %s", err), nil)
		} else {
			out <- health.Ok(name, nil)
		}
	})
}
