package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jonfriesen/adventofcode2018/util"
)

func main() {

	// load input into <time> <string> list
	dic := []DataInput{}
	util.LoadInputFromPath("input", createEntryFunc(&dic))

	// sort
	sort.Slice(dic, func(i, j int) bool { return dic[i].timestamp.Before(dic[j].timestamp) })

	// seperate all Shifts held in Guards (add to GuardCollection)

	grouped := pullShiftFromList(dic)

	guardMap := make(map[int]*GuardProfile)

	for _, v := range grouped {
		shift, id := buildShift(v)
		if guardMap[id] == nil {
			guardMap[id] = &GuardProfile{
				shifts: []Shift{},
			}
		}
		guardMap[id].shifts = append(guardMap[id].shifts, shift)
	}

	// Iterate over guards finding sum of total sleep (keep track of the current guard that sleeps the most)
	topSleeperID := -1
	topSleeperQty := 0

	for i, v := range guardMap {
		gSum := 0
		g := *v
		for _, s := range g.shifts {
			gSum += s.sumSleep()
		}

		if gSum > topSleeperQty {
			topSleeperQty = gSum
			topSleeperID = i
		}
	}

	fmt.Println("Top Sleeper:", topSleeperID)
	fmt.Println("Top Sleeper Qty:", topSleeperQty)

	sleepiestGuard := *guardMap[topSleeperID]
	sleepiestGuard.getMostSleptMinute()

	fmt.Println("Most slept minute", sleepiestGuard.sleepiestMinute)

	fmt.Println("Code:", (topSleeperID * sleepiestGuard.sleepiestMinute))

	// part 2
	mostTimeSleptAtSameMinuteQty := -1
	mostTimeSleptAtSameMinuteID := -1
	for i, v := range guardMap {
		v.getMostSleptMinute()
		if v.sleepiestMinuteCount > mostTimeSleptAtSameMinuteQty {
			mostTimeSleptAtSameMinuteQty = v.sleepiestMinuteCount
			mostTimeSleptAtSameMinuteID = i
		}
	}

	fmt.Println("Guard with most consistent sleep minute is ", mostTimeSleptAtSameMinuteID, "at a count of ", mostTimeSleptAtSameMinuteQty, "for a secret code of", (mostTimeSleptAtSameMinuteID * guardMap[mostTimeSleptAtSameMinuteID].sleepiestMinute))
}

type ShiftEventType int

const (
	Awake ShiftEventType = 0
	Sleep ShiftEventType = 1
)

type GuardCollection struct {
	guards map[int]*GuardProfile
}

type GuardProfile struct {
	// list of events (sorted)
	shifts               []Shift
	sleepiestMinute      int
	sleepiestMinuteCount int
}

type Shift struct {
	// get start time + end time
	startTime time.Time
	endTime   time.Time
	// hold events of session (these are presorted)
	events []ShiftEvent
}

type ShiftEvent struct {
	eType ShiftEventType
	eTime time.Time
}

type DataInput struct {
	timestamp time.Time
	content   string
}

func (g *GuardProfile) getMostSleptMinute() {
	minSleepCount := make(map[int]int)

	startSleep := -1
	for _, v := range g.shifts {
		for _, e := range v.events {
			if e.eType == Sleep {
				startSleep = e.eTime.Minute()
			}
			if e.eType == Awake {
				for c := startSleep; c < e.eTime.Minute(); c++ {
					minSleepCount[c]++
				}
			}
		}
	}

	mostSleptMinuteQty := -1
	mostSleptMinute := -1
	for m, ms := range minSleepCount {
		if ms > mostSleptMinuteQty {
			mostSleptMinuteQty = ms
			mostSleptMinute = m
		}
	}

	g.sleepiestMinute = mostSleptMinute
	g.sleepiestMinuteCount = mostSleptMinuteQty
}

func (s *Shift) sumSleep() int {
	var startSleep time.Time
	totalSleep := 0
	isSleeping := false
	for _, v := range s.events {
		if v.eType == Sleep {
			isSleeping = true
			startSleep = v.eTime
		}
		if v.eType == Awake {
			isSleeping = false
			totalSleep += int(v.eTime.Sub(startSleep).Minutes())
		}
	}

	if isSleeping {
		totalSleep += int(s.endTime.Sub(startSleep).Minutes())
	}

	return totalSleep
}

func pullShiftFromList(dic []DataInput) [][]DataInput {
	out := [][]DataInput{}
	var di []DataInput
	for _, v := range dic {
		if strings.Contains(v.content, "begins shift") {
			if di != nil {
				out = append(out, di)
			}
			di = []DataInput{}
		}
		di = append(di, v)
	}

	// append the last record ;)
	out = append(out, di)

	return out
}

func buildShift(di []DataInput) (Shift, int) {
	s := Shift{}
	id := -1
	for _, v := range di {
		if strings.Contains(v.content, "begins shift") {
			s.startTime = v.timestamp
			id = getIDFromInput(v.content)
		}
		if strings.Contains(v.content, "falls asleep") {
			s.events = append(s.events, ShiftEvent{Sleep, v.timestamp})
		}
		if strings.Contains(v.content, "wakes up") {
			s.events = append(s.events, ShiftEvent{Awake, v.timestamp})
		}
	}
	lastEvent := di[len(di)-1].timestamp
	s.endTime = time.Date(
		lastEvent.Year(),
		lastEvent.Month(),
		lastEvent.Day(),
		0,  // hour
		59, // min
		0,
		0,
		lastEvent.Location(),
	)
	return s, id
}

// This function loads the input into an array /w
// time + the value of that line
func createEntryFunc(dic *[]DataInput) util.InputFunc {

	return func(line string) {
		*dic = append(*dic, DataInput{getTimeStampFromLine(line), line})
	}
}

func getTimeStampFromLine(line string) time.Time {
	re := regexp.MustCompile(`(?m)\[(.*?)\]`)

	in := re.FindStringSubmatch(line)
	layout := "2006-01-02 15:04"

	t, _ := time.Parse(layout, in[1])

	return t
}

func getIDFromInput(line string) int {
	s := strings.Split(line, " ")
	for _, v := range s {
		if strings.HasPrefix(v, "#") {
			id, _ := strconv.Atoi(strings.TrimPrefix(v, "#"))
			return id
		}
	}
	return -1
}
