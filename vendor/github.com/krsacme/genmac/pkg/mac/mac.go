package mac

import (
	"math/rand"
	"net"
	"time"
)

var DefaultBaseMacString = []string{
	"f2:3b:21:00:00:00",
	"f6:3b:21:00:00:00",
	"fa:3b:21:00:00:00",
	"fe:3b:21:00:00:00",
}

var baseMacList []net.HardwareAddr

// Configure the list of mac addresses to be used as base mac address to generate one
// Entries should be a valid mac address in the locally administered address range
// (x2:xx:xx:xx:xx:xx, x6:xx:xx:xx:xx:xx, xa:xx:xx:xx:xx:xx, xe:xx:xx:xx:xx:xx)
// Entries should have trailing 0's (like x6:xx:xx:00:00:00)
// Values with 0's will be generated with random numbers
// Uniqueness of the genaretaed mac address is guranteed by storing the generated values cluster wide
func ConfigureBaseMacRange(baseMacs []string) error {
	baseMacList = nil
	for _, baseMac := range baseMacs {
		hw, err := net.ParseMAC(baseMac)
		if err != nil {
			return err
		}

		// TODO: Add validation for locally administered addresss

		baseMacList = append(baseMacList, hw)
	}
	return nil
}

func GenerateMacAddress() (net.HardwareAddr, error) {
	// Set Default value if base mac list is empty
	if len(baseMacList) == 0 {
		err := ConfigureBaseMacRange(DefaultBaseMacString)
		if err != nil {
			return nil, err
		}
	}
	newMac := generateMac(baseMacList)
	return newMac, nil
}

func generateMac(baseMacs []net.HardwareAddr) net.HardwareAddr {
	// Choose a base mac address from the configured list
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(baseMacs))
	baseMac := baseMacs[idx]

	// Find the number of entries with 00
	zeroIdx := 0
	for ; zeroIdx < len(baseMac); zeroIdx++ {
		rIdx := len(baseMac) - zeroIdx - 1
		if baseMac[rIdx] != 0 {
			break
		}
	}

	// Generate Random number for bytes with 0's
	randMac := make([]byte, len(baseMac)-zeroIdx)
	rand.Read(randMac)

	// Append fixed mac with generated mac
	mac := baseMac[:zeroIdx]
	mac = append(mac, randMac...)
	return mac
}
