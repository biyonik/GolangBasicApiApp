package postgresql

type Config struct {
	Host           string
	Port           string
	UserName       string
	Password       string
	Database       string
	MaxConnections string
	MaxConnectionIdleTime string
}
