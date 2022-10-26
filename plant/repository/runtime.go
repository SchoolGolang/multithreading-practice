package repository

import (
	"github.com/SchoolGolang/multithreading-practice/plant"
	"sync"
)

type PlantsRepository struct {
	plantsData map[string]plant.PlantData
	mu         *sync.RWMutex
}

func NewRepository() *PlantsRepository {
	return &PlantsRepository{
		plantsData: make(map[string]plant.PlantData),
		mu:         new(sync.RWMutex),
	}
}

func (pr *PlantsRepository) AddPlant(data plant.PlantData) {
	pr.mu.RLock()
	_, ok := pr.plantsData[data.ID]
	pr.mu.RUnlock()

	if ok {
		return
	}

	pr.mu.Lock()
	pr.plantsData[data.ID] = data
	pr.mu.Unlock()
}

func (pr *PlantsRepository) RemovePlant(id string) {
	pr.mu.Lock()
	delete(pr.plantsData, id)
	pr.mu.Unlock()
}

func (pr *PlantsRepository) GetPlant(id string) plant.PlantData {
	pr.mu.RLock()
	plnt := pr.plantsData[id]
	pr.mu.RUnlock()

	return plnt
}

func (pr *PlantsRepository) GetHydration(id string) float64 {
	pr.mu.RLock()
	plntHdrtn := pr.plantsData[id].CurrentHydration
	pr.mu.RUnlock()

	return plntHdrtn
}

func (pr *PlantsRepository) GetPh(id string) int {
	pr.mu.RLock()
	plntPh := pr.plantsData[id].CurrentPh
	pr.mu.RUnlock()

	return plntPh
}

func (pr *PlantsRepository) GetHealth(id string) plant.HealthData {
	pr.mu.RLock()
	plntHealth := pr.plantsData[id].CurrentHealth
	pr.mu.RUnlock()

	return plntHealth
}

func (pr *PlantsRepository) GetNormalHydration(id string) float64 {
	pr.mu.RLock()
	plntNormalHdrtn := pr.plantsData[id].NormalHydration
	pr.mu.RUnlock()

	return plntNormalHdrtn
}

func (pr *PlantsRepository) GetNormalPh(id string) (int, int) {
	pr.mu.RLock()
	plnt := pr.plantsData[id]
	pr.mu.RUnlock()

	return plnt.NormalLowerPh, plnt.NormalUpperPh
}

func (pr *PlantsRepository) GetPlantIds() []string {
	pr.mu.RLock()
	ids := make([]string, 0, len(pr.plantsData))
	for k := range pr.plantsData {
		ids = append(ids, k)
	}

	pr.mu.RUnlock()

	return ids
}

func (pr *PlantsRepository) SetPh(id string, ph int) {
	pr.mu.Lock()
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentPh = ph
	}
	pr.mu.Unlock()
}

func (pr *PlantsRepository) SetHydration(id string, hydration float64) {
	pr.mu.Lock()
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentHydration = hydration
	}
	pr.mu.Unlock()
}

func (pr *PlantsRepository) SetHealth(id string, data plant.HealthData) {
	pr.mu.Lock()
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentHealth = data
	}
	pr.mu.Unlock()
}
