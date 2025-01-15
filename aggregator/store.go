package main

import "traffic-toll-calculator/types"

type Storer interface {
}

type MemoryStore struct {
	data types.MapStorage
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}
