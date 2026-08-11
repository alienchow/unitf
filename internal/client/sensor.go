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

func (c *Client) GetSensor(ctx context.Context, sensorID string) (*SensorDto, error) {
	path := "/proxy/protect/api/sensors/" + sensorID
	var resp SensorDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSensor(ctx context.Context, sensorID string, req *SensorDto) (*SensorDto, error) {
	path := "/proxy/protect/api/sensors/" + sensorID
	var resp SensorDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
