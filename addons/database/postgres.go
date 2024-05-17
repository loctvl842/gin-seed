package database

import (
	"app/addons"
	"app/addons/logger"
	"app/gosdk"
	"flag"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDbOpt struct {
	Prefix   addons.AddOnPrefix
	Host     string
	User     string
	Password string
	DbName   string
	Port     int
}

type postgresDb struct {
	*PostgresDbOpt
	logger         logger.Logger
	applicationCtx gosdk.ApplicationContext
	name           addons.AddOnName
	db             *gorm.DB
}

func NewPgDatabase(name addons.AddOnName, prefix addons.AddOnPrefix) *postgresDb {
	return &postgresDb{
		name: name,
		PostgresDbOpt: &PostgresDbOpt{
			Prefix: prefix,
		},
	}
}

// Implement gosdk.Application interface
func (s *postgresDb) Name() string { return string(s.name) }

func (s *postgresDb) GetPrefix() string { return string(s.Prefix) }

func (s *postgresDb) Get() interface{} { return s.db }

// Implement gosdk.PrefixRunnable
func (s *postgresDb) InitFlags() {
	prefix := s.GetPrefix()
	if prefix != "" {
		prefix += "-"
	}

	flag.StringVar(&s.Host, prefix+"host", "localhost", "Postgres host")
	flag.StringVar(&s.User, prefix+"user", "postgres", "Postgres user")
	flag.StringVar(&s.Password, prefix+"password", "postgres", "Postgres password")
	flag.StringVar(&s.DbName, prefix+"db-name", "postgres", "Postgres database name")
	flag.IntVar(&s.Port, prefix+"port", 5432, "Port to listen on")
}

func (s *postgresDb) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(string(s.name))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", s.Host, s.User, s.Password, s.DbName, s.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		s.logger.With("msg", err).Error("Error opening database")
		return err
	}
	s.db = db

	return nil
}

func (s *postgresDb) Run() error {
	if err := s.Configure(); err != nil {
		return err
	}
	s.logger.Info("Running postgres database")

	// Attempt to ping the database to test the connection
	var result int
	if err := s.db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		s.logger.Errorf("Error pinging database: %v", err)
		return err
	} else {
		s.logger.Withs(logger.Fields{
			"host":     s.Host,
			"port":     s.Port,
			"database": s.DbName,
		}).Info("Connected to database")
	}

	return nil
}

func (s *postgresDb) Stop() <-chan bool {
	s.logger.Infof("Stopping %s", s.Name())
	c := make(chan bool)

	go func() {
		c <- true
		dbSQL, err := s.db.DB()
		if err != nil {
			s.logger.Errorf("Error getting database: %v", err)
		}
		if err := dbSQL.Close(); err != nil {
			s.logger.Errorf("Error closing database: %v", err)
		}
	}()

	return c
}
