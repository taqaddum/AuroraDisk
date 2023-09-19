package config

type database struct {
	Type   string
	Passwd string
	User   string
	Host   string
	Port   uint16
	Name   string
	URL    string
}

type system struct {
	Mode          string `validate:"eq=master | eq=slave"`
	Addr          string `validate:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
	Proxy         string `validate:"required_with=Addr"`
	Period        int64  `validate:"gte=0"`
}

type tls struct {
	CertPath string
	KeyPath  string
	Addr     string `validate:"required"`
}

type slave struct {
	Secret       string `validate:"omitempty,gte=64"`
	CallbackTime int    `validate:"omitempty,gte=1"`
	SignatureTTL int    `validate:"omitempty,gte=1"`
}

type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	DB       string
}

type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	SameSite         string
	Secure           bool
}
