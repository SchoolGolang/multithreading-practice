package processor

import (
	"context"
	"fmt"

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
		case sensorPlant := <-p.input:
			if sensorPlant.PlantID == "" {
				continue
			}
			plant := p.plantsRepo.GetPlant(sensorPlant.PlantID)
			hd := sensorPlant.Data
			hdNormal := plant.NormalHydration
			fmt.Println("[Hydration Processor] plant: ", plant.Name, plant.ID, "current hydration:", hd, "normal: ", hdNormal)
			if hd < hdNormal {
				p.dronesRepo.Hydrate(plant.ID, hdNormal)
			}
		case <-ctx.Done():
			return
		}
	}
}
