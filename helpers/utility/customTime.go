package utility

import (
	"strconv"
	"strings"
	"time"
)

func AddTime(Minute int) string {
	timein := time.Now().Local().Add(time.Minute * time.Duration(Minute))
	return timein.Format("15:04")
}

func splitTime(time string, index int) string {
	aArr := strings.Split(time, ":")

	return aArr[index]
}

func GetTime() string {
	timein := time.Now()
	return timein.Format("15:04")
}

func TimeLimit(currentTime, TimeNow string) bool {
	var check bool = false
	hour1 := splitTime(currentTime, 0)
	hour2 := splitTime(TimeNow, 0)
	timeNow, _ := strconv.ParseInt(splitTime(TimeNow, 1), 10, 64)
	timeCurrent, _ := strconv.ParseInt(splitTime(currentTime, 1), 10, 64)
	if currentTime == TimeNow {
		check = false
	} else if hour1 != hour2 {
		check = true
	} else if timeCurrent < timeNow {
		check = true
	}
	return check
}
