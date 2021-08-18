package main

import "fmt"

const uxisteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type Car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c Car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / uxisteenbitmax)
}

func (c Car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / uxisteenbitmax / kmh_multiple)
}

func (c *Car) new_top_speed(speed float64) {
	c.top_speed_kmh = speed
}

func another_new_top_speed(c Car, speed float64) Car {
	c.top_speed_kmh = speed

	return c
}

func structs() {
	a_car := Car{
		gas_pedal:      65000,
		brake_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0,
	}

	fmt.Printf("Gas Pedal: %d", a_car.gas_pedal)
	fmt.Println()
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	// a_car.new_top_speed(500)
	a_car = another_new_top_speed(a_car, 500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
