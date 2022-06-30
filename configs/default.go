package configs

import "github.com/danial-riazati/http-monitoring-server/database"

func Default() Config {
	return Config{
		Listen: ":1234",
		DataBase: database.Config{
			ConnectionString: "mongodb+srv://user:1qaz%21QAZ@cluster0.9yafuwb.mongodb.net/test",
		},
	}
}
