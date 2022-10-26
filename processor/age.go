package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
)

type AgeProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[int]
	dronesRepo droneRepository.DroneRepo
}

func NewAgeProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[int],
	dronesRepo droneRepository.DroneRepo,
) *AgeProcessor {
	return &AgeProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *AgeProcessor) RunProcessor(ctx context.Context) {
	//TODO: implement process functionality

	for{
		select {
		case RecDate:=<-p.input:
			plant := p.plantsRepo.GetPlant(RecDate.PlantID)
			switch {
			case int(RecDate.Data) == 30:
				log.Printf("Рослина %s, з ID: %s: дозріла!!!",plant.Name, RecDate.PlantID)
			default :
				leftDay:=30-int(RecDate.Data)// 30 днів термін достигання
				log.Printf("Вік рослини %s, з ID: %s: %v, рослина дозріє через %v днів",plant.Name, RecDate.PlantID,RecDate.Data,leftDay)
			}
		case <-ctx.Done():
			return
		}
	}
}