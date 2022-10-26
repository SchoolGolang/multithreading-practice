package repository

import (
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"sync"
)

type SensorRepo[T any] struct {
	sensors map[string]sensor.Sensor[T]
	mu      *sync.RWMutex
}

func NewRepository[T any]() *SensorRepo[T] {
	return &SensorRepo[T]{
		sensors: make(map[string]sensor.Sensor[T]),
		mu:      new(sync.RWMutex),
	}
}

func (r *SensorRepo[T]) AddSensor(sensor *sensor.Sensor[T]) {
	r.mu.RLock()
	_, ok := r.sensors[sensor.ID];
	r.mu.RUnlock()

	if ok {
		return
	}

	r.mu.Lock()
	r.sensors[sensor.ID] = *sensor
	r.mu.Unlock()
}

func (r *SensorRepo[T]) GetSensor(id string) sensor.Sensor[T] {
	r.mu.RLock()
	sensr := r.sensors[id]
	r.mu.RUnlock()

	return sensr
}

func (r *SensorRepo[T]) GetSensorByPlantID(plantID string) *sensor.Sensor[T] {
	r.mu.RLock()
	for _, v := range r.sensors {
		if v.PlantID == plantID {
			return &v
		}
	}
	r.mu.RUnlock()

	return nil
}

func (r *SensorRepo[T]) RemoveSensorByPlantID(plantID string) {
	r.mu.Lock()
	for key := range r.sensors {
		s := r.sensors[key]
		if s.PlantID == plantID {
			s.Disconnect()
			delete(r.sensors, key)
		}
	}
	r.mu.Unlock()
}

func (r *SensorRepo[T]) GetAll() []sensor.Sensor[T] {
	r.mu.RLock()
	sensors := make([]sensor.Sensor[T], 0, len(r.sensors))

	for _, item := range r.sensors {
		sensors = append(sensors, item)
	}
	r.mu.RUnlock()

	return sensors
}
