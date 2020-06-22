package main

import (
	"fmt"
	"time"

	"github.com/KofClubs/soccerdash"
	"github.com/go-co-op/gocron"
)

func main() {
	r := soccerdash.Reporter{
		Id:            "node_1",
		TargetAddress: "0.0.0.0:2020",
	}
	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Seconds().Do(func() {
		r.Report("TxPoolNum", "1", false)
		fmt.Println("TxPoolNum")
	})
	s.Every(10).Seconds().Do(func() {
		r.Report("ConnNum", "10", false)
		r.Report("IsProducer", false, false)
		fmt.Println("ConnNum")
		fmt.Println("IsProducer")
	})
	/*
		s.Every(30).Seconds().Do(func() {
			r.Report("NodeDelay", "30", false)
			fmt.Println("NodeDelay")
		})
	*/
	s.Every(1).Minute().Do(func() {
		r.Report("Version", "1.0", false)
		r.Report("NodeName", "123", false)
		fmt.Println("Version")
		fmt.Println("NodeName")
	})
	s.Every(1).Seconds().Do(func() {
		SeqMsg := `{
					"Hash": "0x48e44af30be97556e377b58de3e5d77e6801a82de906e647e121fe1497effac2",
					"Height": 2157,
					"Timestamp": 3141592653590
					}`
		r.Report("LatestSequencer", SeqMsg, false)
		fmt.Println("LatestSequencer")
	})
	s.StartBlocking()
}
