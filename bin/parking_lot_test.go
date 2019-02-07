package main

import (
	"log"
	"testing"
)

func TestParkingLot(t *testing.T) {
	input := [][]string{
		{"create_parking_lot,6"},
		{"park KA-01-HH-1234,White"},
		{"park KA-01-HH-9999 White"},
	}
	actionParking := &ResParkingLot{}

	t.Run("input type", func(t *testing.T) {
		for _, v := range input {
			InputType(actionParking, v)
		}

	})

	t.Run("CreateParkingLot", func(t *testing.T) {
		lot := []int{1, 2, 3}
		for i := range lot {
			res := actionParking.CreateParkingLot(lot[i])
			log.Println(res)
		}
	})

	t.Run("Allocated", func(t *testing.T) {
		regis := []string{"KA-01-HH-1234", "KA-01-HH-12344", "KA-01-HH-12341"}
		color := []string{"white", "red", "blue"}
		for i := range regis {
			res, _ := actionParking.Allocated(regis[i], color[i])
			log.Println(res)
		}

	})
	t.Run("Leave", func(t *testing.T) {
		lot := []int{1, 2, 3}
		for i := range lot {
			res, _ := actionParking.Leave(lot[i])
			log.Println(res)
		}
	})
	t.Run("status", func(t *testing.T) {
		actionParking.Status()
	})

	t.Run("FindByColor", func(t *testing.T) {
		color := []string{"white", "red", "blue"}
		for i := range color {
			res, _ := actionParking.FindByColor(color[i])
			log.Println(res)
		}
	})

	t.Run("FindByRegis", func(t *testing.T) {
		regis := []string{"KA-01-HH-1234", "KA-01-HH-12344", "KA-01-HH-12341"}
		for i := range regis {
			res, _ := actionParking.FindByRegis(regis[i])
			log.Println(res)
		}
	})

}
