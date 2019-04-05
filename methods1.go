package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 // 0 to 65535
	brake_pedal    uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh  float64
}

func (c *car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{
		gas_pedal:      22341,
		brake_pedal:    0,
		steering_wheel: 1234,
		top_speed_kmh:  225.0}
	//a_car := car{223,0,123,335.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
