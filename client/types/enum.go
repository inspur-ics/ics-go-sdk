package types

import "reflect"

type VirtualMachinePowerState string

const (
	VirtualMachinePowerStatePoweredOff = VirtualMachinePowerState("poweredOff")
	VirtualMachinePowerStatePoweredOn  = VirtualMachinePowerState("poweredOn")
	VirtualMachinePowerStateSuspended  = VirtualMachinePowerState("suspended")
)

func init() {
	t["VirtualMachinePowerState"] = reflect.TypeOf((*VirtualMachinePowerState)(nil)).Elem()
}