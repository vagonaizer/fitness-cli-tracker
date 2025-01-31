package ui

import (
	"bufio"
	"fitness-cli-tracker/internal/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetInput() (models.Record, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("вес: ")
	weightStr, _ := reader.ReadString('\n')
	weight, err := strconv.ParseFloat(strings.TrimSpace(weightStr), 64)
	if err != nil {
		return models.Record{}, err
	}

	fmt.Print("Тренировался седня?: (y/n)")
	trainedStr, _ := reader.ReadString('\n')
	trained := strings.TrimSpace(trainedStr) == "y"

	fmt.Print("Калории: ")
	caloriesStr, _ := reader.ReadString('\n')
	calories, err := strconv.Atoi(strings.TrimSpace(caloriesStr))
	if err != nil {
		return models.Record{}, nil
	}

	return models.Record{
		Date:     time.Now(),
		Weight:   weight,
		Trained:  trained,
		Calories: calories,
	}, nil
}
