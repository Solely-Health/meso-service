package facilities

import (
	"github.com/asaskevich/EventBus"
	"github.com/meso-org/meso/workers"
)

func initializeListener(eventBus EventBus.Bus) {
	go func() {
		eventBus.Subscribe("workers:NewWorkerRegistered", workers.NewWorkerRegistered)
	}()
}
