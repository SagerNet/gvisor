// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linux

const (
	// NumControlCharacters is the number of control characters in Termios.
	NumControlCharacters = 19
	// disabledChar is used to indicate that a control character is
	// disabled.
	disabledChar = 0
)

// Winsize is struct winsize, defined in uapi/asm-generic/termios.h.
//
// +marshal
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Termios is struct termios, defined in uapi/asm-generic/termbits.h.
//
// +marshal
type Termios struct {
	InputFlags        uint32
	OutputFlags       uint32
	ControlFlags      uint32
	LocalFlags        uint32
	LineDiscipline    uint8
	ControlCharacters [NumControlCharacters]uint8
}

// KernelTermios is struct ktermios/struct termios2, defined in
// uapi/asm-generic/termbits.h.
//
// +stateify savable
type KernelTermios struct {
	InputFlags        uint32
	OutputFlags       uint32
	ControlFlags      uint32
	LocalFlags        uint32
	LineDiscipline    uint8
	ControlCharacters [NumControlCharacters]uint8
	InputSpeed        uint32
	OutputSpeed       uint32
}

// IEnabled returns whether flag is enabled in termios input flags.
func (t *KernelTermios) IEnabled(flag uint32) bool {
	return t.InputFlags&flag == flag
}

// OEnabled returns whether flag is enabled in termios output flags.
func (t *KernelTermios) OEnabled(flag uint32) bool {
	return t.OutputFlags&flag == flag
}

// CEnabled returns whether flag is enabled in termios control flags.
func (t *KernelTermios) CEnabled(flag uint32) bool {
	return t.ControlFlags&flag == flag
}

// LEnabled returns whether flag is enabled in termios local flags.
func (t *KernelTermios) LEnabled(flag uint32) bool {
	return t.LocalFlags&flag == flag
}

// ToTermios copies fields that are shared with Termios into a new Termios
// struct.
func (t *KernelTermios) ToTermios() Termios {
	return Termios{
		InputFlags:        t.InputFlags,
		OutputFlags:       t.OutputFlags,
		ControlFlags:      t.ControlFlags,
		LocalFlags:        t.LocalFlags,
		LineDiscipline:    t.LineDiscipline,
		ControlCharacters: t.ControlCharacters,
	}
}

// FromTermios copies fields that are shared with Termios into this
// KernelTermios struct.
func (t *KernelTermios) FromTermios(term Termios) {
	t.InputFlags = term.InputFlags
	t.OutputFlags = term.OutputFlags
	t.ControlFlags = term.ControlFlags
	t.LocalFlags = term.LocalFlags
	t.LineDiscipline = term.LineDiscipline
	t.ControlCharacters = term.ControlCharacters
}

// IsTerminating returns whether c is a line terminating character.
func (t *KernelTermios) IsTerminating(cBytes []byte) bool {
	// All terminating characters are 1 byte.
	if len(cBytes) != 1 {
		return false
	}
	c := cBytes[0]

	// Is this the user-set EOF character?
	if t.IsEOF(c) {
		return true
	}

	switch c {
	case disabledChar:
		return false
	case '\n', t.ControlCharacters[VEOL]:
		return true
	case t.ControlCharacters[VEOL2]:
		return t.LEnabled(IEXTEN)
	}
	return false
}

// IsEOF returns whether c is the EOF character.
func (t *KernelTermios) IsEOF(c byte) bool {
	return c == t.ControlCharacters[VEOF] && t.ControlCharacters[VEOF] != disabledChar
}

// Input flags.
const (
	IGNBRK  = 0o000001
	BRKINT  = 0o000002
	IGNPAR  = 0o000004
	PARMRK  = 0o000010
	INPCK   = 0o000020
	ISTRIP  = 0o000040
	INLCR   = 0o000100
	IGNCR   = 0o000200
	ICRNL   = 0o000400
	IUCLC   = 0o001000
	IXON    = 0o002000
	IXANY   = 0o004000
	IXOFF   = 0o010000
	IMAXBEL = 0o020000
	IUTF8   = 0o040000
)

// Output flags.
const (
	OPOST  = 0o000001
	OLCUC  = 0o000002
	ONLCR  = 0o000004
	OCRNL  = 0o000010
	ONOCR  = 0o000020
	ONLRET = 0o000040
	OFILL  = 0o000100
	OFDEL  = 0o000200
	NLDLY  = 0o000400
	NL0    = 0o000000
	NL1    = 0o000400
	CRDLY  = 0o003000
	CR0    = 0o000000
	CR1    = 0o001000
	CR2    = 0o002000
	CR3    = 0o003000
	TABDLY = 0o014000
	TAB0   = 0o000000
	TAB1   = 0o004000
	TAB2   = 0o010000
	TAB3   = 0o014000
	XTABS  = 0o014000
	BSDLY  = 0o020000
	BS0    = 0o000000
	BS1    = 0o020000
	VTDLY  = 0o040000
	VT0    = 0o000000
	VT1    = 0o040000
	FFDLY  = 0o100000
	FF0    = 0o000000
	FF1    = 0o100000
)

