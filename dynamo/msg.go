package dynamo

import (
	"encoding/gob"

	"github.com/steamgjk/paxi"
)

func init() {
	gob.Register(Replicate{})
}

type Replicate struct {
	Command paxi.Command
}
