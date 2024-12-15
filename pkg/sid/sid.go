package sid

import (
	"net"

	"github.com/sony/sonyflake"
)

type Sid struct {
	sf *sonyflake.Sonyflake
}

func GetDeviceMacLow16Bit() (uint16, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}
	for _, nif := range netInterfaces {
		if len(nif.HardwareAddr) > 2 {
			addr := nif.HardwareAddr
			return uint16(addr[4])<<8 + uint16(addr[5]), nil
		}
	}
	return 0, nil
}

func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(
		sonyflake.Settings{
			MachineID: GetDeviceMacLow16Bit,
		},
	)
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}
func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", err
	}
	return IntToBase62(int(id)), nil
}
func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}
