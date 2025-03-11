package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cron_parser \"cron_expression\"")
		os.Exit(1)
	}

	cronExp := os.Args[1]
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
	command := strings.Join(fields[6], " ")

	fmt.Printf("%-14s %v\n", "minute", minute)
	fmt.Printf("%-14s %v\n", "hour", hour)
	fmt.Printf("%-14s %v\n", "day of month", dayOfMonth)
	fmt.Printf("%-14s %v\n", "month", month)
	fmt.Printf("%-14s %v\n", "day of week", dayOfWeek)
	fmt.Printf("%-14s %s\n", "command", command)
}
