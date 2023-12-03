package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"

	"example.com/error-handling/db"
	"example.com/error-handling/utils"

	"github.com/google/uuid"
)

var mockDb *db.MockDb

func CalculateHourWage(salary float64, hour int) (result float64, err error) {
	if hour == 0 {
		err = utils.NewApiError(
			uuid.New().String(),
			http.StatusBadRequest,
			utils.BadRequestCode,
			"can not calculate hourly wage for person without wouring hours",
			errors.New("zero devisor error"),
		)
		return
	}

	result = salary / float64(hour)
	return
}

func CalculatePersonHourWage(name string) (result int, err error) {
	person, err := mockDb.GetPerson(name)
	if err != nil {
		return
	}

	floatResult, err := CalculateHourWage(person.Salary, person.WorkingHours)
	if err != nil {
		return
	}

	result = int(math.Floor(floatResult))
	return
}

func init() {
	mockDb = db.NewMockDb()
	mockDb.AddPerson(db.Person{"Domingo", 1000.0, 80})
	mockDb.AddPerson(db.Person{"Domingooo", 2000.0, 0})
}

func main() {
	result, err := CalculatePersonHourWage("Domingo")
	if err != nil {
		utils.HandleError(err)
	} else {
		fmt.Printf("Result of hour salary for %s is %d\n", "Domingo", result)
	}

	result, err = CalculatePersonHourWage("SomeName")
	if err != nil {
		utils.HandleError(err)
	} else {
		fmt.Printf("Result of hour salary for %s is %d\n", "SomeName", result)
	}

	result, err = CalculatePersonHourWage("Domingooo")
	if err != nil {
		utils.HandleError(err)
	} else {
		fmt.Printf("Result of hour salary for %s is %d\n", "SomeName", result)
	}
}

//func main() {
//	names := []string{"Domingo", "SomeName", "Domingooo"}
//	for _, name := range names {
//		result, err := CalculatePersonHourWage(name)
//		if err != nil {
//			utils.HandleError(err)
//		} else {
//			fmt.Printf("Hourly wage for %s is %d\n", name, result)
//		}
//	}
//}
