package auth

import (
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func (a *Auth_Server) sendBroadcasts(stream AuthService_WaitCallbackServer, tkn string) {

	streamClient := a.openStream(tkn)
	defer a.closeStream(tkn)

	for {
		select {
		case <-stream.Context().Done():
			return
		case res := <-streamClient:
			if s, ok := status.FromError(stream.Send(&res)); ok {
				switch s.Code() {
				case codes.OK:
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					return
				default:
					return
				}
			}
		}
	}
}

func (c *Auth_Server) MakeBroadcast(ctx context.Context) {
	for res := range c.Broadcast {
		c.streamsMtx.RLock()
		for _, stream := range c.ClientStreams {
			select {
			case stream <- res:
			default:
			}
		}
		c.streamsMtx.RUnlock()
	}
}

func (c *Auth_Server) openStream(tkn string) (stream chan CallbackData) {

	stream = make(chan CallbackData, 100)
	c.streamsMtx.Lock()
	c.ClientStreams[tkn] = stream
	c.streamsMtx.Unlock()

	return
}

func (c *Auth_Server) closeStream(tkn string) {

	c.streamsMtx.Lock()
	if stream, ok := c.ClientStreams[tkn]; ok {
		delete(c.ClientStreams, tkn)
		close(stream)
	}

	c.streamsMtx.Unlock()
}
