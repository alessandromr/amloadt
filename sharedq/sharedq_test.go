package sharedq

import "testing"

/*
	Test
*/
func TestPushPull(t *testing.T) {
	list := SharedList{}
	list.Push("Node")
	result := list.Pull()
	checkNodes(result, t)
}

func TestPushPop(t *testing.T) {
	list := SharedList{}
	list.Push("Node")
	result := list.Pop()
	checkNodes(result, t)
}

func TestDirectPop(t *testing.T) {
	list := SharedList{}
	list.Push("Node")
	result := list.DirectPop("Node")
	checkNodes(result, t)
}

func checkNodes(result Node, t *testing.T) {
	if result.value != "Node" {
		t.Errorf("Error pushing or pulling and delete")
	}
	if result.nextNode != nil {
		t.Errorf("Error with next Node during p&p")
	}
	if result.prevNode != nil {
		t.Errorf("Error with prev Node during p&p")
	}
}

/*
	Benchmark
*/

func BenchmarkPush(b *testing.B) {
	list := []SharedList{}
	for n := 0; n < b.N; n++ {
		list.Push("Node")
	}
}

func BenchmarkPull(b *testing.B) {
	list := SharedList{}
	list.Push("Node")
	for n := 0; n < b.N; n++ {
		list.Pull()
	}
}

func BenchmarkComplete(b *testing.B) {
	List := SharedList{}
	List.Push("Node")
	for n := 0; n < b.N; n++ {
		List.Push("Node")
		List.Pop()
	}
}

func BenchmarkGrowingPush(b *testing.B) {
	list := SharedList{}
	for n := 0; n < b.N; n++ {
		for n := 0; n < b.N/10; n++ {
			list.Push("Node")
		}
	}
}
