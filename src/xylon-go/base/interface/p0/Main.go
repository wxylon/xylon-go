package main

/*import (
	"fmt"
)*/

type IHello interface {
	hi(str string) error
}

func main() {
	iHelloes := map[string]IHello{}
	iHelloes["pig"] = new(HelloPig)
	iHelloes["person"] = new(HelloPerson)

	_ = iHelloes["person"].hi("ss")
	_ = iHelloes["pig"].hi("ss")
}
