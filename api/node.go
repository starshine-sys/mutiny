package api

// QueryNodeResponse ...
type QueryNodeResponse struct {
	Version   string `json:"revolt"`
	Websocket string `json:"ws"`
	App       string `json:"app"`
	Vapid     string `json:"vapid"`
	Features  QueryNodeFeatures
}

// QueryNodeFeatures ...
type QueryNodeFeatures struct {
	Registration bool `json:"registration"`

	Captcha QueryNodeFeaturesCaptcha `json:"captcha"`

	Email      bool `json:"email"`
	InviteOnly bool `json:"invite_only"`

	Autumn  QueryNodeFeaturesService `json:"autumn"`
	January QueryNodeFeaturesService `json:"january"`
	Voso    QueryNodeFeaturesService `json:"voso"`
}

// QueryNodeFeaturesCaptcha ...
type QueryNodeFeaturesCaptcha struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}

// QueryNodeFeaturesService ...
type QueryNodeFeaturesService struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
	WS      string `json:"ws,omitempty"`
}

// QueryNode returns information about the remote node.
func (c *Client) QueryNode() (*QueryNodeResponse, error) {
	var q QueryNodeResponse
	err := c.RequestJSON(&q, "GET", EndpointQueryNode)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
