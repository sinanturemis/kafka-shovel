package configuration

func ReadShovelConfig() []ShovelConfig {
	return []ShovelConfig{
		{
			ErrorTopic: "tex.lm-collect.cancelled-package-not-found.0.tex.lm-collect.collect-delivery-service.error",
			RetryTopic: "tex.lm-collect.cancelled-package-not-found.0.tex.lm-collect.collect-delivery-service.retry",
			Active:     true,
			RetryCount: 5,
			Interval:   "1m",
		},
	}
}
