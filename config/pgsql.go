package config

func (p *PgsqlConfig) Dsn() string {
	return "host=" + p.Host + " user=" + p.Name + " password=" + p.Password + " dbname=" + p.DbName + " port=" + p.Port
}
