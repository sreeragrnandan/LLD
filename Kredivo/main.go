package main

import (
	"errors"
	"fmt"
	"time"
)

type Vehicle struct {
	Colour         string
	RegistrationNo string
}

func NewVehicle(color, registrationNo string) *Vehicle {
	return &Vehicle{
		Colour:         color,
		RegistrationNo: registrationNo,
	}
}

type Slot struct {
	SlotNo      int
	IsAvaliable bool
}

type ParkingLotSytem struct {
	MaxCapacity    int
	SlotVehicleMap map[int64]Vehicle
}

func NewParkingLotSytem(maxCapacity int) *ParkingLotSytem {
	return &ParkingLotSytem{
		MaxCapacity:    maxCapacity,
		SlotVehicleMap: make(map[int64]Vehicle, maxCapacity),
	}

}

type ParkingLotSystem interface {
	ParkVehicles(vehicle Vehicle) error
	RemoveVechicles(registrationNo string) error

	GetReistrationByColor(color string) (registrationNo []string)
	GetSlotByReistration(color string) (slotNo []string)

	GetSlotByColor(color string) (slotNo []string)
}

func (p *ParkingLotSytem) ParkVehicles(vehicle Vehicle) error {

	if len(p.SlotVehicleMap) > p.MaxCapacity {
		fmt.Println("Parking lot is full")
		return errors.New("Parking lot is full")
	}

	uuid := time.Now().UnixMicro()
	p.SlotVehicleMap[uuid] = vehicle

	fmt.Printf("parking vehicle %+v \n", p.SlotVehicleMap)

	return nil
}

func (p *ParkingLotSytem) RemoveVechicles(registrationNo string) error {
	for slotNum, v := range p.SlotVehicleMap {
		if v.RegistrationNo == registrationNo {
			delete(p.SlotVehicleMap, slotNum)
		}
	}

	return nil
}

func (p *ParkingLotSytem) GetReistrationByColor(color string) (registrationNo []string) {
	for _, v := range p.SlotVehicleMap {
		if v.Colour == color {
			registrationNo = append(registrationNo, v.RegistrationNo)
			fmt.Println(registrationNo)
		}
	}

	return
}

func main() {
	vehicle1 := NewVehicle("black", "KA-1235")
	vehicle2 := NewVehicle("blue", "KL-1235")
	// vehicle3 := NewVehicle("blue", "KL-1235")
	p := NewParkingLotSytem(2)
	p.ParkVehicles(*vehicle1)
	p.ParkVehicles(*vehicle2)
	p.GetReistrationByColor("black")

	// p.ParkVehicles(*vehicle3)
}
