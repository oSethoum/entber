package entber

import "log"

func WithDB(config ...*DBConfig) option {
	return func(e *extension) {
		if len(config) == 0 {
			e.data.DBConfig = new(DBConfig)
		} else {
			e.data.DBConfig = config[0]
		}
		if e.data.DBConfig.Path == "" {
			e.data.DBConfig.Path = "db"
		}
		if e.data.DBConfig.Driver == "" {
			e.data.DBConfig.Driver = SQLite
		} else {
			if !in(e.data.DBConfig.Driver, []string{MySQL, SQLite, PostgreSQL}) {
				log.Fatalln("driver", e.data.DBConfig.Driver, "is not supported")
			}
		}
		switch e.data.DBConfig.Driver {
		case SQLite:
			if e.data.DBConfig.Dsn == "" {
				e.data.DBConfig.Dsn = "file:entber.sqlite?_fk=1&cache=shared"
			}
		case MySQL:
			if e.data.DBConfig.Dsn == "" {
				e.data.DBConfig.Dsn = "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True"
			}
		case PostgreSQL:
			if e.data.DBConfig.Dsn == "" {
				e.data.DBConfig.Dsn = "host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
			}
		}
	}
}

func WithFiber(config ...*FiberConfig) option {
	return func(e *extension) {
		if len(config) > 0 {
			e.data.FiberConfig = config[0]
		}
		e.data.WithFiber = true
	}
}

func WithTS(config ...*TSConfig) option {
	return func(e *extension) {
		if len(config) == 0 {
			e.data.TSConfig = new(TSConfig)
		} else {
			e.data.TSConfig = config[0]
		}
		if e.data.TSConfig.ApiPath == "" {
			e.data.TSConfig.ApiPath = "ts/"
		}
	}
}

func WithAppConfig(config *AppConfig) option {
	return func(e *extension) {
		e.data.AppConfig = config
	}
}
