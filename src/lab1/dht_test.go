package dht

/** go test -test.run TestDHT1 */

import (
	"fmt"
	"testing"
)

func TestDHT1(t *testing.T) {
	id0 := "00"
	id1 := "01"
	id2 := "02"
	id3 := "03"
	id4 := "04"
	id5 := "05"
	id6 := "06"
	id7 := "07"

	node0b := makeDHTNode(&id0, "localhost", "1111")
	node1b := makeDHTNode(&id1, "localhost", "1112")
	node2b := makeDHTNode(&id2, "localhost", "1113")
	node3b := makeDHTNode(&id3, "localhost", "1114")
	node4b := makeDHTNode(&id4, "localhost", "1115")
	node5b := makeDHTNode(&id5, "localhost", "1116")
	node6b := makeDHTNode(&id6, "localhost", "1117")
	node7b := makeDHTNode(&id7, "localhost", "1118")

	node0b.addToRing(node1b)
	node1b.addToRing(node2b)
	node1b.addToRing(node3b)
	node1b.addToRing(node4b)
	node4b.addToRing(node5b)
	node3b.addToRing(node6b)
	node3b.addToRing(node7b)

	fmt.Println("-> ring structure")
	node1b.printRing()
/**
	node3b.testCalcFingers(0, 3)
	node3b.testCalcFingers(1, 3)
	node3b.testCalcFingers(2, 3)
	node3b.testCalcFingers(3, 3)
	*/
}

func TestDHT2(t *testing.T) {
	node1 := makeDHTNode(nil, "localhost", "1111")
	node2 := makeDHTNode(nil, "localhost", "1112")
	node3 := makeDHTNode(nil, "localhost", "1113")
	node4 := makeDHTNode(nil, "localhost", "1114")
	node5 := makeDHTNode(nil, "localhost", "1115")
	node6 := makeDHTNode(nil, "localhost", "1116")
	node7 := makeDHTNode(nil, "localhost", "1117")
	node8 := makeDHTNode(nil, "localhost", "1118")
	node9 := makeDHTNode(nil, "localhost", "1119")

	key1 := "2b230fe12d1c9c60a8e489d028417ac89de57635"
	key2 := "87adb987ebbd55db2c5309fd4b23203450ab0083"
	key3 := "74475501523a71c34f945ae4e87d571c2c57f6f3"


	node1.addToRing(node2)
	node1.addToRing(node3)
	node1.addToRing(node4)
	node4.addToRing(node5)
	node3.addToRing(node6)
	node3.addToRing(node7)
	node3.addToRing(node8)
	node7.addToRing(node9)

	fmt.Println("TEST: " + node1.lookup(key1).nodeId + " is responsible for " + key1)
	fmt.Println("TEST: " + node1.lookup(key2).nodeId + " is responsible for " + key2)
	fmt.Println("TEST: " + node1.lookup(key3).nodeId + " is responsible for " + key3)


	fmt.Println("-> ring structure")
	node1.printRing()

	nodeForKey1 := node1.lookup(key1)
	fmt.Println("dht node " + nodeForKey1.nodeId + " running at " + nodeForKey1.contact.ip + ":" + nodeForKey1.contact.port + " is responsible for " + key1)

	nodeForKey2 := node1.lookup(key2)
	fmt.Println("dht node " + nodeForKey2.nodeId + " running at " + nodeForKey2.contact.ip + ":" + nodeForKey2.contact.port + " is responsible for " + key2)

	nodeForKey3 := node1.lookup(key3)
	fmt.Println("dht node " + nodeForKey3.nodeId + " running at " + nodeForKey3.contact.ip + ":" + nodeForKey3.contact.port + " is responsible for " + key3)
}


func TestDHT0(t *testing.T) {
        id0 := "00"
        id1 := "01"
        id2 := "02"
        id3 := "03"
        id4 := "04"
        id5 := "05"
        id6 := "06"
        id7 := "07"

        node0 := makeDHTNode(&id0, "localhost", "1111")
        node1 := makeDHTNode(&id1, "localhost", "1112")
        node2 := makeDHTNode(&id2, "localhost", "1113")
        node3 := makeDHTNode(&id3, "localhost", "1114")
        node4 := makeDHTNode(&id4, "localhost", "1115")
        node5 := makeDHTNode(&id5, "localhost", "1116")
        node6 := makeDHTNode(&id6, "localhost", "1117")
        node7 := makeDHTNode(&id7, "localhost", "1118")


	node0.addToRing(node1)
	node1.addToRing(node2)
	node1.addToRing(node3)
	node1.addToRing(node4)
	node4.addToRing(node5)
	node3.addToRing(node6)
	node3.addToRing(node7)


/**
	fmt.Println("node 3 lookups 04, should be 04")
	fmt.Println(node3.lookup("04").nodeId)
	fmt.Println("node 5 lookups 02, should be 03")
        fmt.Println(node5.lookup("02").nodeId)
	fmt.Println("node 3 lookups 02, should be 03")
        fmt.Println(node3.lookup("02").nodeId)
	fmt.Println("node 3 lookups 01, should be 01")
        fmt.Println(node3.lookup("01").nodeId)
*/

	fmt.Println("-> ring structure")
       	node1.printRing()
//	fmt.Println("-> ring structure")
//	node4.printRing()


       	node0.printRingFingers()

       	fmt.Println(node1.acceleratedLookupUsingFingers("05").nodeId)
       	//fmt.Println(node3.acceleratedLookupUsingFingers("02").nodeId)

	fmt.Println("node 3 lookups 09")
        fmt.Println(node3.lookup("09").nodeId)
	fmt.Println("node 3 lookups 10")
        fmt.Println(node3.lookup("10").nodeId)

}

func TestDHT4(t *testing.T) {
        id1 := "01"
        id8 := "08"
        id32 := "32"
        id67 := "67"
        id72 := "72"
        id82 := "82"
        id86 := "86"
        id87 := "87"

        node1 := makeDHTNode(&id1, "localhost", "1111")
        node8 := makeDHTNode(&id8, "localhost", "1112")
        node32 := makeDHTNode(&id32, "localhost", "1113")
        node67 := makeDHTNode(&id67, "localhost", "1114")
        node72 := makeDHTNode(&id72, "localhost", "1115")
        node82 := makeDHTNode(&id82, "localhost", "1116")
        node86 := makeDHTNode(&id86, "localhost", "1117")
        node87 := makeDHTNode(&id87, "localhost", "1118")


	node87.addToRing(node1)
	node87.addToRing(node8)
	node8.addToRing(node32)
	node8.addToRing(node67)
	node67.addToRing(node72)
	node32.addToRing(node82)
	node1.addToRing(node86)

	fmt.Println("-> ring structure")
       	node1.printRing()


       	node1.printRingFingers()
        fmt.Println(node1.lookup("210").nodeId)

       	//fmt.Println(node1.acceleratedLookupUsingFingers("05").nodeId)

}

