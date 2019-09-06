package Main

import (
	"./src"
	"strconv"
)

func main() {
	print("test")
	var node1 = createNode(5000, "10.0.0.3")
	var node2 = createNode(5000, "10.0.0.4")

	node1.CalcDistance(node2.ID)
	print("1")

}

func createNode(port int, ip string) src.Contact {
	id := src.NewRandomKademliaID()
	address := ip + ":" + strconv.Itoa(port)
	me := src.NewContact(id, address)
	return me
}
