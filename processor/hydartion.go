package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
	//"log"
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
	//TODO: implement process functionality
	for {
		select {
		case data := <-p.input:
			log.Printf("type: %T, value: %[1]v", data)
			normal := p.plantsRepo.GetHydration(data.PlantID)
			if normal > float64(data.Data) {
				p.dronesRepo.Hydrate(data.PlantID, normal)
			}
		case <-ctx.Done():
			return
		}
	}
}
