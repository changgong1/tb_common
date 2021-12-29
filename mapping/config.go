package mapping

type AddressConfig struct {
	MongoAddr string `default:"mongodb://127.0.0.1:27017"`
	Database  string `default:"mapping"`
	Col       string `default:"addr"`
	GidAddr   string `default:"http://gid.mytokenpocket.vip"`
}
