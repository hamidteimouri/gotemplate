package di

import (
	"fmt"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	"gitlab.com/hamidteimouri/core-lens/data/postgres"
	"gitlab.com/hamidteimouri/core-lens/domain"
	"gitlab.com/hamidteimouri/core-lens/presentation/httppr/handlers"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db           *gorm.DB
	dbDatasource *postgres.Postgres
	lensService  *domain.LensService
	lensHandler  *handlers.LensHandler
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbHost := htenvier.Env("DB_HOST")
	dbPort := htenvier.Env("DB_PORT")
	dbName := htenvier.Env("DB_NAME")
	dbUsername := htenvier.Env("DB_USERNAME")
	dbPassword := htenvier.Env("DB_PASSWORD")
	dbTimezone := htenvier.Env("DB_TIMEZONE")

	// logger of gorm
	gormLogger := logger.Default.LogMode(logger.Silent)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
	db, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{Logger: gormLogger, SkipDefaultTransaction: true, TranslateError: true})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database connection error")
	}
	d, err := db.DB()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to get DB instance")
	}
	if err = d.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database has no ping")
	}

	return db
}

func DbDatasource() *postgres.Postgres {
	if dbDatasource == nil {
		dbDatasource = postgres.NewPostgres(DB())
	}

	return dbDatasource
}

func DomainLensService() *domain.LensService {
	if lensService == nil {
		lensService = domain.NewLensService(DbDatasource())
	}

	return lensService
}

func LensHandler() *handlers.LensHandler {
	if lensHandler == nil {
		lensHandler = handlers.NewLensHandler(DomainLensService())
	}

	return lensHandler
}
