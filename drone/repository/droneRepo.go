package repository

import (
	"fmt"

	"github.com/SchoolGolang/multithreading-practice/drone"
	"github.com/SchoolGolang/multithreading-practice/mock"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
)

type DroneRepo struct {
	plantsService *mock.PlantsServiceMock
	plantsRepo    repository.Repository

	drones map[string]drone.Drone
}

func NewDroneRepo(plantsService *mock.PlantsServiceMock, plantsRepository repository.Repository) DroneRepo {
	return DroneRepo{
		plantsService: plantsService,
		plantsRepo:    plantsRepository,
		drones:        make(map[string]drone.Drone),
	}
}

func (d *DroneRepo) AdjustSoils(plantId string, ph int) {
	fmt.Println("AdjustSoils")
	d.plantsService.UpdatePlantPH(plantId, ph)
}

func (d *DroneRepo) Hydrate(plantId string, value float64) {
	fmt.Println("Hydrate")
	d.plantsService.UpdatePlantHydration(plantId, value)
}

func (d *DroneRepo) ReplacePlant(plantId string) string {
	fmt.Println("ReplacePlant")
	d.plantsService.RemovePlant(plantId)
	return d.plantsService.AddPlant()
}
