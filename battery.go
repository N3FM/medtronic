package medtronic

const (
	Battery CommandCode = 0x72
)

type BatteryInfo struct {
	MilliVolts int
	LowBattery bool
}

func (pump *Pump) Battery() BatteryInfo {
	result := pump.Execute(Battery, func(data []byte) interface{} {
		if len(data) < 4 || data[0] != 3 {
			return nil
		}
		return BatteryInfo{
			LowBattery: data[1] != 0,
			MilliVolts: twoByteInt(data[2:4]) * 10,
		}
	})
	if pump.Error() != nil {
		return BatteryInfo{}
	}
	return result.(BatteryInfo)
}
