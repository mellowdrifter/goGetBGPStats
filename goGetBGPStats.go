package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// A struct to hold the AS information all together.
type BGPStat struct {
	time             int
	v4Count          uint32
	v6Count          uint32
	peersConfigured  uint8
	peers6Configured uint8
	peersUp          uint8
	peers6Up         uint8
	v4Total          uint32
	v6Total          uint32
}

// Function attatched to the struct that will print all this information out
// String() actually overrides, to allow you to print a string representation of the struct
func (b BGPStat) String() string {
	return fmt.Sprintf("Latest info is:\nTime = %d\nV4 RIB = %d\nV4 FIB = %d\nV6 RIB = %d\nV6 FIB = %d\nV4 configured = %d\nV4 up = %d\nV6 configured = %d\nV6 up = %d",
		b.time, b.v4Total, b.v4Count, b.v6Total,
		b.v6Count, b.peersConfigured, b.peersUp,
		b.peers6Configured, b.peers6Up)
}

// Strut function
func (b BGPStat) AllUp() bool {
	if b.peersConfigured != b.peersUp && b.peers6Configured != b.peers6Up {
		return false
	}
	return true
}

func main() {

	// connect to the local database
	db, err := sql.Open("mysql", "bgpreader:k3hj4b5iu2yb@/doc_bgp_statistics")
	if err != nil {
		panic(err.Error())
	}

	// We will close the database connection at the end
	defer db.Close()

	// Ensure we can reach the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// create an empty struct
	bgpInfo := BGPStat{}

	// grab the required information and stick it directly into our struct
	err = db.QueryRow(`select TIME, V4COUNT, V6COUNT, V4TOTAL, V6TOTAL, PEERS_CONFIGURED,
		PEERS6_CONFIGURED, PEERS_UP, PEERS6_UP 
		from INFO ORDER by TIME DESC limit 1`).Scan(
		&bgpInfo.time,
		&bgpInfo.v4Count,
		&bgpInfo.v6Count,
		&bgpInfo.v4Total,
		&bgpInfo.v6Total,
		&bgpInfo.peersConfigured,
		&bgpInfo.peers6Configured,
		&bgpInfo.peersUp,
		&bgpInfo.peers6Up,
	)
	if err != nil {
		panic(err.Error())
	}

	// print out what we have
	fmt.Println(bgpInfo)

	switch bgpInfo.AllUp() {
	case true:
		fmt.Println("All configured peers are up")
	case false:
		fmt.Println("Not all configured peers are up")
	}

}
