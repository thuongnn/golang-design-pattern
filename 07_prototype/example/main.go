package main

import "fmt"

/*------------------ Init prototype pattern ------------------*/

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

/*------------------ Define object need to clone ------------------*/

type Computer struct {
	os        string
	office    string
	antivirus string
	browser   string
	others    string
}

func (c *Computer) Clone() Cloneable {
	tc := *c
	return &tc
}

/*------------------ Using prototype pattern ------------------*/

var manager *PrototypeManager

func init() {
	manager = NewPrototypeManager()
	manager.Set("computer", &Computer{
		os:        "Window 10",
		office:    "Word 2019",
		antivirus: "BKAV",
		browser:   "Chrome v69",
		others:    "Skype",
	})
}

func main() {
	c1 := manager.Get("computer").Clone()
	forUser1 := c1.(*Computer)
	forUser1.os = "Linux"

	c2 := manager.Get("computer").Clone()
	forUser2 := c2.(*Computer)
	forUser2.os = "Ubuntu"

	fmt.Println(forUser1)
	fmt.Println(forUser2)
}