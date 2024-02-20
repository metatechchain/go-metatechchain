package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://9c55681bbf99e7465d1bfb6625c0df6593c52eaf1641ad601b78524db0b65614c107050318981056e596901605a19d1930d653a50a644a0808bba1230e154d7c@185.185.83.124:5050",
		"enode://b619c596994998424854161b268e3e9b637befbdc35412ad8dbf6af76beafe109b3591949121a56b9bc9ef67906264971ed134ab112f587f7baae460fa7c6d2b@84.46.244.240:5050",
		"enode://e0e8f42b43e7c308058d80a59463afb7635127c5b9502ec7e68d1a7528a54eb6950cb4d4ef87a33c22f0cc43a57862d4a7133e4d427acf8f6042ac9da2dbccce@62.171.163.222:5050",
		"enode://03181e2f0ff8623c52bd774570e7f727c23c2a02b812b73b3100ed61ee1c6ff1d633d2c9c25ae9ac2142a25e645fa2ddfe068872361c8261319ca2819a444cc1@89.117.56.240:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
