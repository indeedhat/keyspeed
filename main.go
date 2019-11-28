package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MarinX/keylogger"
	"github.com/indeedhat/gli"
)

const (
	WordKeys = 5
)

type KeySpeed struct {
	Help     bool   `gli:"help,h" description:"display this message"`
	Cpm      bool   `gli:"cpm,c" description:"display characters per minute (wpm is default)"`
	Pad      int    `gli:"pad,p" description:"pad the display with leading 0's" default:"3"`
	Best     bool   `gli:"best,b" description:"keep track of your best time for this session"`
	Interval int64  `gli:"interval,i" description:"Polling interval to uodate the count in seconds" default:"5"`
	Device   string `gli:"device,d" description:"manually set the keyboard device"`

	log  []int64
	best int
}

func (ks *KeySpeed) NeedHelp() bool {
	return ks.Help
}

func (ks *KeySpeed) Run() int {
	if "" == ks.Device {
		ks.Device = keylogger.FindKeyboardDevice()
	}

	if 1 > len(ks.Device) {
		log.Fatal("no keyboard found")
	}

	logger, err := keylogger.New(ks.Device)
	if nil != err {
		log.Fatal(err)
	}

	defer logger.Close()

	events := logger.Read()
	go ks.watcher()

	for ev := range events {
		switch ev.Type {
		case keylogger.EvKey:
			if !ev.KeyPress() {
				break
			}

			ks.log = append(ks.log, time.Now().Unix())
		}
	}

	return 0
}

func (ks *KeySpeed) watcher() {
	for {
		select {
		case <-time.After(time.Second * time.Duration(ks.Interval)):
			threshold := time.Now().Unix() - 60

			var i int
			for i = 0; i < len(ks.log); i++ {
				if ks.log[i] > threshold {
					break
				}
			}

			var best string
			ks.log = ks.log[i:]
			count := len(ks.log)
			pad := strconv.Itoa(ks.Pad)
			key := "CPM"

			if !ks.Cpm {
				count = count / WordKeys
				key = "WPM"
			}

			if ks.best < count {
				ks.best = count
			}

			if ks.Best {
				best = fmt.Sprintf("/%0"+pad+"d", ks.best)
			}

			fmt.Printf("%s: %0"+pad+"d%s\n", key, count, best)
		}
	}
}

func main() {
	app := gli.NewApplication(&KeySpeed{}, "Typing speed for i3blocks")

	app.Run()
}
