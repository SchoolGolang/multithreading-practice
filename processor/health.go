package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
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
	//TODO: implement process functionality

	for {
		select {
		case data := <-p.input:
			log.Printf("type: %T, value: %[1]v", data)
			negative := p.plantsRepo.GetPlant(data.PlantID)
			if negative.CurrentHealth.LeavesState < 50 || negative.CurrentHealth.RootsState < 50 {
				p.dronesRepo.ReplacePlant(data.PlantID)
			}
		case <-ctx.Done():
			return
		}
	}
}
