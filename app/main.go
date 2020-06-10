package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/latifrons/soccerdash"
)

func main() {
	r := soccerdash.Reporter{
		Id:            "node_1",
		TargetAddress: "127.0.0.1:8080",
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Seconds().Do(func() {
		r.Report("TxPoolNum", "1", false)
	})

	s.Every(10).Seconds().Do(func() {
		r.Report("ConnNum", "10", false)
		r.Report("IsProducer", false, false)
	})

	s.Every(30).Seconds().Do(func() {
		r.Report("NodeDelay", "30", false)
	})

	s.Every(1).Minute().Do(func() {
		r.Report("Version", "1.0", false)
		r.Report("NodeName", "123", false)
	})

	s.Every(1).Minute().Do(func() {
		SeqMsg := `{
			data: {
			Type: 1,
			Hash: "0x48e44af30be97556e377b58de3e5d77e6801a82de906e647e121fe1497effac2",
			ParentsHash: [
			"0x0b82bab454743a0f0cae194b15c8bfa049cfc4fa230b1bb0fd42e4790ba39a69"
			],
			AccountNonce: 2157,
			Height: 2157,
			PublicKey: "BIDG6ARHwZspJ1gUJTuHrjre0pyL7YJvG7E+UnbujaOVym+AjVi6jZIhptr92lwqi0xkXyXnI10gHt8RtnRHzrI=",
			Signature: "0x823d16432ba343f3a79d04ebadbab1683a5b93074d84e3c367c4815e568c682b156bff6c1b0a574319651dbe2d3e21f3ce21b5ea3085ea7de290371fb4b1297b01",
			MineNonce: 0,
			Weight: 2157,
			Version: 0,
			Issuer: "0x7349f7a6f622378d5fb0e2c16b9d4a3e5237c187",
			BlsJointSig: "0x",
			BlsJointPubKey: "0x",
			StateRoot: "0x980e8a14f85aaa1dfc7a4377f2d7ef7feaada75b08ab8b8e73590304bee548dc",
			Proposing: false,
			Timestamp: 1591769404219
			},
			err: ""
		}`
		r.Report("LatestSequencer", SeqMsg, false)
	})

	s.StartBlocking()
}
