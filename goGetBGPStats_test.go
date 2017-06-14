package main

import (
	"testing"
)

func TestString(t *testing.T) {
	bgpValues := []struct {
		desc string
		data BGPStat
		want string
	}{
		{
			desc: "High values test",
			data: BGPStat{
				time:             148555999,
				v4Count:          650000,
				v6Count:          42000,
				peersConfigured:  6,
				peers6Configured: 6,
				peersUp:          5,
				peers6Up:         5,
				v4Total:          3500000,
				v6Total:          100000,
			},
			want: "Latest info is:\nTime = 148555999\nV4 RIB = 3500000\nV4 FIB = 650000\nV6 RIB = 100000\nV6 FIB = 42000\nV4 configured = 6\nV4 up = 5\nV6 configured = 6\nV6 up = 5",
		}, {
			desc: "Low values test",
			data: BGPStat{
				time:             500,
				v4Count:          111,
				v6Count:          11,
				peersConfigured:  16,
				peers6Configured: 16,
				peersUp:          15,
				peers6Up:         15,
				v4Total:          1111,
				v6Total:          111,
			},
			want: "Latest info is:\nTime = 500\nV4 RIB = 1111\nV4 FIB = 111\nV6 RIB = 111\nV6 FIB = 11\nV4 configured = 16\nV4 up = 15\nV6 configured = 16\nV6 up = 15",
		},
	}

	for _, v := range bgpValues {
		got := v.data.String()
		if got != v.want {
			t.Errorf("Test desc (%v) failed.\nGot = %v\n Want = %v", v.desc, got, v.want)
		}
	}
}

func TestAllUp(t *testing.T) {
	upValues := []struct {
		desc string
		data BGPStat
		want bool
	}{
		{
			desc: "True test",
			data: BGPStat{
				peersConfigured:  5,
				peers6Configured: 5,
				peersUp:          5,
				peers6Up:         5,
			},
			want: true,
		}, {
			desc: "False test",
			data: BGPStat{
				peersConfigured:  6,
				peers6Configured: 6,
				peersUp:          5,
				peers6Up:         5,
			},
			want: false,
		},
	}

	for _, v := range upValues {
		got := v.data.AllUp()
		if got != v.want {
			t.Errorf("Test desc(%v): Unexpected boolean received\n Got = %v\n Want = %v", v.desc, got, v.want)
		}
	}
}
