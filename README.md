# D7024E-DDS
This is an implementation of a kademlia network with a distributed hashtable  from the [Kademlia specification](http://xlattice.sourceforge.net/components/protocol/kademlia/specs.html). 

## How to run the program
1. Clone the project
2. Navigate into the project directory
3. Run `./restart.sh`
4. Answer `y` to the questions if you want to make sure the run is clean from previous runs
5. You have now started a distributed kademlia network

## How the program works
* In a terminal run `docker attach [node name]` to get into the node's cli 
* Inside the terminal one can type `h` to get a description of the commands one can use
* To change the number of nodes to run edit the x in `kademliaNodes=[x]` in the restart.sh script
* To edit the constants k or alpha edit the constans in main.go
