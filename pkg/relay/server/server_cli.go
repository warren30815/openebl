package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	formatter "github.com/bluexlab/logrus-formatter"
	"github.com/gobuffalo/pop"
	"github.com/openebl/openebl/pkg/config"
	"github.com/openebl/openebl/pkg/relay/server/storage/postgres"
	"github.com/sirupsen/logrus"
)

type RelayServerApp struct{}

type RelayServerCli struct {
	Server  struct{} `cmd help:"Run relay server."`
	Migrate struct {
		Migrations string `type:"path" default:"migrations" help:"Path to migration folder."`
	} `cmd help:"Migrate database."`
	Config string `type:"path" default:"config.yaml" help:"Path to config file."`
}

type RelayServerConfig struct {
	Database     postgres.DatabaseConfig `yaml:"database"`
	LocalAddress string                  `yaml:"local_address"`
	OtherPeers   []string                `yaml:"other_peers"`
}

func (r *RelayServerApp) Run() error {
	formatter.InitLogger()

	var cli RelayServerCli
	kctx := kong.Parse(&cli)
	switch kctx.Command() {
	case "server":
		return r.runServer(cli)
	case "migrate":
		return r.runMigrate(cli)
	}

	return nil
}

func (r *RelayServerApp) runServer(cli RelayServerCli) error {
	cfg := RelayServerConfig{}
	if err := config.FromFile(cli.Config, &cfg); err != nil {
		logrus.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	eventStorage, err := postgres.NewEventStorageWithConfig(cfg.Database)
	if err != nil {
		logrus.Errorf("failed to create event storage: %v", err)
		os.Exit(1)
	}

	relayServer, err := NewServer(
		WithLocalAddress(cfg.LocalAddress),
		WithPeers(cfg.OtherPeers),
		WithStorage(eventStorage),
	)
	if err != nil {
		logrus.Errorf("failed to create relay server: %v", err)
		os.Exit(1)
	}

	logrus.Info("starting relay server.")
	go func() {
		if err := relayServer.Run(); err != nil {
			logrus.Errorf("failed to start relay server: %v", err)
			os.Exit(1)
		}
	}()

	r.waitForInterrupt()
	relayServer.Close()
	return nil
}

func (r *RelayServerApp) runMigrate(cli RelayServerCli) error {
	cfg := RelayServerConfig{}
	if err := config.FromFile(cli.Config, &cfg); err != nil {
		logrus.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	cd := pop.ConnectionDetails{
		Dialect:  "postgres",
		Database: cfg.Database.Database,
		Host:     cfg.Database.Host,
		Port:     fmt.Sprintf("%d", cfg.Database.Port),
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
	}
	conn, err := pop.NewConnection(&cd)
	if err != nil {
		logrus.Errorf("failed to create connection: %v", err)
		os.Exit(1)
	}

	migrator, err := pop.NewFileMigrator(cli.Migrate.Migrations, conn)
	if err != nil {
		logrus.Errorf("failed to create migrator: %v", err)
		os.Exit(1)
	}
	if err := migrator.Up(); err != nil {
		logrus.Errorf("failed to migrate: %v", err)
		os.Exit(1)
	}

	return nil
}

func (r *RelayServerApp) waitForInterrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server......")
}
