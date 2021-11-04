package client

type ClientOption func(*ClientConfig)

func NewClientConfig(opts ...ClientOption) *ClientConfig {
	clientConfig := &ClientConfig{
		TimeoutMs:            10 * 1000,
		BeatInterval:         5 * 1000,
		OpenKMS:              false,
		UpdateThreadNum:      20,
		NotLoadCacheAtStart:  false,
		UpdateCacheWhenEmpty: false,
		RotateTime:           "24h",
		MaxAge:               3,
		LogLevel:             "info",
	}

	for _, opt := range opts {
		opt(clientConfig)
	}


	return clientConfig
}

