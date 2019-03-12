package workers

import (
    "testing"
    "time"
)

// deal deploy
func Test_DeployDealOnce(t *testing.T) {
    DefaultDeployWorker.DealOnce()

    time.Sleep(time.Second * 600)
}
