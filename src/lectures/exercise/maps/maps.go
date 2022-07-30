//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printStatus(serv map[string]int) {
	fmt.Println("Total servers", len(serv))
	status := make(map[int]int)

	for _, k := range serv {
		if _, ok := status[k]; ok {
			status[k] += 1
		} else {
			status[k] = 1
		}
	}

	for k, v := range status {
		switch k {
		case Online:
			fmt.Println(v, "\t servers online")
		case Offline:
			fmt.Println(v, "\t servers offline")
		case Maintenance:
			fmt.Println(v, "\t servers are in maintainance")
		case Retired:
			fmt.Println(v, "\t servers have been retired")
		default:
			panic("Unhandled server status")
		}
	}

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)
	for _, s := range servers {
		serverStatus[s] = Online
	}

	printStatus(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	printStatus(serverStatus)

	for k, _ := range serverStatus {
		serverStatus[k] = Maintenance
	}
	printStatus(serverStatus)
}
