package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Field ranges for cron expressions
var (
	minuteRange     = []int{0, 59}
	hourRange       = []int{0, 23}
	dayOfMonthRange = []int{1, 31}
	monthRange      = []int{1, 12}
	dayOfWeekRange  = []int{0, 6}
)

func main() {

	var cronExpFlag = flag.String("exp", "", "cronExpression")
	flag.Parse()
	cronExp := *cronExpFlag

	fmt.Println("cron expression:: ", cronExp)

	var fields = strings.Fields(cronExp)
	if len(fields) < 6 || len(fields) > 6 {
		fmt.Println("Invalid cron expression. Expected 5 time fields and a command.")
		fmt.Println(fields)
		os.Exit(1)
	}

	minute := fieldParser(fields[0], minuteRange)
	hour := fieldParser(fields[1], hourRange)
	dayOfMonth := fieldParser(fields[2], dayOfMonthRange)
	month := fieldParser(fields[3], monthRange)
	dayOfWeek := fieldParser(fields[4], dayOfWeekRange)
	command := fields[5] //expecting single word command

	fmt.Printf("%-14s %v\n", "minute", minute)
	fmt.Printf("%-14s %v\n", "hour", hour)
	fmt.Printf("%-14s %v\n", "day of month", dayOfMonth)
	fmt.Printf("%-14s %v\n", "month", month)
	fmt.Printf("%-14s %v\n", "day of week", dayOfWeek)
	fmt.Printf("%-14s %s\n", "command", command)
}
