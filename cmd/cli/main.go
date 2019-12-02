package main

import (
	"fmt"
	"github.com/icamonkey/VolumePUI/pkg/volume"
)

func main() {

	i, e := volume.GetVolume()
	if e != nil{
		fmt.Println("Volume Error:", e)
	}

	fmt.Println(i)
}
