package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type HydrationProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[float64]
	dronesRepo droneRepository.DroneRepo
}

func NewHydrationProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[float64],
	dronesRepo droneRepository.DroneRepo,
) *HydrationProcessor {
	return &HydrationProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HydrationProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case measurement := <-p.input:
			go func() {
				plantID := measurement.PlantID
				currHydration := measurement.Data
				normalHydration := p.plantsRepo.GetNormalHydration(plantID)

				if currHydration < normalHydration {
					p.dronesRepo.Hydrate(plantID, 1.0-currHydration)
					// p.dronesRepo.Hydrate(plantID, normalHydration-currHydration) Немає інформації як поливати.
				} // За умовою невідомо чи треба строга нерівність чи ні.
			}()
		}
	}
}
