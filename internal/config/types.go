package bloc4_config

type ServerCnf struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Https             bool   `yaml:"https"`
	CertPath          string `yaml:"certPath"`
	ReadTimeout       int64  `yaml:"readTimeout"`
	ReadHeaderTimeout int64  `yaml:"readHeaderTimeout"`
	WriteTimeout      int64  `yaml:"writeTimeout"`
	IdleTimeout       int64  `yaml:"idleTimeout"`
	MaxHeaderBytes    int    `yaml:"maxHeaderBytes"`
	// MaxConnections    int    `yaml:"MaxConnections"`
}

type ApiCnf struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	GatewayName     string `yaml:"gatewayName"`
	NeedAuth        bool   `yaml:"needAuth"`
	AuthType        string `yaml:"authType"`
	Description     string `yaml:"description"`
	Protocol        string `yaml:"protocol"`
	GlobalRateLimit int64  `yaml:"globalRateLimit"`
}

type Config struct {
	Server ServerCnf
	Apis   []ApiCnf
}

type configType int

const (
	cSERVER configType = iota
	cAPI
)

type configPair struct {
	kind configType
	path string
}
