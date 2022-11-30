package controller

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewUserController, NewCardController, NewAttachmentController, NewVoteController, NewTeamController, NewMarkdownController, NewTodoController, NewObjectController)
