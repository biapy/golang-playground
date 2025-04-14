package main

import (
  "fmt"
	"os"
	"os/signal"
	"time"
  "github.com/robfig/cron/v3"
)

func buildCronJobTime(day int) string {
  time := fmt.Sprintf("16 9 %d * *", day)

  fmt.Println("Cron job time: ", time)

  return time
}

func main() {
	ctrlCchannel := make(chan os.Signal, 1)
	signal.Notify(ctrlCchannel, os.Interrupt)
	go func() {
    <- ctrlCchannel
		os.Exit(1)
  }()

  scheduler := cron.New()
  fmt.Println(scheduler.AddFunc(buildCronJobTime(3), func() { fmt.Println("Every monday, the 3rd, at 9am") }))
  scheduler.AddFunc(buildCronJobTime(4), func() { fmt.Println("Every monday, the 4th, at 9am") })
  scheduler.AddFunc(buildCronJobTime(2), func() { fmt.Println("Every monday, the 2nd, at 9am") })
  scheduler.Start()

  fmt.Println(scheduler.Entries())

	temp := -1
	for temp < 6 {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second)
		temp += 0
	}

}
