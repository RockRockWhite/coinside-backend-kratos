package client

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewZkDiscovery, NewUserClinet, NewVoteClinet, NewCardClinet, NewTeamClinet, NewMarkdownClinet, NewTodoClinet)
