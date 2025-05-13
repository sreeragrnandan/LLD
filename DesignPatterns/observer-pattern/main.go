package main

import "fmt"

// Observer interface - any observer must implement Update()
type Observer interface {
	Update(temperature float64)
}

// Subject interface - methods for registering, deregistering, and notifying observers
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

// WeatherStation Concrete Subject
type WeatherStation struct {
	observers   []Observer // List of registered observers
	temperature float64    // Current temperature state
}

// Register adds a new observer to the list
func (ws *WeatherStation) Register(o Observer) {
	ws.observers = append(ws.observers, o)
}

// Deregister removes an observer from the list
func (ws *WeatherStation) Deregister(o Observer) {
	for i, observer := range ws.observers {
		if observer == o {
			// Remove observer by slicing
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll sends the current temperature to all observers
func (ws *WeatherStation) NotifyAll() {
	for _, observer := range ws.observers {
		observer.Update(ws.temperature)
	}
}

// SetTemperature updates the temperature and notifies observers
func (ws *WeatherStation) SetTemperature(temp float64) {
	fmt.Printf("WeatherStation: new temperature is %.2f°C\n", temp)
	ws.temperature = temp
	ws.NotifyAll()
}

// PhoneDisplay Concrete Observer
type PhoneDisplay struct {
	name string
}

// Update is called when the temperature changes
func (p *PhoneDisplay) Update(temp float64) {
	fmt.Printf("%s: received temperature update: %.2f°C\n", p.name, temp)
}

// WindowDisplay Concrete Observer
type WindowDisplay struct{}

// Update is called when the temperature changes
func (w *WindowDisplay) Update(temp float64) {
	fmt.Printf("WindowDisplay: current temperature is %.2f°C\n", temp)
}

// Main - simulate the observer pattern in action
func main() {
	// Create subject (WeatherStation)
	station := &WeatherStation{}

	// Create observers
	phone1 := &PhoneDisplay{name: "Phone 1"}
	phone2 := &PhoneDisplay{name: "Phone 2"}
	window := &WindowDisplay{}

	// Register observers to the station
	station.Register(phone1)
	station.Register(phone2)
	station.Register(window)

	// Change temperature - all observers will be notified
	station.SetTemperature(25.0)
	station.SetTemperature(30.5)

	// Deregister one observer and change temperature again
	station.Deregister(phone2)
	station.SetTemperature(27.3)
}
