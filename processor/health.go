package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type HealthProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[plant.HealthData]
	dronesRepo droneRepository.DroneRepo
}

func NewHealthProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[plant.HealthData],
	dronesRepo droneRepository.DroneRepo,
) *HealthProcessor {
	return &HealthProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HealthProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case measurement := <-p.input:
			go func() {
				minHealth := 50.0
				plantID := measurement.PlantID
				currRootsHealth := measurement.Data.RootsState
				currLeavesHealth := measurement.Data.LeavesState

				if currRootsHealth < minHealth || currLeavesHealth < minHealth {
					p.dronesRepo.ReplacePlant(plantID)
				}
			}()
		}
	}
}
