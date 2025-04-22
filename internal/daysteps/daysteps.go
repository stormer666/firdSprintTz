package daysteps

import (
	"/c/Users/storm/petprojects/firdSprintTz/spentcalories.go"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000

	errParse = errors.New("ошибка парсинга строки")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	//Деление строки на слайс строк
	stepsndWalk := strings.Split(data, ",")

	//Проверка длины слайса

	if len(stepsndWalk) != 2 {
		return 0, 0, err
	}
	//Преобразование первого элемента слайса в число типа int
	steps, err := int(stepsndWalk[0])
	if err != nil {
		return 0, 0, err
	}
	if steps == 0 {
		return 0, 0, errors.New("количество шагов не должно быть равным нулю")
	}
	//Преобразование второго элемента слайса в time.Duration
	timeOfWalk, err := time.ParseDuration(stepsndWalk[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, timeOfWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {

	//Парсинг строки с данными
	walkDetails, err := parsePackage(data)
	if err != nil {
		fmt.Println("ошибка парсинга строки")
		return ""
	}

	step := walkDetails[0]
	if walkDetails[0] == 0 {
		return ""
	}

	distanceM := (float64(step) * stepLength) / mInKm

	ccals := spentcalories.WalkingSpentCalories()

	output := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", step, distanceM, ccals)
	return output

}
