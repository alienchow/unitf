package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alienchow/unitf/internal/client"
)

func TestClient_ProtectCameras_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "protect_cameras_massive.json")

	// Server for List operations (returns array)
	tsList := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer tsList.Close()

	cList, err := client.NewClient(tsList.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	cams, err := cList.Protect.ListCameras(ctx)
	if err != nil {
		t.Fatalf("ListCameras failed: %v", err)
	}
	if len(cams) != 15 {
		t.Fatalf("Expected 15 cameras, got %d", len(cams))
	}

	// Server for Single Item operations (returns object)
	tsItem := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": "cam-1", "name": "Front Door Camera"}`))
	}))
	defer tsItem.Close()

	cItem, err := client.NewClient(tsItem.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = cItem.Protect.GetCamera(ctx, cams[0].ID)
	if err != nil {
		t.Errorf("GetCamera failed: %v", err)
	}
	_, err = cItem.Protect.UpdateCamera(ctx, cams[0].ID, &cams[0])
	if err != nil {
		t.Errorf("UpdateCamera failed: %v", err)
	}
}

func TestClient_ProtectSensors_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "protect_sensors_massive.json")

	tsList := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer tsList.Close()

	cList, err := client.NewClient(tsList.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	sensors, err := cList.Protect.ListSensors(ctx)
	if err != nil {
		t.Fatalf("ListSensors failed: %v", err)
	}
	if len(sensors) != 10 {
		t.Fatalf("Expected 10 sensors, got %d", len(sensors))
	}

	tsItem := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": "sensor-1", "name": "Sensor-1"}`))
	}))
	defer tsItem.Close()

	cItem, err := client.NewClient(tsItem.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = cItem.Protect.GetSensor(ctx, sensors[0].ID)
	if err != nil {
		t.Errorf("GetSensor failed: %v", err)
	}
	_, err = cItem.Protect.UpdateSensor(ctx, sensors[0].ID, &client.SensorDto{Name: "Sensor-1"})
	if err != nil {
		t.Errorf("UpdateSensor failed: %v", err)
	}
}

func TestClient_ProtectPeripherals(t *testing.T) {
	// Item server returning single JSON object
	tsItem := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": "peri-1", "name": "Peripheral-1"}`))
	}))
	defer tsItem.Close()

	cItem, err := client.NewClient(tsItem.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// List server returning array
	tsList := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id": "peri-1", "name": "Peripheral-1"}]`))
	}))
	defer tsList.Close()

	cList, err := client.NewClient(tsList.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Chime
	_, err = cItem.Protect.GetChime(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetChime failed: %v", err)
	}
	_, err = cItem.Protect.UpdateChime(ctx, "peri-1", &client.ChimeDto{Name: "Chime-1"})
	if err != nil {
		t.Errorf("UpdateChime failed: %v", err)
	}
	_, err = cList.Protect.ListChimes(ctx)
	if err != nil {
		t.Errorf("ListChimes failed: %v", err)
	}

	// Light
	_, err = cItem.Protect.GetLight(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetLight failed: %v", err)
	}
	_, err = cItem.Protect.UpdateLight(ctx, "peri-1", &client.LightDto{Name: "Light-1"})
	if err != nil {
		t.Errorf("UpdateLight failed: %v", err)
	}
	_, err = cList.Protect.ListLights(ctx)
	if err != nil {
		t.Errorf("ListLights failed: %v", err)
	}

	// Liveview
	_, err = cItem.Protect.GetLiveview(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetLiveview failed: %v", err)
	}
	_, err = cItem.Protect.CreateLiveview(ctx, &client.LiveviewDto{Name: "New LV"})
	if err != nil {
		t.Errorf("CreateLiveview failed: %v", err)
	}
	_, err = cItem.Protect.UpdateLiveview(ctx, "peri-1", &client.LiveviewDto{Name: "Liveview-1"})
	if err != nil {
		t.Errorf("UpdateLiveview failed: %v", err)
	}
	err = cItem.Protect.DeleteLiveview(ctx, "peri-1")
	if err != nil {
		t.Errorf("DeleteLiveview failed: %v", err)
	}
	_, err = cList.Protect.ListLiveviews(ctx)
	if err != nil {
		t.Errorf("ListLiveviews failed: %v", err)
	}

	// Relay
	_, err = cItem.Protect.GetRelay(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetRelay failed: %v", err)
	}
	_, err = cItem.Protect.UpdateRelay(ctx, "peri-1", &client.RelayDto{Name: "Relay-1"})
	if err != nil {
		t.Errorf("UpdateRelay failed: %v", err)
	}
	_, err = cList.Protect.ListRelays(ctx)
	if err != nil {
		t.Errorf("ListRelays failed: %v", err)
	}

	// Siren
	_, err = cItem.Protect.GetSiren(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetSiren failed: %v", err)
	}
	_, err = cItem.Protect.UpdateSiren(ctx, "peri-1", &client.SirenDto{Name: "Siren-1"})
	if err != nil {
		t.Errorf("UpdateSiren failed: %v", err)
	}
	_, err = cList.Protect.ListSirens(ctx)
	if err != nil {
		t.Errorf("ListSirens failed: %v", err)
	}

	// Viewer
	_, err = cItem.Protect.GetViewer(ctx, "peri-1")
	if err != nil {
		t.Errorf("GetViewer failed: %v", err)
	}
	_, err = cItem.Protect.UpdateViewer(ctx, "peri-1", &client.ViewerDto{Name: "Viewer-1"})
	if err != nil {
		t.Errorf("UpdateViewer failed: %v", err)
	}
	_, err = cList.Protect.ListViewers(ctx)
	if err != nil {
		t.Errorf("ListViewers failed: %v", err)
	}

	// Protect Stubs
	_, _ = cList.Protect.ListFobs(ctx)
	_, _ = cList.Protect.ListSpeakers(ctx)
	_, _ = cList.Protect.ListUsers(ctx)
}
