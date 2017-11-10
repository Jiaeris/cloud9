package utils

import "time"

func TimeInt64() int64 {
	return time.Now().Unix()
}

//if parse error ,return time now object .
func String2Time(stime string) time.Time {
	timeObj, err := time.Parse("2006-01-02 15:04:05", stime)
	if err != nil {
		panic(err)
		return time.Now()
	}
	return timeObj
}
