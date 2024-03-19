package main

import (
	"flag"
	"sync"

	"github.com/steamgjk/paxi"
	"github.com/steamgjk/paxi/abd"
	"github.com/steamgjk/paxi/blockchain"
	"github.com/steamgjk/paxi/chain"
	"github.com/steamgjk/paxi/dynamo"
	"github.com/steamgjk/paxi/epaxos"
	"github.com/steamgjk/paxi/hpaxos"
	"github.com/steamgjk/paxi/kpaxos"
	"github.com/steamgjk/paxi/log"
	"github.com/steamgjk/paxi/m2paxos"
	"github.com/steamgjk/paxi/paxos"
	"github.com/steamgjk/paxi/paxos_group"
	"github.com/steamgjk/paxi/sdpaxos"
	"github.com/steamgjk/paxi/vpaxos"
	"github.com/steamgjk/paxi/wankeeper"
	"github.com/steamgjk/paxi/wpaxos"
)

var algorithm = flag.String("algorithm", "paxos", "Distributed algorithm")
var id = flag.String("id", "", "ID in format of Zone.Node.")
var simulation = flag.Bool("sim", false, "simulation mode")

var master = flag.String("master", "", "Master address.")

func replica(id paxi.ID) {
	if *master != "" {
		paxi.ConnectToMaster(*master, false, id)
	}

	log.Infof("node %v starting...", id)

	switch *algorithm {

	case "paxos":
		paxos.NewReplica(id).Run()

	case "epaxos":
		epaxos.NewReplica(id).Run()

	case "sdpaxos":
		sdpaxos.NewReplica(id).Run()

	case "wpaxos":
		wpaxos.NewReplica(id).Run()

	case "abd":
		abd.NewReplica(id).Run()

	case "chain":
		chain.NewReplica(id).Run()

	case "vpaxos":
		vpaxos.NewReplica(id).Run()

	case "wankeeper":
		wankeeper.NewReplica(id).Run()

	case "kpaxos":
		kpaxos.NewReplica(id).Run()

	case "paxos_groups":
		paxos_group.NewReplica(id).Run()

	case "dynamo":
		dynamo.NewReplica(id).Run()

	case "blockchain":
		blockchain.NewMiner(id).Run()

	case "m2paxos":
		m2paxos.NewReplica(id).Run()

	case "hpaxos":
		hpaxos.NewReplica(id).Run()

	default:
		panic("Unknown algorithm")
	}
}

func main() {
	log.Infof("Paxi Start Init")
	paxi.Init()
	log.Infof("Paxi Init Complete")
	if *simulation {
		var wg sync.WaitGroup
		wg.Add(1)
		paxi.Simulation()
		for id := range paxi.GetConfig().Addrs {
			n := id
			go replica(n)
		}
		wg.Wait()
	} else {
		replica(paxi.ID(*id))
	}
}
