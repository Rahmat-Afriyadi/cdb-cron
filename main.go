package main

import (
	"cron/service"
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func Log(content string) {
	fileName := "log.txt"

	// Open the file for appending (create if it doesn't exist)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append the new line to the file
	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	Log("Hari ini komputer dinyalakan " + time.Now().Format("2006-01-02 15:04:05"))

	defer scheduler.Stop()
	_, err := scheduler.AddFunc("15 18 * * *", func() {
		a := time.Now()
		if a.Weekday() == 5 && (a.Format("02") == "12" || a.Format("02") == "13") {
			Log("running task " + time.Now().Format("2006-01-02 15:04:05"))
			service.UpdateMohonFakturWKMTDKValid()
		}
		if a.Format("02") == "14" {
			Log("running task " + time.Now().Format("2006-01-02 15:04:05"))
			service.UpdateMohonFakturWKMTDKValid()
		}
		Log("running task " + time.Now().Format("2006-01-02 15:04:05") + " Hari ini tidak update")
	})
	if err != nil {
		fmt.Println("cron error ", err)
	}

	// _, err = scheduler.AddFunc("* * * * *", func() {
	// 	a := time.Now()
	// 	Log("running task " + a.Format("2006-01-02 15:04:05") + " Test Menit")
	// })
	// if err != nil {
	// 	fmt.Println("cron error ", err)
	// }

	scheduler.Start()

	select {}

}

// {JM04E1872694 08979620183 08979620183 ardidwicahyono@gmail.com 2b}
