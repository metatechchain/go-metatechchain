package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://53fef6d8ccd675cc997aaf1798cea0774f2b8fd32a74f40aef97f54066be9dde09620c569ec9595bbf474bad3f6e8c31c7abbeb2a4cc1e83cb662f404eb5651d@89.117.56.240:5050",
		"enode://4bec37fd321019986073f373b0a1c12163646e1bad989537fab4367173f47c87862ebf169826a13a38b67fa74e5310d841046bc4de3a975e43d0a144f17ae573@185.185.83.124:5050",
		"enode://f457e71fc54625fe90ad1b43b3ff7cc82805bbea749452729e6a53f91e10aeb0c3cb1e352207cec8c165ee12875b8407154bdb5107e0bb863edab1727b7e5230@62.171.163.222:5050",
		"enode://3dda45a6d424987394b6032d62d1bf1bcc9f741ab9bdecf60df401fc68ee7f3f3e4b55ec0bf4e599feb301c03071647bcd95fc285aced4e03be80f80db524936@84.46.244.240:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
