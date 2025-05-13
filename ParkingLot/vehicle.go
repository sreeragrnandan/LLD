package parkinglot

type ParkingSpot struct {
	spotNumber    int
	vehicleType   VehicleType
	parkedVehicle Vehicle
}

func NewParkingSpot(spotNumber int, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{spotNumber: spotNumber, vehicleType: vehicleType}
}

type ParkingLot struct {
	levels []*Level
}

type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

type BaseVehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (v *BaseVehicle) GetLicensePlate() string {
	return v.licensePlate
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}

func NewCar(licensePlate string) Vehicle {
	return &BaseVehicle{licensePlate: licensePlate, vehicleType: CAR}
}

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor int, numSpots int) *Level {
	level := &Level{floor: floor}
	bikeSpots := int(float64(numSpots) * 0.50)
	carSpots := int(float64(numSpots) * 0.40)

	for i := 1; i <= bikeSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, MOTORCYCLE))
	}
	for i := bikeSpots + 1; i <= bikeSpots+carSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, CAR))
	}
	for i := bikeSpots + carSpots + 1; i <= numSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, TRUCK))
	}
	return level
}
