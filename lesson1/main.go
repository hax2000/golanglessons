package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {

	ntpTime, err := ntp.Time("time.apple.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(3)
	}else{
		hours, minutes, seconds := ntpTime.Clock()
		fmt.Printf("Текущее время: %d:%d:%d\n",hours,minutes,seconds)
		fmt.Printf("Точное время: %d:%d:%d\n",hours-3,minutes,seconds)
	}

}