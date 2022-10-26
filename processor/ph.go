package processor

import (
	"context"
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
		case <-ctx.Done():
			return
		case measurement := <-p.input:
			go func() {
				plantID := measurement.PlantID
				currPh := measurement.Data
				minPh, maxPh := p.plantsRepo.GetNormalPh(plantID)
				desiredPh := (minPh + maxPh) / 2 // Щоб не на граничне змінювати, а дати простір для змін в обидві сторони.

				if currPh < minPh {
					p.dronesRepo.AdjustSoils(plantID, desiredPh-currPh)
				} else if maxPh > currPh {
					p.dronesRepo.AdjustSoils(plantID, desiredPh-currPh)
				}
			}()
		}
	}
}
