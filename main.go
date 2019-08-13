package main

import (
	"fmt"
	"log"

	"github.com/alessandromr/amloadt/sharedq"
	"github.com/alessandromr/amloadt/ualyzer"
)

func main() {

	list := sharedq.SharedList{}

	list.Push("Node 1")

	log.Println(list.DirectPop("Node 1"))

	url := ualyzer.URL{Raw: "lighsail.aws.amazon.com/test/index.html"}

	fmt.Println(url.GetDomain())
	fmt.Println(url.GetFQDN())

	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	n, err := list.Join([]string{"1.2.3.4"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}

}
