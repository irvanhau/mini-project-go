package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	res = readData()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func readData() *ProgramConfig {
	var data = new(ProgramConfig)

<<<<<<< HEAD
	data = loadConfig()

	if data == nil {
		err := godotenv.Load(".env")
		data = loadConfig()
		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	var permit = true
=======
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file, ", err.Error())
		return nil
	}
>>>>>>> e57ef60ee5fa38f86220564d2f80a2908eb1e0c6

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Port Value, ", err.Error())
			permit = false
		}
		res.ServerPort = port
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Port Value, ", err.Error())
			permit = false
		}
		res.DBPort = port
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("MT_SERVER_KEY"); found {
		res.MTServerKey = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("MT_CLIENT_KEY"); found {
		res.MTClientKey = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		res.CloudName = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		res.CloudAPIKey = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		res.CloudAPISecret = val
	}else{
		permit = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		res.CloudFolderName = val
	}else{
		permit = false
	}

	if !permit{
		return nil
	}

	return res
}
