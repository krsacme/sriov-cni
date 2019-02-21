package config

import (
	"encoding/json"
	"fmt"

	sriovtypes "github.com/intel/sriov-cni/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
)

// Code generated by mockery v1.0.0. DO NOT EDIT.
// MockNetlinkManager is an autogenerated mock type for the NetlinkManager type
type MockNetlinkManager struct {
	mock.Mock
}

// LinkByName provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkByName(_a0 string) (netlink.Link, error) {
	ret := _m.Called(_a0)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeLink is a dummy netlink struct used during testing
type FakeLink struct {
	netlink.Link
}

var _ = Describe("Config", func() {
	Context("Checking LoadConf function", func() {
		It("Assuming correct config file", func() {
			conf := []byte(`{
	"name": "mynet",
	"type": "sriov",
	"master": "enp175s0f1",
	"mac":"66:77:88:99:aa:bb",
	"vf": 0,
	"ipam": {
	    "type": "host-local",
	    "subnet": "10.55.206.0/26",
	    "routes": [
	        { "dst": "0.0.0.0/0" }
	    ],
	    "gateway": "10.55.206.1"
	}
			}`)
			_, err := LoadConf(conf)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Assuming correct config file - existing DeviceID", func() {
			conf := []byte(`{
        "name": "mynet",
        "type": "sriov",
        "master": "enp175s0f1",
        "deviceID": "0000:af:06.1",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			_, err := LoadConf(conf)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Assuming incorrect config file - not existing DeviceID", func() {
			conf := []byte(`{
        "name": "mynet",
        "type": "sriov",
        "master": "enp175s0f1",
        "deviceID": "0000:af:06.3",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			_, err := LoadConf(conf)
			Expect(err).To(HaveOccurred())
		})
		It("Assuming incorrect config file - broken json", func() {
			conf := []byte(`{
        "name": "mynet"
        "type": "sriov",
        "master": "enp175s0f1",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			_, err := LoadConf(conf)
			Expect(err).To(HaveOccurred())
		})
		It("Assuming incorrect config file - missing master", func() {
			conf := []byte(`{
        "name": "mynet",
        "type": "sriov",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			_, err := LoadConf(conf)
			Expect(err).Should(MatchError("error: SRIOV-CNI loadConf: VF pci addr OR Master name is required"))
		})
	})
	Context("Checking getVfInfo function", func() {
		It("Assuming existing PF", func() {
			_, err := getVfInfo("0000:af:06.0")
			Expect(err).NotTo(HaveOccurred())
		})
		It("Assuming not existing PF", func() {
			_, err := getVfInfo("0000:af:07.0")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("Checking AssignFreeVF function", func() {
		mocked := &MockNetlinkManager{}
		fakeLink := FakeLink{}
		It("Assuming existing interface", func() {
			conf := []byte(`{
        "name": "mynet",
        "type": "sriov",
        "master": "enp175s0f1",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			var netconf sriovtypes.NetConf
			json.Unmarshal(conf, &netconf)
			mocked.On("LinkByName", mock.AnythingOfType("string")).Return(fakeLink, nil)
			nLink = mocked
			err := AssignFreeVF(&netconf)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Assuming not existing interface", func() {
			conf := []byte(`{
        "name": "mynet",
        "type": "sriov",
        "master": "enp175s0f2",
        "mac":"66:77:88:99:aa:bb",
        "vf": 0,
        "ipam": {
            "type": "host-local",
            "subnet": "10.55.206.0/26",
            "routes": [
                { "dst": "0.0.0.0/0" }
            ],
            "gateway": "10.55.206.1"
        }
                        }`)
			var netconf sriovtypes.NetConf
			json.Unmarshal(conf, &netconf)
			mocked.On("LinkByName", mock.AnythingOfType("string")).Return(nil, fmt.Errorf("No such interface"))
			nLink = mocked
			err := AssignFreeVF(&netconf)
			Expect(err).To(HaveOccurred())
		})
	})
})
