package main
import (
	"fmt"
)

type IProcess interface {        //IProces interface
	process()
}
type Adapter struct {          //Adapter struct
	adaptee Adaptee
}
func (adapter Adapter) process() {     //Adapter class method process
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct { //Adaptee Struct
	adapterType int
}

func (adaptee Adaptee) convert() {                     // Adaptee class method convert
	fmt.Println("Adaptee convert method")
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
