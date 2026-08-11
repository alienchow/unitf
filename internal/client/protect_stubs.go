package client

import (
	"context"
)

func (h *ProtectHandler) listIDs(ctx context.Context, path string) ([]struct{ ID string }, error) {
	var resp []struct{ ID string }
	err := h.Request(ctx, "GET", path, nil, &resp)
	return resp, err
}

func (h *ProtectHandler) ListCameras(ctx context.Context) ([]CameraDto, error) {
	var resp []CameraDto
	if err := h.Request(ctx, "GET", "/v1/cameras", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *ProtectHandler) ListNvr(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/nvrs")
}

func (h *ProtectHandler) ListSensors(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/sensors")
}

func (h *ProtectHandler) ListLights(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/lights")
}

func (h *ProtectHandler) ListSpeakers(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/speakers")
}

func (h *ProtectHandler) ListLiveviews(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/liveviews")
}

func (h *ProtectHandler) ListRelays(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/relays")
}

func (h *ProtectHandler) ListSirens(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/sirens")
}

func (h *ProtectHandler) ListChimes(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/chimes")
}

func (h *ProtectHandler) ListViewers(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/viewers")
}

func (h *ProtectHandler) ListFobs(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/fobs")
}

func (h *ProtectHandler) ListUsers(ctx context.Context) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/users")
}
