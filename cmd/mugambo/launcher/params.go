package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://a0caa4808e578eb4e0fbaaffa012d76f39b6b8a4cfaf7e5ec7e8170cbff0e5a69d2f7872cdcbb4a7db077fbca5b91c2ab0b237344c04c7d47a9bda98994156de@65.108.15.158:30301",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
