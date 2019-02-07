package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var path = "./functional_spec/fixtures/file_input.txt"

type ParkingLot struct {
	LotNumber          int
	Color, RegisNumber string
}

type ResParkingLot struct {
	Data []map[string]string
}

func (pl *ResParkingLot) CreateParkingLot(lot int) *ResParkingLot {

	rpl := ResParkingLot{}
	for i := 1; i <= lot; i++ {
		e := strconv.Itoa(i)
		rpl.Data = append(rpl.Data, map[string]string{"lot_number": e, "regis_number": "", "color": ""})
	}
	return &rpl
}

func (pl *ResParkingLot) Allocated(regis, colors string) (*ResParkingLot, string) {

	var res string
	for _, v := range pl.Data {
		if v["regis_number"] == "" && v["color"] == "" {
			v["regis_number"] = regis
			v["color"] = colors
			res = fmt.Sprintf("Allocated slot number: %v", v["lot_number"])
			break
		} else {
			res = fmt.Sprintf("Sorry, parking lot is full")
		}
	}

	return pl, res
}

func (pl *ResParkingLot) Leave(lot int) (*ResParkingLot, string) {
	var res string
	num := strconv.Itoa(lot)
	for _, v := range pl.Data {
		if v["lot_number"] == num {
			v["regis_number"], v["color"] = "", ""
		}
	}
	res = fmt.Sprintf("Slot number %v is free", lot)
	return pl, res
}

func (pl *ResParkingLot) Status() *ResParkingLot {
	log.Println("Slot No.				Registration No.			Colour         ")
	for _, v := range pl.Data {
		a := fmt.Sprintf("%v					%v				%v", v["lot_number"], v["regis_number"], v["color"])
		log.Println(a)
	}
	return pl
}

func (pl *ResParkingLot) FindByColor(color string) (*ResParkingLot, string) {
	var res string
	var list []string
	for _, v := range pl.Data {
		for _, val := range v {
			if val == color {
				list = append(list, v["regis_number"])

			}
		}

	}

	if len(list) == 0 {
		res = fmt.Sprintf("Not Found")
	} else {
		res = fmt.Sprintf(strings.Join(list, ","))
	}
	return pl, res
}

func (pl *ResParkingLot) FindByRegis(regis string) (*ResParkingLot, string) {

	var res string
	var list []string
	for _, v := range pl.Data {
		for _, val := range v {
			if val == regis {
				list = append(list, v["lot_number"])
			}
		}
	}
	if len(list) == 0 {
		res = fmt.Sprintf("Not Found")
	} else {
		res = fmt.Sprintf(strings.Join(list, ","))
	}
	return pl, res
}

func (pl *ResParkingLot) FindByLot(color string) (*ResParkingLot, string) {

	var res string
	var list []string
	for _, v := range pl.Data {
		for _, val := range v {
			if val == color {
				list = append(list, v["lot_number"])
			}
		}
	}
	if len(list) == 0 {
		res = fmt.Sprintf("Not Found")
	} else {
		res = fmt.Sprintf(strings.Join(list, ","))
	}
	return pl, res
}

func main() {
	actionParking := &ResParkingLot{}
	if len(os.Args) == 2 {
		var input []string
		input = FileInputParam(path)

		for _, v := range input {
			strs := strings.Split(v, " ")
			actionParking = InputType(actionParking, strs)
		}
	} else {
		Reader(actionParking)
	}

}

func Reader(actionParking *ResParkingLot) {

	for {
		reader := bufio.NewReader(os.Stdin)
		log.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		strs := strings.Split(text, " ")

		actionParking = InputType(actionParking, strs)
	}
}

func InputType(actionParking *ResParkingLot, input []string) *ResParkingLot {
	var msg string
	switch input[0] {
	case "create_parking_lot":
		lot, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal("convert failed", err)
		}

		actionParking = actionParking.CreateParkingLot(lot)
		msg = fmt.Sprintf("Created a parking lot with %v slots", len(actionParking.Data))
		log.Println(msg)
	case "park":
		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		actionParking, msg = actionParking.Allocated(input[1], input[2])
		log.Println(msg)
	case "leave":

		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		lot, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal("convert failed", err)
		}
		actionParking, msg = actionParking.Leave(lot)
		log.Println(msg)
	case "status":
		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		actionParking = actionParking.Status()
	case "registration_numbers_for_cars_with_colour":
		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		actionParking, msg = actionParking.FindByColor(input[1])
		log.Println(msg)
	case "slot_number_for_registration_number":
		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		actionParking, msg = actionParking.FindByRegis(input[1])
		log.Println(msg)
	case "slot_numbers_for_cars_with_colour":
		if len(actionParking.Data) == 0 {
			log.Fatal("Must CreateParkingLot first ")
		}
		actionParking, msg = actionParking.FindByLot(input[1])
		log.Println(msg)
	case "exit":
		os.Exit(3)
	}
	return actionParking

}

func FileInputParam(path string) []string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("cannot load file ", err)
	}
	tags := strings.Split(string(dat), "\n")
	return tags
}
