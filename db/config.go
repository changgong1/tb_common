package db

type DbInstConfig struct {
	Driver       string
	Url          string
	MaxLifetime  int //wait_timeout/2
	MaxIdleConns int
	MaxOpenConns int
}

func (s *DbInstConfig) GetMaxLifetime() int {
	if s.MaxLifetime <= 0 {
		return 12
	}
	return s.MaxLifetime
}

func (s *DbInstConfig) GetMaxIdleConns() int {
	if s.MaxIdleConns <= 0 {
		return 2
	}
	return s.MaxIdleConns
}

func (s *DbInstConfig) GetMaxOpenConns() int {
	if s.MaxOpenConns <= 0 {
		return 0
	}
	return s.MaxOpenConns
}

type DBConfig struct {
	Instances map[string]DbInstConfig
}
