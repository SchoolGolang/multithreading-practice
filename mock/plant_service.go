package mock

import (
	"context"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	sensorRepo "github.com/SchoolGolang/multithreading-practice/sensor/repository"
	"github.com/SchoolGolang/multithreading-practice/util"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

type PlantsServiceMock struct {
	plantsRepo    *repository.PlantsRepository
	phRepo        *sensorRepo.SensorRepo[int]
	hydrationRepo *sensorRepo.SensorRepo[float64]
	healthRepo    *sensorRepo.SensorRepo[plant.HealthData]

	frequency int
}

func NewPlantsServiceMock(
	plantsRepo *repository.PlantsRepository,
	phRepo *sensorRepo.SensorRepo[int],
	hydrationRepo *sensorRepo.SensorRepo[float64],
	healthRepo *sensorRepo.SensorRepo[plant.HealthData],
	frequency int,
) *PlantsServiceMock {
	return &PlantsServiceMock{
		plantsRepo:    plantsRepo,
		phRepo:        phRepo,
		hydrationRepo: hydrationRepo,
		healthRepo:    healthRepo,
		frequency:     frequency,
	}
}

func (ps *PlantsServiceMock) SendRandomUpdates(ctx context.Context) {
	for {
		plantIDs := ps.plantsRepo.GetPlantIds()
		plantID := plantIDs[util.GetRandomIndex(len(plantIDs))]
		select {
		case <-ctx.Done():
			return
		default:
			switch rand.Intn(20) {
			case 0:
				pID, ph := plantID, GetPHData()
				log.Printf("Plant ID: %v, nes ph: %d", pID, ph)
				ps.UpdatePlantPH(pID, ph)
			case 1:
				pID, gd := plantID, GetHydrationData()
				log.Printf("Plant ID: %v, nes gd: %d", pID, gd)
				ps.UpdatePlantHydration(pID, gd)

			case 2:
				pID, hp := plantID, GetHealthData()
				log.Printf("Plant ID: %v, nes hp: %d", pID, hp)
				ps.UpdatePlantHealth(pID, hp)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func (ps *PlantsServiceMock) AddPlant() string {
	plantData := GetPlantData()
	plantId := plantData.ID
	ps.plantsRepo.AddPlant(plantData)

	phSensor := sensor.NewSensor[int](uuid.New().String(), plantId)
	ps.phRepo.AddSensor(phSensor)

	hydrationSensor := sensor.NewSensor[float64](uuid.New().String(), plantId)
	ps.hydrationRepo.AddSensor(hydrationSensor)

	healthSensor := sensor.NewSensor[plant.HealthData](uuid.New().String(), plantId)
	ps.healthRepo.AddSensor(healthSensor)

	return plantId
}

func (ps *PlantsServiceMock) RemovePlant(plantID string) {
	ps.plantsRepo.RemovePlant(plantID)

	ps.phRepo.RemoveSensorByPlantID(plantID)
	ps.hydrationRepo.RemoveSensorByPlantID(plantID)
	ps.healthRepo.RemoveSensorByPlantID(plantID)
}

func (ps *PlantsServiceMock) UpdatePlantPH(plantId string, ph int) {
	s := ps.phRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[int]{
		PlantID: s.PlantID,
		Data:    ph,
	}
}

func (ps *PlantsServiceMock) UpdatePlantHydration(plantId string, hydration float64) {
	s := ps.hydrationRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[float64]{
		PlantID: s.PlantID,
		Data:    hydration,
	}
}

func (ps *PlantsServiceMock) UpdatePlantHealth(plantId string, health plant.HealthData) {
	s := ps.healthRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[plant.HealthData]{
		PlantID: s.PlantID,
		Data:    health,
	}
}
