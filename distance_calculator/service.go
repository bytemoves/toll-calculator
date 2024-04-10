package main

import (
	"fmt"

	"github.com/bytemoves/toll-calculator/types"
)

type CalculatorServicer interface{
	calculateDistance(types.OBUData) (float64,error)
}
type CalculatorService struct {

}
func NewCalculatorService() CalculatorServicer{
	return &CalculatorService{}
}

func (s *CalculatorService) calculateDistance(data types.OBUData) (float64,error){
	fmt.Println("calculating the distance")
	return 0.0 ,  nil
}