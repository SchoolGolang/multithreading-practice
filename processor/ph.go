package processor

import (
	"context"
	"fmt"

	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
)

type PHProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[int]
	dronesRepo droneRepository.DroneRepo
}

func NewPHProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[int],
	dronesRepo droneRepository.DroneRepo,
) *PHProcessor {
	return &PHProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *PHProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case sensorPlant := <-p.input:
			if sensorPlant.PlantID == "" {
				continue
			}
			plant := p.plantsRepo.GetPlant(sensorPlant.PlantID)
			ph := sensorPlant.Data
			normalLowerPh := plant.NormalLowerPh
			normalUpperPh := plant.NormalUpperPh
			normalPh := (normalLowerPh + normalUpperPh) / 2
			fmt.Println("[PHP Processor] plant: ", plant.Name, plant.ID, "current ph:", ph, "normal [", normalLowerPh, ", ", normalUpperPh, "]")
			if ph < normalLowerPh || normalUpperPh < ph {
				p.dronesRepo.AdjustSoils(plant.ID, normalPh)
			}
		case <-ctx.Done():
			return
		}
	}
}
