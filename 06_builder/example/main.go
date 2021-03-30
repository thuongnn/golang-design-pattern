package main

import (
	"fmt"
)

func main() {
	newCar := &Car{}
	NewCarBuilder(newCar).Construct("LUXURY", "v12", "KIA SOLUTO")
	fmt.Println(newCar.getResult())
}

type ICar interface {
	addChassis(chassis string)
	addEngine(engine string)
	addBody(body string)
	getResult() string
}

type CarBuilder struct {
	builder ICar
}

func NewCarBuilder(car ICar) *CarBuilder {
	return &CarBuilder{builder: car}
}

func (cb *CarBuilder) Construct(chassis, engine, body string) {
	cb.builder.addChassis(chassis)
	cb.builder.addEngine(engine)
	cb.builder.addBody(body)
}

type Car struct {
	chassis string
	engine  string
	body    string
}

func (c *Car) addChassis(chassis string) {
	c.chassis = chassis
}

func (c *Car) addEngine(engine string) {
	c.engine = engine
}

func (c *Car) addBody(body string) {
	c.body = body
}

func (c *Car) getResult() string {
	return fmt.Sprintf("New car with information:\n"+
		"chassis: %s\n"+
		"engine: %s\n"+
		"body: %s\n", c.chassis, c.engine, c.body)
}
