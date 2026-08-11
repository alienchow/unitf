package client

import "context"

type SiteOverview struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	InternalReference string `json:"internalReference"`
}

type SiteOverviewPage struct {
	Count      int            `json:"count"`
	Data       []SiteOverview `json:"data"`
	Limit      int            `json:"limit"`
	Offset     int64          `json:"offset"`
	TotalCount int64          `json:"totalCount"`
}

func (c *Client) ListSites(ctx context.Context) ([]SiteOverview, error) {
	path := "/v1/sites"
	var page SiteOverviewPage
	if err := c.Network.Request(ctx, "GET", path, nil, &page); err != nil {
		return nil, err
	}
	return page.Data, nil
}
