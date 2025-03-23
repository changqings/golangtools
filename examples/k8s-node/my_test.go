package examples

import (
	"fmt"
	"testing"

	"github.com/changqings/golangtools/slice"
)

type MyNode struct {
	PoolName string
	NGName   string
	Name     []string
}
type Node struct {
	PoolName string
	NgName   string
	Name     string
}

func TestK8sNode(t *testing.T) {

	NodeList := []Node{
		{"pool1", "ng1", "node_a1"},
		{"pool1", "ng1", "node_a2"},
		{"pool1", "ng2", "node_b1"},
		{"pool1", "ng2", "node_b2"},
		{"pool2", "ng21", "node_c1"},
		{"pool2", "ng21", "node_c2"},
		{"pool2", "ng22", ""},
	}

	var ngMap = make(map[string][]string)
	var poolMap = make(map[string][]string)
	var poolWithDepNgMap = make(map[string][]string)

	for _, n := range NodeList {
		ngMap[n.NgName] = append(ngMap[n.NgName], n.Name)
		poolMap[n.PoolName] = append(poolMap[n.PoolName], n.NgName)
	}

	fmt.Printf("ngMap=%v\n", ngMap)
	fmt.Printf("poolMap=%v\n", poolMap)

	for k, v := range poolMap {
		poolWithDepNgMap[k] = slice.Depulicate(v)
	}
	fmt.Printf("poolWithDepNgMap=%v\n", poolWithDepNgMap)

	MyNodes := []MyNode{}
	for k, v := range poolWithDepNgMap {
		for _, n := range v {
			MyNode := MyNode{
				PoolName: k,
				NGName:   n,
				Name:     ngMap[n],
			}
			MyNodes = append(MyNodes, MyNode)
		}
	}
	for _, v := range MyNodes {
		fmt.Printf("poolName=%s,ngName=%s,nodeNames=%v\n", v.PoolName, v.NGName, v.Name)
	}

}
