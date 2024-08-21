package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	fb "github.com/Eugene-Usachev/fastbytes"
	"github.com/Eugene-Usachev/fst"
	"log"
	"messager/src/internal/handler"
	"messager/src/internal/repository"
	"messager/src/internal/server"
	"messager/src/internal/service"
	"messager/src/pkg/logger"
	"os"
	"strconv"
	"time"
)

type config struct {
	isProduction bool
	port         int

	esAddr []string
	esUser string
	esPass string

	postgresHost   string
	postgresPort   int
	postgresUser   string
	postgresPass   string
	postgresDBName string

	redisAddr     string
	redisPassword string

	prometheusAddr string

	fstAccessKey  string
	fstRefreshKey string
}

func getConfig() *config {
	c := &config{}

	isProduction := os.Getenv("IS_PRODUCTION")
	if isProduction != "" {
		c.isProduction, _ = strconv.ParseBool(isProduction)
	} else {
		log.Fatal("IS_PRODUCTION is not set")
	}

	port := os.Getenv("PORT")
	if port != "" {
		var err error
		c.port, err = strconv.Atoi(port)
		if err != nil {
			log.Fatal("Failed to parse PORT: ", err)
		}
	} else {
		log.Fatal("PORT is not set")
	}

	esAddresses := os.Getenv("ES_ADDRESSES")
	if esAddresses != "" {
		var addresses []string
		if err := json.Unmarshal(fb.S2B(esAddresses), &addresses); err != nil {
			log.Fatal("Failed to unmarshal ES_ADDRESSES: ", err)
		}
		c.esAddr = addresses
	} else {
		log.Fatal("ES_ADDR is not set")
	}

	esUser := os.Getenv("ES_USERNAME")
	if esUser != "" {
		c.esUser = esUser
	} else {
		log.Fatal("ES_USERNAME is not set")
	}

	esPass := os.Getenv("ES_PASSWORD")
	if esPass != "" {
		c.esPass = esPass
	} else {
		log.Fatal("ES_PASSWORD is not set")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost != "" {
		c.postgresHost = postgresHost
	} else {
		log.Fatal("POSTGRES_HOST is not set")
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort != "" {
		var err error
		c.postgresPort, err = strconv.Atoi(postgresPort)
		if err != nil {
			log.Fatal("POSTGRES_PORT is not a number")
		}
	} else {
		log.Fatal("POSTGRES_PORT is not set")
	}

	postgresUser := os.Getenv("POSTGRES_USERNAME")
	if postgresUser != "" {
		c.postgresUser = postgresUser
	} else {
		log.Fatal("POSTGRES_USERNAME is not set")
	}

	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	if postgresPass != "" {
		c.postgresPass = postgresPass
	} else {
		log.Fatal("POSTGRES_PASSWORD is not set")
	}

	postgresDBName := os.Getenv("POSTGRES_DB_NAME")
	if postgresDBName != "" {
		c.postgresDBName = postgresDBName
	} else {
		log.Fatal("POSTGRES_DATABASE is not set")
	}

	redisAddr := os.Getenv("REDIS_ADDRESS")
	if redisAddr != "" {
		c.redisAddr = redisAddr
	} else {
		log.Fatal("REDIS_ADDRESS is not set")
	}

	redisPort := os.Getenv("REDIS_PASSWORD")
	if redisPort != "" {
		c.redisPassword = redisPort
	} else {
		log.Fatal("REDIS_PORT is not set")
	}

	fstAccessKey := os.Getenv("FST_ACCESS_KEY")
	if fstAccessKey != "" {
		c.fstAccessKey = fstAccessKey
	} else {
		log.Fatal("FST_ACCESS_KEY is not set")
	}

	fstRefreshKey := os.Getenv("FST_REFRESH_KEY")
	if fstRefreshKey != "" {
		c.fstRefreshKey = fstRefreshKey
	} else {
		log.Fatal("FST_REFRESH_KEY is not set")
	}

	prometheusAddr := os.Getenv("PROMETHEUS_ADDRESS")
	if prometheusAddr != "" {
		c.prometheusAddr = prometheusAddr
	} else {
		log.Fatal("PROMETHEUS_ADDRESS is not set")
	}

	return c
}

func main() {
	logger := logger.NewTempLogger()
	postgresLogger := repository.NewPostgresLogger(logger)

	pool, err := repository.NewPostgresDB(context.Background(), 8, repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USERNAME"),
		UserPass: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}, postgresLogger)

	if err != nil {
		log.Fatal("cant conect to database " + err.Error())
	}

	repositoryImpl := repository.NewRepository(pool, logger)

	accessConverter := fst.NewEncodedConverter(&fst.ConverterConfig{
		SecretKey:      fb.S2B(os.Getenv("JWT_SECRET_KEY")),
		Postfix:        nil,
		ExpirationTime: 15 * time.Minute,
		HashType:       sha256.New,
		//WithExpirationTime: true,
	})
	refreshConverter := fst.NewEncodedConverter(&fst.ConverterConfig{
		SecretKey:      fb.S2B(os.Getenv("JWT_SECRET_KEY_FOR_REFRESH_TOKEN")),
		Postfix:        nil,
		ExpirationTime: 31 * time.Hour * 24,
		HashType:       sha256.New,
		//WithExpirationTime: true,
	})

	serviceImpl := service.NewService(logger, repositoryImpl, accessConverter, refreshConverter)

	handlerImpl := handler.NewHandler(logger, serviceImpl)

	server := server.NewEchoServer()
	server.InitRoutes(handlerImpl)
	server.Run(":8080")
}
