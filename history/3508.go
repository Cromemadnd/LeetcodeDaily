package main

import "github.com/emirpasic/gods/sets/treeset"

type Packet struct {
	Source, Destination, Timestamp int
}

type Router struct {
	PacketsHash map[Packet]struct{}
	Packets     *treeset.Set
	MemoryLimit int
	Counter     int
}

func PacketsComparator(a, b interface{}) int {
	p1 := a.(Packet)
	p2 := b.(Packet)
	if p1.Timestamp == p2.Timestamp {
		return p1.ID - p2.ID
	}
	return p1.Timestamp - p2.Timestamp
}

func Constructor(memoryLimit int) Router {
	return Router{
		PacketsHash: make(map[Packet]struct{},
		Packets:     treeset.NewWith(PacketsComparator),
		MemoryLimit: memoryLimit,
		Counter:     0,
	}
}

func (r *Router) popFirstPacket() *Packet {
	iter := r.Packets.Iterator()
	if iter.Next() {
		packet := iter.Value().(Packet)
		packetWithoutID := Packet{Source: packet.Source, Destination: packet.Destination, Timestamp: packet.Timestamp, ID: -1}
		r.Packets.Remove(packet)
		delete(r.PacketsHash, packetWithoutID)
		return &packet
	}
	return nil
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := Packet{Source: source, Destination: destination, Timestamp: timestamp, ID: -1}
	if r.PacketsHash[packet] {
		return false
	}
	r.PacketsHash[packet] = true

	packetWithID := Packet{Source: source, Destination: destination, Timestamp: timestamp, ID: r.Counter}
	r.Counter++
	r.Packets.Add(packetWithID)
	if r.Packets.Size() > r.MemoryLimit {
		r.popFirstPacket()
	}
	return true
}

func (r *Router) ForwardPacket() []int {
	if packet := r.popFirstPacket(); packet != nil {
		return []int{packet.Source, packet.Destination, packet.Timestamp}
	}
	return []int{}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	count := 0
	minPacket := Packet{Source: -1, Destination: -1, Timestamp: startTime, ID: -1}
	// r.Packets.Ceiling(minPacket)

	treeset.


	for iter := r.Packets.Iterator(); iter.Next(); {
		packet := iter.Value().(Packet)
		if packet.Timestamp < startTime {
			continue
		}
		if packet.Timestamp > endTime {
			break
		}
		if packet.Destination == destination {
			count++
		}
	}
	return count
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
