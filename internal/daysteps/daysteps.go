package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")

	if len(s) != 2 {
		return errors.New("Invalid input\n")
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("Invalid duration\n")
	}

	distance := spentenergy.Distance(ds.Steps)

	ccals := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	res := fmt.Sprintf("Количество шагов: %d."+
		"\nДистанция составила %.2f км."+
		"\nВы сожгли %.2f ккал.\n", ds.Steps, distance, ccals)

	return res, nil
}
