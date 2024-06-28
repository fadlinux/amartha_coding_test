package config

type (
	apiAmarthaDB struct {
		ConfigDB apiAmarthaObject `json:"config-db"`
	}

	apiAmarthaObject struct {
		Data config `json:"data"`
	}

	config struct {
		MysqlDBConnection string `json:"mysql_db"`
	}
)
