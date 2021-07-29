package conta

import (
	"fmt"
	"math"
)

func Sum() {
	const PI = 3.1415
	raio := 3.2
	area := PI * math.Pow(raio, 2)
	fmt.Println("O tamanho da área é", area)
}
