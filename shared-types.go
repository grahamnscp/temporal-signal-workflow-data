package signalwkfldata

import (
	"os"
	"time"
	"strings"
)

var TaskQueueName = os.Getenv("TASK_QUEUE_NAME")
var NumberOfCustomers = 1
var PeriodLengthSeconds = 30
var log_level = strings.ToLower(os.Getenv("LOG_LEVEL"))

type Subscription struct {
	BillingPeriod       time.Duration
	MaxBillingPeriods   int
	BillingPeriodCharge int
	Active              bool
}

type Customer struct {
	FirstName    string
	LastName     string
	Id           string
	Email        string
	Subscription Subscription
}

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[94m"
var ColorMagenta = "\033[35m"
var ColorCyan = "\033[36m"
var ColorWhite = "\033[37m"

