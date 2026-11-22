package spentenergy

import (
	"errors"
	"math"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func spentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, errors.New("количество шагов должно быть больше нуля")
	case !isPositiveFinite(weight):
		return 0, errors.New("вес должен быть положительным конечным числом")
	case !isPositiveFinite(height):
		return 0, errors.New("рост должен быть положительным конечным числом")
	case duration <= 0:
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	speed := MeanSpeed(steps, height, duration)

	return weight * speed * duration.Minutes() / minInH, nil
}

func isPositiveFinite(value float64) bool {
	return value > 0 && !math.IsInf(value, 0)
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := spentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	return spentCalories(steps, weight, height, duration)
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}
