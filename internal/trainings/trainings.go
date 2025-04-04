package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// создайте структуру Training
type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")

	if len(s) != 3 {
		return errors.New("Invalid data\n")
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	if s[1] != "Бег" && s[1] != "Ходьба" {
		return errors.New("Invalid data\n")
	}
	t.TrainingType = s[1]

	trainingTime, err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	t.Duration = trainingTime

	return nil

}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", errors.New("Invalid data\n")
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	ccals := 0.0
	switch t.TrainingType {
	case "Бег":
		ccals = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		ccals = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type\n")
	}

	res := fmt.Sprintf("Тип тренировки: %s"+
		"\nДлительность: %.2f ч."+
		"\nДистанция: %.2f км."+
		"\nСкорость: %.2f км/ч"+
		"\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, ccals)

	return res, nil

}
