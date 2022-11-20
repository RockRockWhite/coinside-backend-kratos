package client

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewZkDiscovery, NewUserClinet, NewCardClinet, NewTeamClinet, NewMarkdownClinet)