// Control flags.
const (
	CBAUD    = 0o010017
	B0       = 0o000000
	B50      = 0o000001
	B75      = 0o000002
	B110     = 0o000003
	B134     = 0o000004
	B150     = 0o000005
	B200     = 0o000006
	B300     = 0o000007
	B600     = 0o000010
	B1200    = 0o000011
	B1800    = 0o000012
	B2400    = 0o000013
	B4800    = 0o000014
	B9600    = 0o000015
	B19200   = 0o000016
	B38400   = 0o000017
	EXTA     = B19200
	EXTB     = B38400
	CSIZE    = 0o000060
	CS5      = 0o000000
	CS6      = 0o000020
	CS7      = 0o000040
	CS8      = 0o000060
	CSTOPB   = 0o000100
	CREAD    = 0o000200
	PARENB   = 0o000400
	PARODD   = 0o001000
	HUPCL    = 0o002000
	CLOCAL   = 0o004000
	CBAUDEX  = 0o010000
	BOTHER   = 0o010000
	B57600   = 0o010001
	B115200  = 0o010002
	B230400  = 0o010003
	B460800  = 0o010004
	B500000  = 0o010005
	B576000  = 0o010006
	B921600  = 0o010007
	B1000000 = 0o010010
	B1152000 = 0o010011
	B1500000 = 0o010012
	B2000000 = 0o010013
	B2500000 = 0o010014
	B3000000 = 0o010015
	B3500000 = 0o010016
	B4000000 = 0o010017
	CIBAUD   = 0o02003600000
	CMSPAR   = 0o10000000000
	CRTSCTS  = 0o20000000000

	// IBSHIFT is the shift from CBAUD to CIBAUD.
	IBSHIFT = 16
)

// Local flags.
const (
	ISIG    = 0o000001
	ICANON  = 0o000002
	XCASE   = 0o000004
	ECHO    = 0o000010
	ECHOE   = 0o000020
	ECHOK   = 0o000040
	ECHONL  = 0o000100
	NOFLSH  = 0o000200
	TOSTOP  = 0o000400
	ECHOCTL = 0o001000
	ECHOPRT = 0o002000
	ECHOKE  = 0o004000
	FLUSHO  = 0o010000
	PENDIN  = 0o040000
	IEXTEN  = 0o100000
	EXTPROC = 0o200000
)

// Control Character indices.
const (
	VINTR    = 0
	VQUIT    = 1
	VERASE   = 2
	VKILL    = 3
	VEOF     = 4
	VTIME    = 5
	VMIN     = 6
	VSWTC    = 7
	VSTART   = 8
	VSTOP    = 9
	VSUSP    = 10
	VEOL     = 11
	VREPRINT = 12
	VDISCARD = 13
	VWERASE  = 14
	VLNEXT   = 15
	VEOL2    = 16
)

// ControlCharacter returns the termios-style control character for the passed
// character.
//
// e.g., for Ctrl-C, i.e., ^C, call ControlCharacter('C').
//
// Standard control characters are ASCII bytes 0 through 31.
func ControlCharacter(c byte) uint8 {
	// A is 1, B is 2, etc.
	return uint8(c - 'A' + 1)
}

// DefaultControlCharacters is the default set of Termios control characters.
var DefaultControlCharacters = [NumControlCharacters]uint8{
	ControlCharacter('C'),  // VINTR = ^C
	ControlCharacter('\\'), // VQUIT = ^\
	'\x7f',                 // VERASE = DEL
	ControlCharacter('U'),  // VKILL = ^U
	ControlCharacter('D'),  // VEOF = ^D
	0,                      // VTIME
	1,                      // VMIN
	0,                      // VSWTC
	ControlCharacter('Q'),  // VSTART = ^Q
	ControlCharacter('S'),  // VSTOP = ^S
	ControlCharacter('Z'),  // VSUSP = ^Z
	0,                      // VEOL
	ControlCharacter('R'),  // VREPRINT = ^R
	ControlCharacter('O'),  // VDISCARD = ^O
	ControlCharacter('W'),  // VWERASE = ^W
	ControlCharacter('V'),  // VLNEXT = ^V
	0,                      // VEOL2
}

// MasterTermios is the terminal configuration of the master end of a Unix98
// pseudoterminal.
var MasterTermios = KernelTermios{
	ControlFlags:      B38400 | CS8 | CREAD,
	ControlCharacters: DefaultControlCharacters,
	InputSpeed:        38400,
	OutputSpeed:       38400,
}

// DefaultReplicaTermios is the default terminal configuration of the replica
// end of a Unix98 pseudoterminal.
var DefaultReplicaTermios = KernelTermios{
	InputFlags:        ICRNL | IXON,
	OutputFlags:       OPOST | ONLCR,
	ControlFlags:      B38400 | CS8 | CREAD,
	LocalFlags:        ISIG | ICANON | ECHO | ECHOE | ECHOK | ECHOCTL | ECHOKE | IEXTEN,
	ControlCharacters: DefaultControlCharacters,
	InputSpeed:        38400,
	OutputSpeed:       38400,
}

// WindowSize corresponds to struct winsize defined in
// include/uapi/asm-generic/termios.h.
//
// +stateify savable
// +marshal
type WindowSize struct {
	Rows uint16
	Cols uint16
	_    [4]byte // Padding for 2 unused shorts.
}
