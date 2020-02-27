package main

import "fmt"

type USB interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

// implement USB start
func (p *Phone) start() {
	fmt.Printf("%s is starting.\n", p.Name)
}

// implement USB stop
func (p *Phone) stop() {
	fmt.Printf("%s is stop.\n", p.Name)
}

type Camera struct {
	Name string
}

// implement USB start
func (c Camera) start() {
	fmt.Printf("%s is starting.\n", c.Name)
}

// implement USB stop
func (c Camera) stop() {
	fmt.Printf("%s is stop.\n", c.Name)
}

// work with USB
type Computer struct{}

func (c Computer) Work(usb USB) {
	if p, ok := usb.(*Phone); ok {
		p.Name = "newPhone"
	}
	usb.start()
	usb.stop()
}

func main() {
	phone := &Phone{Name: "myphone"}
	camera := Camera{Name: "newcamera"}
	computer := Computer{}

	computer.Work(phone)
	computer.Work(camera)

	fmt.Println(phone)

}

// Test01 duplication
// type IA interface{
// 	Test01()
// 	Test02()
// }

// type IB interface{
// 	Test01()
// 	Test03()
// }

// type C interface{
// 	IA
// 	IB
// }
