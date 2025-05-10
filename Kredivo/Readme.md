1. Registration numbers of all cars of a particular colour.
2. Slot number in which a car with a given registration number is parked.
3. Slot numbers of all slots where a car of a particular colour is parked.

- maxCapacity => n

```Go
type Vehicle struct{
    Colour string
    RegistrationNo string
}

type Slot struct{
    SlotNo      int
    IsAvaliable bool
}

type ParkinglotSystem struct{
    MaxCapacity int
    SlotVehicleMap map[int]Vehicle
}

```