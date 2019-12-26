package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAddAchievement represents the MSG_MHF_ADD_ACHIEVEMENT
type MsgMhfAddAchievement struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAddAchievement) Opcode() network.PacketID {
	return network.MSG_MHF_ADD_ACHIEVEMENT
}

// Parse parses the packet from binary
func (m *MsgMhfAddAchievement) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAddAchievement) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}