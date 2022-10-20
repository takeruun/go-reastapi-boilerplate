package service

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/wader/gormstore/v2"
)

type SessionService interface {
	GetSession(ctx context.Context, name string) (*sessions.Session, error)
	GetSessionValue(ctx context.Context, key string) (interface{}, error)
	SaveSession(ctx context.Context, key string, value interface{}) error
	DeleteSession(ctx context.Context) error
}

type sessionService struct {
	store *gormstore.Store
}

func NewSessionService(s *gormstore.Store) SessionService {
	return &sessionService{
		store: s,
	}
}

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

// GetSession returns a cached session of the given name
func (service *sessionService) GetSession(ctx context.Context, name string) (*sessions.Session, error) {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	// Ignore err because a session is always returned even if one doesn't exist
	session, err := service.store.Get(httpContext.R, name)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (service *sessionService) GetSessionValue(ctx context.Context, key string) (interface{}, error) {
	session, err := service.GetSession(ctx, "_goreset_session")
	if err != nil {
		return "", err
	}

	return session.Values[key], nil
}

func (service *sessionService) SaveSession(ctx context.Context, key string, value interface{}) error {
	session, err := service.GetSession(ctx, "_goreset_session")
	if err != nil {
		return err
	}

	session.Values[key] = value

	httpContext := ctx.Value(HTTPKey("http")).(HTTP)
	err = service.store.Save(httpContext.R, *httpContext.W, session)
	if err != nil {
		return err
	}

	return nil
}

func (service *sessionService) DeleteSession(ctx context.Context) error {
	session, err := service.GetSession(ctx, "_goreset_session")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)
	err = service.store.Save(httpContext.R, *httpContext.W, session)
	if err != nil {
		return err
	}

	return nil
}
