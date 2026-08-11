package client

import (
	"context"
)

// ProtectClient defines the interface for Protect API operations.
type ProtectClient interface {
	GetCamera(ctx context.Context, cameraID string) (*CameraDto, error)
	UpdateCamera(ctx context.Context, cameraID string, req *CameraDto) (*CameraDto, error)

	GetSensor(ctx context.Context, sensorID string) (*SensorDto, error)
	UpdateSensor(ctx context.Context, sensorID string, req *SensorDto) (*SensorDto, error)

	GetLight(ctx context.Context, id string) (*LightDto, error)
	UpdateLight(ctx context.Context, id string, req *LightDto) (*LightDto, error)

	GetRelay(ctx context.Context, id string) (*RelayDto, error)
	UpdateRelay(ctx context.Context, id string, req *RelayDto) (*RelayDto, error)

	GetSiren(ctx context.Context, id string) (*SirenDto, error)
	UpdateSiren(ctx context.Context, id string, req *SirenDto) (*SirenDto, error)

	GetChime(ctx context.Context, id string) (*ChimeDto, error)
	UpdateChime(ctx context.Context, id string, req *ChimeDto) (*ChimeDto, error)

	GetViewer(ctx context.Context, id string) (*ViewerDto, error)
	UpdateViewer(ctx context.Context, id string, req *ViewerDto) (*ViewerDto, error)

	CreateLiveview(ctx context.Context, req *LiveviewDto) (*LiveviewDto, error)
	GetLiveview(ctx context.Context, id string) (*LiveviewDto, error)
	UpdateLiveview(ctx context.Context, id string, req *LiveviewDto) (*LiveviewDto, error)
	DeleteLiveview(ctx context.Context, id string) error

	ListCameras(ctx context.Context) ([]CameraDto, error)
	ListNvr(ctx context.Context) ([]struct{ ID string }, error)
	ListSensors(ctx context.Context) ([]struct{ ID string }, error)
	ListLights(ctx context.Context) ([]struct{ ID string }, error)
	ListSpeakers(ctx context.Context) ([]struct{ ID string }, error)
	ListLiveviews(ctx context.Context) ([]struct{ ID string }, error)
	ListRelays(ctx context.Context) ([]struct{ ID string }, error)
	ListSirens(ctx context.Context) ([]struct{ ID string }, error)
	ListChimes(ctx context.Context) ([]struct{ ID string }, error)
	ListViewers(ctx context.Context) ([]struct{ ID string }, error)
	ListFobs(ctx context.Context) ([]struct{ ID string }, error)
	ListUsers(ctx context.Context) ([]struct{ ID string }, error)
}

// Client implements ProtectClient
var _ ProtectClient = (*Client)(nil)
