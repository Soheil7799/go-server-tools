package config

type Server struct {
	Name        string `yaml:"name"`
	Host        string `yaml:"host"`
	Description string `yaml:"description"`
}

type SSHKey struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	Description string `yaml:"description"`
}

type Config struct {
	Servers []Server `yaml:"servers"`
	SSHKeys []SSHKey `yaml:"ssh_keys"`
}
