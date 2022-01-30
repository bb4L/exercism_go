package elon

import "fmt"

// Drive drives forward
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.distance = car.distance + car.speed
	car.battery = car.battery - car.batteryDrain
}

// DisplayDistance displays the driven distance
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the battery level
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish returns true if the car can finish the track
func (car *Car) CanFinish(trackDistance int) bool {
	return trackDistance <= (car.battery/car.batteryDrain)*car.speed
}
