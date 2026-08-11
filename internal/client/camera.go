package client

import (
	"context"
)

type CameraDto struct {
	ID               string             `json:"id,omitempty"`
	Name             string             `json:"name,omitempty"`
	VideoMode        string             `json:"videoMode,omitempty"`
	RecordEverything bool               `json:"recordEverything,omitempty"`
	OSDSettings      *CameraOSDSettings `json:"osdSettings,omitempty"`
	SmartDetect      *CameraSmartDetect `json:"smartDetectSettings,omitempty"`
}

type CameraOSDSettings struct {
	IsNameEnabled bool `json:"isNameEnabled"`
	IsDateEnabled bool `json:"isDateEnabled"`
}

type CameraSmartDetect struct {
	ObjectTypes []string `json:"objectTypes"`
}

func (h *ProtectHandler) GetCamera(ctx context.Context, cameraID string) (*CameraDto, error) {
	path := "/v1/cameras/" + cameraID
	var resp CameraDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *ProtectHandler) UpdateCamera(ctx context.Context, cameraID string, req *CameraDto) (*CameraDto, error) {
	path := "/v1/cameras/" + cameraID
	var resp CameraDto
	if err := h.Request(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
