package config

type Config struct {
}

// WechatConfig 微信相关配置
type WechatConfig struct {
	CorpID     string `yaml:"corp_id"`     // corp_id
	CorpSecret string `yaml:"corp_secret"` // corp_secret
	AgentID    string `yaml:"agent_id"`    // agent_id
}
