package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	slicestr := strings.Split(data, ",")

	if len(slicestr) != 3 {
		err := errors.New("неверное количество элементов")
		log.Println(err)
		return 0, "", 0, err
	}

	steps, err := strconv.Atoi(slicestr[0])
	if err != nil {
		err := errors.New("ошибка преобразования количества шагов")
		log.Println(err)
		return 0, "", 0, err
	}

	if steps <= 0 {
		err := errors.New("количество шагов должно быть больше нуля")
		log.Println(err)
		return 0, "", 0, err
	}

	kindOfActivity := slicestr[1]

	timeOfWalk, err := time.ParseDuration(slicestr[2])
	if err != nil {
		err := errors.New("ошибка парсинга времени")
		log.Println(err)
		return 0, "", 0, err
	}
	if timeOfWalk <= 0 {
		err := errors.New("продолжительность должна быть больше нуля")
		log.Println(err)
		return 0, "", 0, err
	}

	return steps, kindOfActivity, timeOfWalk, nil
}

func distance(steps int, height float64) float64 {
	stepsLength := height * stepLengthCoefficient
	stepsFinished := float64(steps) * stepsLength
	distanceM := stepsFinished / mInKm
	return distanceM
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	durHours := duration.Hours()
	averageSpeed := dist / durHours
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	howSteps, typeOfActivity, howTime, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var output string
	hours := howTime.Seconds() / 3600
	howLong := distance(howSteps, height)
	avgSp := meanSpeed(howSteps, height, howTime)

	switch typeOfActivity {
	case "Бег":
		spentCcal, err := RunningSpentCalories(howSteps, weight, height, howTime)
		if err != nil {
			return "", err
		}
		output = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfActivity, hours, howLong, avgSp, spentCcal)

	case "Ходьба":
		spentCcal, err := WalkingSpentCalories(howSteps, weight, height, howTime)
		if err != nil {
			return "", err
		}
		output = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfActivity, hours, howLong, avgSp, spentCcal)

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	return output, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		err := errors.New("некорректные значения steps")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("некорректные значения weight")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("некорректные значения height")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("некорректные значения duration")
		log.Println(err)
		return 0, err
	}

	averageSp := meanSpeed(steps, height, duration)

	durationMin := duration.Minutes()
	result := (weight * averageSp * durationMin) / minInH

	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		err := errors.New("некорректные значения steps")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("некорректные значения weight")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("некорректные значения height")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("некорректные значения duration")
		log.Println(err)
		return 0, err
	}

	averageSp := meanSpeed(steps, height, duration)

	durationMin := duration.Minutes()
	preResult := (weight * averageSp * durationMin) / minInH
	result := preResult * walkingCaloriesCoefficient

	return result, nil
}
