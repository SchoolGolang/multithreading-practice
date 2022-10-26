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
	//TODO: implement process functionality
	for{
	select {
	case RecData := <-p.input:

		plant := p.plantsRepo.GetPlant(RecData.PlantID)
		switch  {
		case  int(RecData.Data) < plant.NormalLowerPh || int(RecData.Data) > plant.NormalUpperPh:
			p.dronesRepo.AdjustSoils(RecData.PlantID, (plant.NormalUpperPh+plant.NormalLowerPh)/2)
			log.Printf("У рослини %s з ID: %s: Стан кислотності: %v, при нормі: %v-%v  Встановити: %v  ",plant.Name,RecData.PlantID, RecData.Data, plant.NormalLowerPh, plant.NormalUpperPh ,(plant.NormalUpperPh+plant.NormalLowerPh)/2)
		default:
			log.Printf("У рослини %s з ID: %s: Стан кислотності: %v, при нормі %v-%v",plant.Name,RecData.PlantID,RecData.Data,plant.NormalLowerPh,plant.NormalUpperPh)
		}
	case <- ctx.Done():
		return
	}
	}
}
