package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet"
)

// serverbound
const PacketIdSynchronizePlayerPosition = 0x40

const (
	SyncPosRelX = 1 << iota
	SyncPosRelY
	SyncPosRelZ
	SyncPosRelPitch
	SyncPosRelYaw
)

type SynchronizePlayerPosition struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	Flags      int8
	TeleportID int32
}

func (SynchronizePlayerPosition) ID() int32 {
	return PacketIdSynchronizePlayerPosition
}

func (s *SynchronizePlayerPosition) Encode(w io.Writer) error {
	if err := w.Double(s.X); err != nil {
		return err
	}
	if err := w.Double(s.Y); err != nil {
		return err
	}
	if err := w.Double(s.Z); err != nil {
		return err
	}
	if err := w.Float(s.Yaw); err != nil {
		return err
	}
	if err := w.Float(s.Pitch); err != nil {
		return err
	}
	if err := w.Byte(s.Flags); err != nil {
		return err
	}
	return w.VarInt(s.TeleportID)
}

func (s *SynchronizePlayerPosition) Decode(r io.Reader) error {
	if err := r.Double(&s.X); err != nil {
		return err
	}
	if err := r.Double(&s.Y); err != nil {
		return err
	}
	if err := r.Double(&s.Z); err != nil {
		return err
	}
	if err := r.Float(&s.Yaw); err != nil {
		return err
	}
	if err := r.Float(&s.Pitch); err != nil {
		return err
	}
	if err := r.Byte(&s.Flags); err != nil {
		return err
	}
	_, err := r.VarInt(&s.TeleportID)
	return err
}

// serverbound
const PacketIdConfirmTeleporation = 0x00

type ConfirmTeleporation struct {
	packet.EmptyPacket
}

func (ConfirmTeleporation) ID() int32 {
	return PacketIdConfirmTeleporation
}
