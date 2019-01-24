package auth

import (
	"context"
	"google.golang.org/grpc/metadata"
)
const tokenHeader = "x-callback-token"

func (s *Auth_Server) extractToken(ctx context.Context) (tkn string, ok bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[tokenHeader]) == 0 {
		return "", false
	}

	return md[tokenHeader][0], true
}
