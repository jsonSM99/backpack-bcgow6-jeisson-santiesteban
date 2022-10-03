package main

import (
	"fmt"
)

type CATEGORIES = map[string]map[string]float64
type WORKER = map[string]interface{}
type WORKERS = map[string]WORKER

var categories CATEGORIES = CATEGORIES{
	"C": {
		"salary_per_hour": 1000,
		"percent":         0,
	},
	"B": {
		"salary_per_hour": 1500,
		"percent":         .2,
	},
	"A": {
		"salary_per_hour": 3000,
		"percent":         .5,
	},
}

var workers WORKERS = WORKERS{
	"jeissom": {
		"category":   "A",
		"minutes":    1000,
		"month_days": 32,
	},
	"fernando": {
		"category":   "B",
		"minutes":    3000,
		"month_days": 35,
	},
}

func calcular_salario(worker WORKER) float64 {
	var result float64

	minutes, ok_minutes := worker["minutes"].(int)
	category, ok_category := worker["category"].(string)
	month_days, ok_month_days := worker["month_days"].(int)

	if !ok_category || !ok_minutes || !ok_month_days {
		return result
	}
	salary_per_hour, ok_salary_per_hour := categories[category]["salary_per_hour"]

	if !ok_salary_per_hour {
		return result
	}

	hours := (minutes / 60)
	result = float64(hours) * salary_per_hour * float64(month_days)

	return result
}

func main() {
	for name, worker := range workers {
		fmt.Printf("Trabajador %s: %.2f\n", name, calcular_salario(worker))
	}
}
