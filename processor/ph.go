package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
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
		case data := <-p.input:
			log.Printf("type ph: %T, value: %[1]v", data)
			low, high := p.plantsRepo.GetNormalPh(data.PlantID)
			if low > int(data.Data) || int(data.Data) > high {
				p.dronesRepo.AdjustSoils(data.PlantID, (low+high)/2)

			}

		case <-ctx.Done():
			return
		}
	}
}
