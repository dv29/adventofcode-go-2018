package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type sleepDuration struct {
	from, to time.Time
}
type sleepPattern struct {
	totalMinutes float64
	sleepTimes   []sleepDuration
}

type GuardSleepPattern map[int]map[string]*sleepPattern

func (g GuardSleepPattern) PrintSleepTimes(currentGuardId int, date string) {
	// fmt.Printf("cgi: %d\n", currentGuardId)
	// fmt.Printf("date: %s\n", date)
	for _, v := range g[currentGuardId][date].sleepTimes {
		fmt.Printf("sleeptime from: %+v/t to %+v\n", v.from, v.to)
	}
	// fmt.Printf("sleepTimes: %v\n", g[currentGuardId][date].sleepTimes)
	// fmt.Printf("curr guard: %v\n", g[currentGuardId][date])
	// fmt.Printf("totalMinutes: %f\n", g[currentGuardId][date].totalMinutes)
}

func (g GuardSleepPattern) CalculateTotalMinutes(currentGuardId int, date string) {
	total := 0.0
	for _, v := range g[currentGuardId][date].sleepTimes {
		total += v.to.Sub(v.from).Minutes()
	}

	g[currentGuardId][date].totalMinutes = total
}

func (g GuardSleepPattern) GetGuardMaxSleep() (guardId int) {
	currMax := 0.0
	for guard := range g {

		// fmt.Printf("%v\t%v\n", guard, guardDate)
		total := 0.0
		for guardDate := range g[guard] {
			total += g[guard][guardDate].totalMinutes
		}

		if total > currMax {
			currMax = total
			guardId = guard
		}
	}
	return
}

func (g GuardSleepPattern) GetMaxOverlap(guardId int) (minute, minuteCount int) {
	minuteCountMap := make(map[int]int)
	for gd := range g[guardId] {
		for i := 0; i < 60; i++ {
			for _, v := range g[guardId][gd].sleepTimes {
				if i >= v.from.Minute() && i < v.to.Minute() {
					// fmt.Printf("sleeptime from: %+v/t to %+v\n", v.from.Minute(), v.to.Minute())
					minuteCountMap[i]++
				}
			}
		}
	}
	currMax := 0
	for k, v := range minuteCountMap {
		if v > currMax {
			currMax = v
			minute = k
			minuteCount = currMax
		}
	}
	return
}

func (g GuardSleepPattern) GetGuardMaxOverlap() (gid, min int) {
	type Minute struct{ minute, minuteCount int }
	guardCount := make(map[int]Minute)
	for guardId := range g {
		minute, minuteCount := g.GetMaxOverlap(guardId)
		guardCount[guardId] = Minute{minute, minuteCount}
	}
	currMax := 0
	for k, v := range guardCount {
		if v.minuteCount > currMax {
			currMax = v.minuteCount
			gid = k
			min = v.minute
		}
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var strs []string
	guards := make(GuardSleepPattern)
	// regex := regexp.MustCompile(`\[(?P<year>\d+)-(?P<month>\d+)-(?P<day>\d+) (?P<hour>\d+):(?P<minute>\d+)\]`)
	// value := regex.FindStringSubmatch(v)

	for scanner.Scan() {
		str := scanner.Text()
		check(err)
		strs = append(strs, str)
	}

	sort.Strings(strs)
	currentGuardId := 0
	for _, v := range strs {
		var hh, mm int
		var date string
		fmt.Sscanf(v, "[%s %d:%d] Guard #%d", &date, &hh, &mm, &currentGuardId)

		if _, ok := guards[currentGuardId]; !ok {
			guards[currentGuardId] = make(map[string]*sleepPattern)
		}
		if _, ok := guards[currentGuardId][date]; !ok {
			guards[currentGuardId][date] = &sleepPattern{}
		}

		currTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %02d:%02d:00", date, hh, mm))
		check(err)

		if strings.Contains(v, "falls") {
			// fmt.Printf("falling asleep\n")
			guards[currentGuardId][date].sleepTimes = append(guards[currentGuardId][date].sleepTimes, sleepDuration{from: currTime})
			// guards.PrintSleepTimes(currentGuardId, date)
		} else if strings.Contains(v, "wake") {
			// fmt.Printf("waking up\n")
			guards[currentGuardId][date].sleepTimes[len(guards[currentGuardId][date].sleepTimes)-1].to = currTime
			guards.CalculateTotalMinutes(currentGuardId, date)
			// guards.PrintSleepTimes(currentGuardId, date)
		}
	}

	guardMaxSlept := guards.GetGuardMaxSleep()
	// for gd := range guards[guardMaxSlept] {
	// 	guards.PrintSleepTimes(guardMaxSlept, gd)
	// }
	maxMinute, _ := guards.GetMaxOverlap(guardMaxSlept)
	fmt.Printf("%v\n", maxMinute*guardMaxSlept)

	gid, min := guards.GetGuardMaxOverlap()
	fmt.Printf("Guard with most overlap time guardid: %v, minute: %v, total = %v\n", gid, min, gid*min)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
