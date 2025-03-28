package observer_test

import (
	"fmt"
	. "gof-in-go/behavioral/observer"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	generator := NewRandomNumberGenerator()
	digitObserver := NewDigitObserver()
	graphObserver := NewGraphObserver()

	generator.AddObserver(digitObserver)
	generator.AddObserver(graphObserver)

	generator.Execute()
	allOutputs := generator.NotifyObservers()

	digitRegex := regexp.MustCompile(`^DigitObserver:(\d+)$`)
	graphRegex := regexp.MustCompile(`^GraphObserver:(\**)$`)

	digitOutputs := make([]string, 0)
	graphOutputs := make([]string, 0)

	for _, output := range allOutputs {
		if digitRegex.MatchString(output) {
			digitOutputs = append(digitOutputs, output)
		} else if graphRegex.MatchString(output) {
			graphOutputs = append(graphOutputs, output)
		}
	}

	assert.Len(t, digitOutputs, 20, "DigitObserverの出力回数が20回ではありません")
	assert.Len(t, graphOutputs, 20, "GraphObserverの出力回数が20回ではありません")

	for i := range 20 {
		digitMatch := digitRegex.FindStringSubmatch(digitOutputs[i])
		graphMatch := graphRegex.FindStringSubmatch(graphOutputs[i])

		if assert.Len(t, digitMatch, 2, fmt.Sprintf("%d回目のDigitObserverの出力形式が不正です: %s", i+1, digitOutputs[i])) &&
			assert.Len(t, graphMatch, 2, fmt.Sprintf("%d回目のGraphObserverの出力形式が不正です: %s", i+1, graphOutputs[i])) {

			digitStr := digitMatch[1]
			graphStr := graphMatch[1]

			digitNum, err := strconv.Atoi(digitStr)
			assert.NoError(t, err, fmt.Sprintf("%d回目のDigitObserverの数値を整数に変換できませんでした: %s", i+1, digitStr))

			assert.Equal(t, len(graphStr), digitNum, fmt.Sprintf("%d回目のDigitObserver (%d) と GraphObserver (%d) の数値が一致しません", i+1, digitNum, len(graphStr)))
		}
	}
}
