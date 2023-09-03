package interviews

var totalNumberOfRequests int
var droppedRequests int

var secondLimit int = 3
var minuteLimit int = 10
var fiveMinuteLimit int = 20

// Implement a fixed counter rate limiter
// Can accept a maximum of
//
//	3 requests every second
//	10 requests every minute
//	40 requests every 5minutes
//
// If a request is dropped, it is still counted in the total requests for the period
// return the number of dropped requests
func rateLimitRequests(requests []int) int {

	// reset counters for requests count in interval

	var secondCounter int = 0
	var minuteCounter int = 0
	var fiveMinuteCounter int = 0

	// loop through requests
	// if second changes, reset second counter
	// at every 60 seconds, reset minuteCounter (using division)
	// at every 300 seconds, reset fiveMinuteCounter (using division)

	// check from longest to shortest time

	// store curr and prev time intervals

	currSecond := requests[0] - 1
	currMinute := (requests[0] / 60) - 1
	currFiveMinute := (requests[0] / 300) - 1

	for i := range requests {

		if requests[i]/300 != currFiveMinute {
			currFiveMinute = requests[i] / 300
			fiveMinuteCounter = 0
		}

		if requests[i]/60 != currMinute {
			currMinute = requests[i] / 60
			minuteCounter = 0
		}

		if requests[i] != currSecond {
			currSecond = requests[i]
			secondCounter = 0
		}

		totalNumberOfRequests++

		if fiveMinuteCounter+1 > fiveMinuteLimit || minuteCounter+1 > minuteLimit || secondCounter+1 > secondLimit {
			droppedRequests++
		}
		fiveMinuteCounter++
		minuteCounter++
		secondCounter++
	}

	//fmt.Println(fiveMinuteCounter, minuteCounter, secondCounter)

	return droppedRequests
}
