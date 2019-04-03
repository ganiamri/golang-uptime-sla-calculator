package main

import (
	"log"

	slacalc "github.com/ganiamri/golang-uptime-sla-calculator/sla-calculator"
)

type UptimeData struct {
	Timestamp int64
	Value     int64
	Exception bool
}

var (
	startTime int64 = 10000
	endTime   int64 = 13000
)

var uptimeSeriesData = []UptimeData{
	{10100, 0, false},
	{10200, 0, false},
	{10300, 0, false},
	{10400, 0, false},
	{10500, 270, false},
	{10600, 370, false},
	{10700, 470, false},
	{10800, 570, false},
	{10900, 670, false},
	{11000, 40, false},
	{11100, 140, false},
	{11200, 240, false},
	{11300, 340, false},
	{11400, 440, false},
	{11500, 540, false},
	{11600, 0, false},
	{11700, 0, false},
	{11800, 840, false},
	{11900, 940, false},
	{12000, 1040, false},
	{12100, 1140, false},
	{12200, 0, false},
	{12300, 0, false},
	{12400, 0, false},
	{12500, 10, false},
	{12600, 110, false},
	{12700, 210, false},
	{12800, 0, true},
	{12900, 0, true},
	{13000, 0, false},
}

func main() {
	// Prepare the data
	uptimeVals := []int64{}
	timestamps := []int64{}
	exceptions := []bool{}
	for _, val := range uptimeSeriesData {
		uptimeVals = append(uptimeVals, val.Value)
		timestamps = append(timestamps, val.Timestamp)
		exceptions = append(exceptions, val.Exception)
	}
	// Create calculator object
	calc := slacalc.NewUptimeSLACalculator()
	// Print the calculation result
	snmpSLA, _ := calc.CalculateSNMPAvailability(startTime, endTime, timestamps, uptimeVals)
	log.Printf("Connectivity SLA: %.5f", snmpSLA)
	uptimeSLA, _ := calc.CalculateUptimeAvailability(startTime, endTime, timestamps, uptimeVals)
	log.Printf("Uptime SLA: %.5f", uptimeSLA)
	SLA1, _ := calc.CalculateSLA1Availability(startTime, endTime, timestamps, uptimeVals)
	log.Printf("(Custom) SLA 1: %.5f", SLA1)
	SLA2, _ := calc.CalculateSLA2Availability(startTime, endTime, timestamps, uptimeVals, exceptions)
	log.Printf("(Custom) SLA 2: %.5f", SLA2)
}
