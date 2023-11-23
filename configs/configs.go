package configs

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	ServerPort      int
	DBPort          int
	DBHost          string
	DBUser          string
	DBPass          string
	DBName          string
	Secret          string
	RefSecret       string
	MTServerKey     string
	MTClientKey     string
	CloudAPIKey     string
	CloudAPISecret  string
	CloudName       string
	CloudFolderName string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file, ", err.Error())
		return nil
	}

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Port Value, ", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Port Value, ", err.Error())
			return nil
		}
		res.DBPort = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	}

	if val, found := os.LookupEnv("MT_SERVER_KEY"); found {
		res.MTServerKey = val
	}

	if val, found := os.LookupEnv("MT_CLIENT_KEY"); found {
		res.MTClientKey = val
	}

	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		res.CloudName = val
	}

	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		res.CloudAPIKey = val
	}

	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		res.CloudAPISecret = val
	}

	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		res.CloudFolderName = val
	}

	return res
}
