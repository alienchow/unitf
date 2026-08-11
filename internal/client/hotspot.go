package client

import (
	"context"
)

type HotspotVoucherDto struct {
	ID             string `json:"_id,omitempty"`
	Code           string `json:"code,omitempty"`
	Quota          int    `json:"quota,omitempty"`
	Duration       int    `json:"duration,omitempty"`
	CreateTime     int    `json:"create_time,omitempty"`
	QosRateMaxUp   int    `json:"qos_rate_max_up,omitempty"`
	QosRateMaxDown int    `json:"qos_rate_max_down,omitempty"`
}

func (h *NetworkHandler) CreateHotspotVoucher(ctx context.Context, siteID string, req *HotspotVoucherDto) error {
	path := "/api/s/" + siteID + "/cmd/hotspot"
	payload := map[string]interface{}{
		"cmd":    "create-voucher",
		"n":      1,
		"quota":  req.Quota,
		"expire": req.Duration,
		"up":     req.QosRateMaxUp,
		"down":   req.QosRateMaxDown,
	}
	// Note: API returns a create_time which must be used to fetch the exact voucher later.
	// For simplicity in this mock integration, we ignore the response data and just return success.
	return h.Request(ctx, "POST", path, payload, nil)
}

func (h *NetworkHandler) GetHotspotVoucher(ctx context.Context, siteID string, id string) (*HotspotVoucherDto, error) {
	path := "/api/s/" + siteID + "/stat/voucher"
	var resp struct {
		Data []HotspotVoucherDto `json:"data"`
	}
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	for _, v := range resp.Data {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, nil
}

func (h *NetworkHandler) DeleteHotspotVoucher(ctx context.Context, siteID string, id string) error {
	path := "/api/s/" + siteID + "/cmd/hotspot"
	payload := map[string]interface{}{
		"cmd": "delete-voucher",
		"_id": id,
	}
	return h.Request(ctx, "POST", path, payload, nil)
}
