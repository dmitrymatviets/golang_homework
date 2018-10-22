package model

import "fmt"

type Budka struct {
	CheckIner CheckIner
}

func (b *Budka) CheckIn(passanger Passenger) (*Pass, error) {
	if b.CheckIner == nil {
		return nil, fmt.Errorf("under construction")
	}
	return b.CheckIner.CheckIn(passanger)
}

func (b *Budka) Register(c CheckIner) {
	b.CheckIner = c
}
