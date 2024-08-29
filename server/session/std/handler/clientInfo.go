package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdClientInformation, handleClientInfo)
	std.RegisterHandler(net.ConfigurationState, configuration.PacketIdClientInformation, handleClientInfo)
}

func handleClientInfo(s *std.StandardSession, p packet.Decodeable) {
	var inf configuration.ClientInformation
	switch pk := p.(type) {
	case *configuration.ClientInformation:
		inf = *pk
		s.ClientInfo.Set(*pk)
	case *play.ClientInformation:
		inf = pk.ClientInformation
		s.ClientInfo.Set(pk.ClientInformation)
	default:
		return
	}

	new := metadata.Metadata{
		metadata.PlayerDisplayedSkinPartsIndex: metadata.Byte(inf.DisplayedSkinParts),
		metadata.PlayerMainHandIndex:           metadata.Byte(inf.MainHand),
	}
	s.Player().SetMetadataIndexes(new)
	s.Broadcast().EntityMetadata(s, new)
}
