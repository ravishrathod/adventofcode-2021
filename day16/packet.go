package main

type Packet struct {
	version    int
	typ        int
	value      int
	subPackets []*Packet
	parent	   *Packet
}
