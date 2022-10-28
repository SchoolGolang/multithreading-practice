package processor

import (
	"context"
	"fmt"

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
		case sensorPlant := <-p.input:
			if sensorPlant.PlantID == "" {
				continue
			}
			fmt.Printf("[Health Processor] got %q\n", sensorPlant.PlantID)
			plant := p.plantsRepo.GetPlant(sensorPlant.PlantID)
			leaves := sensorPlant.Data.LeavesState
			roots := sensorPlant.Data.RootsState
			fmt.Println("[Health Processor] plant: ", plant.Name, plant.ID, "health leaves:", leaves, ", health roots: ", roots)
			if leaves < 50 || roots < 50 {
				a := p.dronesRepo.ReplacePlant(plant.ID)
				fmt.Println(a)
			}
		case <-ctx.Done():
			return
		}
	}
}
