package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("не удалось разобрать %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось сформировать отчёт для записи %q: %v", data, err)
			continue
		}

		fmt.Println(strings.TrimSuffix(info, "\n"))
	}
}
