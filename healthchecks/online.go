package healthchecks

import (
	"fmt"
	"time"

	"github.com/strivesolutions/go-health-checks/pkg/health"
)

func OnlineCheck() health.Checker {
	return health.CreateHealthCheck("online", func(name string, out chan health.CheckResult) {
		out <- health.Ok(name, []string{
			fmt.Sprintf("Server time: %s", time.Now()),
		})
	})
}
