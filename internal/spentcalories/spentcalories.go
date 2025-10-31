package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataParts := strings.Split(data, ",")
	if len(dataParts) != 3 {
		return 0, "", 0, fmt.Errorf("not have arguments")
	}

	typeActive := dataParts[1]

	countSteps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return 0, "", 0, err
	}

	if countSteps <= 0 {
		return 0, "", 0, fmt.Errorf("negative steps")
	}

	durationActiv, err := time.ParseDuration(dataParts[2])
	if err != nil {
		return 0, "", 0, err
	}

	if durationActiv <= 0 {
		return 0, "", 0, fmt.Errorf("negative duration")
	}

	return countSteps, typeActive, durationActiv, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient

	distanceM := float64(steps) * stepLength
	distanceKM := distanceM / mInKm

	return distanceKM
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0.0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeActive, durationTime, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	hours := durationTime.Hours()
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, durationTime)

	var calories float64
	switch typeActive {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, durationTime)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, durationTime)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		log.Println(err)
		return "", err
	}

	answer := `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`

	return fmt.Sprintf(answer, typeActive, hours, dist, speed, calories), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0.0 || height <= 0.0 {
		return 0, fmt.Errorf("running spCal: wrong arguments")
	}

	if duration <= 0.0 {
		return 0, fmt.Errorf("wrong duration")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0.0 || height <= 0.0 {
		return 0, fmt.Errorf("walking spCal: wrong arguments")
	}

	if duration <= 0.0 {
		return 0, fmt.Errorf("wrong duration")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}
