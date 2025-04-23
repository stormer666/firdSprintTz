package daysteps

import (
	"errors"
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

	//Деление строки на слайс строк
	stepsndWalk := strings.Split(data, ",")

	//Проверка длины слайса

	if len(stepsndWalk) != 2 {
		err := errors.New("неверная длина слайса")
		log.Println(err)
		return 0, 0, err

	}
	//Преобразование первого элемента слайса в число типа int
	steps, err := strconv.Atoi(stepsndWalk[0])
	if err != nil {
		err := errors.New("ошибка преобразования")
		log.Println(err)
		return 0, 0, err

	}
	if steps <= 0 {
		err := errors.New("количество шагов не должно быть меньше либо равным нулю")
		log.Println(err)
		return 0, 0, err
	}

	//Преобразование второго элемента слайса в time.Duration
	timeOfWalk, err := time.ParseDuration(stepsndWalk[1])
	if err != nil {
		err := errors.New("ошибка парсирования времени")
		log.Println(err)
		return 0, 0, err
	}
	if timeOfWalk <= 0.0 {
		err := errors.New("продолжительность не должна быть меньше либо равна нулю")
		log.Println(err)
		return 0, 0, err
	}
	return steps, timeOfWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {

	//Парсинг строки с данными
	walkDetails, timeWalk, err := parsePackage(data)
	if err != nil {
		fmt.Println("ошибка парсинга строки")
		return ""
	}

	if walkDetails <= 0 {
		return ""
	}

	distanceM := (float64(walkDetails) * stepLength) / mInKm

	ccals, err := spentcalories.WalkingSpentCalories(walkDetails, weight, height, timeWalk)
	if err != nil {
		return ""
	}

	output := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", walkDetails, distanceM, ccals)
	return output

}
