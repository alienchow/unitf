package client

import (
	"context"
	"errors"
)

func (c *Client) ListCameras(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListNvr(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListSensors(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListLights(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListSpeakers(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListLiveviews(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListRelays(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListSirens(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListChimes(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListViewers(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListFobs(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListUsers(ctx context.Context) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}
