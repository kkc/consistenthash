package consistent

import (
	"errors"
	//"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

// implement sort.Interface
func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// Ring is a network of distributed nodes.
type Ring struct {
	Nodes            Nodes
	NumberOfReplicas int
}

// Nodes is an array of nodes.
type Nodes []*Node

// Node is a single entity in a ring.
type Node struct {
	Id     string
	Num    int
	HashId uint32
}

// Initializes new distribute network of nodes or a ring.
func NewRing() *Ring {
	return &Ring{
		Nodes:            Nodes{},
		NumberOfReplicas: 5,
	}
}

func NewNode(id string, replicaNumber int) *Node {
	return &Node{
		Id:     id,
		Num:    replicaNumber,
		HashId: crc32.ChecksumIEEE([]byte(id + strconv.Itoa(replicaNumber))),
	}
}

func (r *Ring) AddNode(id string) {
	for i := 0; i < r.NumberOfReplicas; i++ {
		node := NewNode(id, i)
		//fmt.Printf("%s => %s\n", node.Id, node.HashId)
		r.Nodes = append(r.Nodes, node)
		sort.Sort(r.Nodes)
	}
}

func (r *Ring) Get(key string) string {
	// fmt.Println(key, crc32.ChecksumIEEE([]byte(key)))
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.ChecksumIEEE([]byte(key))
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}

func (r *Ring) RemoveNode(id string) error {
	for replica_idx := 0; replica_idx < r.NumberOfReplicas; replica_idx++ {
		searchfn := func(i int) bool {
			return r.Nodes[i].HashId >= crc32.ChecksumIEEE([]byte(id+strconv.Itoa(replica_idx)))
		}
		i := sort.Search(r.Nodes.Len(), searchfn)
		if i >= r.Nodes.Len() || r.Nodes[i].Id != id {
			return errors.New("node not found")
		}

		r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)
	}
	return nil
}
