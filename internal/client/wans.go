package client

import "context"

type WanOverview struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IPAddress string `json:"ipAddress,omitempty"`
}

type WanOverviewPage struct {
	Count      int           `json:"count"`
	Data       []WanOverview `json:"data"`
	Limit      int           `json:"limit"`
	Offset     int64         `json:"offset"`
	TotalCount int64         `json:"totalCount"`
}

func (h *NetworkHandler) ListWans(ctx context.Context, siteID string) ([]WanOverview, error) {
	path := "/v1/sites/" + siteID + "/wans"
	var page WanOverviewPage
	if err := h.Request(ctx, "GET", path, nil, &page); err != nil {
		return nil, err
	}
	return page.Data, nil
}
