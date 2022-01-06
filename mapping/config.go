package mapping

const (
	DefaultMongoAddr  = "mongodb://127.0.0.1:27017"
	DefaultDatabase   = "mapping"
	DefaultCollection = "addr"
	DefaultGidAddr    = "http://gid.mytokenpocket.vip"
	DefaultCacheSize  = 1024
	DefaultExpire     = 3600
)

type AddressConfig struct {
	MongoAddr string
	Database  string
	Col       string
	GidAddr   string
	CacheSize int
	Expire    int64
}
