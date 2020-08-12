package proc

import (
	"fmt"
)

const (
	IMAGE_FILE_MACHINE_UNKNOWN   = 0x0
	IMAGE_FILE_MACHINE_AM33      = 0x1d3
	IMAGE_FILE_MACHINE_AMD64     = 0x8664
	IMAGE_FILE_MACHINE_ARM       = 0x1c0
	IMAGE_FILE_MACHINE_ARMNT     = 0x1c4
	IMAGE_FILE_MACHINE_ARM64     = 0xaa64
	IMAGE_FILE_MACHINE_EBC       = 0xebc
	IMAGE_FILE_MACHINE_I386      = 0x14c
	IMAGE_FILE_MACHINE_IA64      = 0x200
	IMAGE_FILE_MACHINE_M32R      = 0x9041
	IMAGE_FILE_MACHINE_MIPS16    = 0x266
	IMAGE_FILE_MACHINE_MIPSFPU   = 0x366
	IMAGE_FILE_MACHINE_MIPSFPU16 = 0x466
	IMAGE_FILE_MACHINE_POWERPC   = 0x1f0
	IMAGE_FILE_MACHINE_POWERPCFP = 0x1f1
	IMAGE_FILE_MACHINE_R4000     = 0x166
	IMAGE_FILE_MACHINE_SH3       = 0x1a2
	IMAGE_FILE_MACHINE_SH3DSP    = 0x1a3
	IMAGE_FILE_MACHINE_SH4       = 0x1a6
	IMAGE_FILE_MACHINE_SH5       = 0x1a8
	IMAGE_FILE_MACHINE_THUMB     = 0x1c2
	IMAGE_FILE_MACHINE_WCEMIPSV2 = 0x169
)

type PEMachine uint16

// PEMachineString map pe machine to name, See $GOROOT/src/debug/pe/pe.go for detail
var PEMachineString = map[uint16]string{
	IMAGE_FILE_MACHINE_UNKNOWN:   "unknown",
	IMAGE_FILE_MACHINE_AM33:      "am33",
	IMAGE_FILE_MACHINE_AMD64:     "amd64",
	IMAGE_FILE_MACHINE_ARM:       "arm",
	IMAGE_FILE_MACHINE_ARMNT:     "armnt",
	IMAGE_FILE_MACHINE_ARM64:     "arm64",
	IMAGE_FILE_MACHINE_EBC:       "ebc",
	IMAGE_FILE_MACHINE_I386:      "i386",
	IMAGE_FILE_MACHINE_IA64:      "ia64",
	IMAGE_FILE_MACHINE_M32R:      "m32r",
	IMAGE_FILE_MACHINE_MIPS16:    "mips16",
	IMAGE_FILE_MACHINE_MIPSFPU:   "mipsfpu",
	IMAGE_FILE_MACHINE_MIPSFPU16: "mipsfpu16",
	IMAGE_FILE_MACHINE_POWERPC:   "powerpc",
	IMAGE_FILE_MACHINE_POWERPCFP: "powerpcfp",
	IMAGE_FILE_MACHINE_R4000:     "r4000",
	IMAGE_FILE_MACHINE_SH3:       "sh3",
	IMAGE_FILE_MACHINE_SH3DSP:    "sh3dsp",
	IMAGE_FILE_MACHINE_SH4:       "sh4",
	IMAGE_FILE_MACHINE_SH5:       "sh5",
	IMAGE_FILE_MACHINE_THUMB:     "thumb",
	IMAGE_FILE_MACHINE_WCEMIPSV2: "wcemipsv2",
}

func (m PEMachine) String() string {
	str, ok := PEMachineString[uint16(m)]
	if ok {
		return str
	}
	return fmt.Sprintf("unkown image file machine code %d\n", uint16(m))
}