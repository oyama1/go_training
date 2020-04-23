package conv

import "fmt"

// 型の宣言
type Kg float64
type Lb float64 

func (kg Kg) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func (lb Lb) String() string {
	return fmt.Sprintf("%glb", lb)
}