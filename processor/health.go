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
	for {
		select {
		case RecData := <-p.input:
			plant := p.plantsRepo.GetPlant(RecData.PlantID)
			switch {
			case plant.CurrentHealth.LeavesState < 50 || plant.CurrentHealth.RootsState < 50:
				p.dronesRepo.ReplacePlant(RecData.PlantID)
				log.Printf("Рівень здоров'я рослини %s з ID: %s <50  потребує заміни", plant.Name, RecData.PlantID)
			default:
				log.Printf("У рослини %s з ID: %s рівень здоров'я листя %v у та коріння: %v, ", plant.Name, RecData.PlantID, plant.CurrentHealth.LeavesState, plant.CurrentHealth.RootsState)
			}
		case <-ctx.Done():
			return
		}
	}
}
