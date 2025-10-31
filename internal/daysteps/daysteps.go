package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataParts := strings.Split(data, ",")
	if len(dataParts) != 2 {
		return 0, 0, fmt.Errorf("wrong arguments")
	}

	countSteps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return 0, 0, err
	}

	if countSteps <= 0 {
		return 0, 0, fmt.Errorf("negative steps")
	}

	duration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("wrong duration")
	}

	return countSteps, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	countSteps, walkDuration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distanceM := float64(countSteps) * stepLength
	distanceKm := distanceM / mInKm

	calories, err := spentcalories.WalkingSpentCalories(countSteps, weight, height, walkDuration)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		countSteps, distanceKm, calories)

}
