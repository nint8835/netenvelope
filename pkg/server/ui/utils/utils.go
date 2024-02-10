package utils

import (
	"context"

	"github.com/nint8835/netenvelope/pkg/database/queries"
)

func GetCurrentUser(ctx context.Context) *queries.User {
	return ctx.Value("current_user").(*queries.User)
}
