package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	Version      string
	BasicAuth    string
	LogPath      string
	ScanPath     string
	HttpPort     int
	LifeWindow   int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	JwtSecret             string
	Threads               int
	Detailed              bool
	Proxy                 bool
	Engin                 bool
	ScanSpeed             uint32
	AdrrArr               []string
	Target                []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type MasscanSettingS struct {
	Rate      string
	IpFile    string
	IpNotScan string
	Port      string
}

type EsSettingS struct {
	Host string
	User string
	Pwd  string
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
