package configs

import "time"

func Default() Config {
	return Config{
		Listen: ":1234",
		DataBase: DbConfig{
			ConnectionString: "mongodb+srv://user:1qaz%21QAZ@cluster0.9yafuwb.mongodb.net/?retryWrites=true&w=majority",
			Timeout:          10 * time.Second,
		},
		SECRETKEY: "12345678",
		Caller: CallerConfig{
			Sleep: 10 * time.Second,
		},
		User: UserConfig{
			NoOfUrls: 20,
		},
	}
}
