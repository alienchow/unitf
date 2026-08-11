package client

import (
	"context"
)

type SensorDto struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Alarm     bool   `json:"alarm,omitempty"`
	TempLimit int    `json:"tempLimit,omitempty"`
}

func (h *ProtectHandler) GetSensor(ctx context.Context, sensorID string) (*SensorDto, error) {
	path := "/v1/sensors/" + sensorID
	var resp SensorDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *ProtectHandler) UpdateSensor(ctx context.Context, sensorID string, req *SensorDto) (*SensorDto, error) {
	path := "/v1/sensors/" + sensorID
	var resp SensorDto
	if err := h.Request(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
