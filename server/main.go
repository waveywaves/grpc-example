package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/waveywaves/grpc-example/server/grpc"
	v1 "github.com/waveywaves/grpc-example/server/service"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "3000", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", os.Getenv("MYSQL_PORT_3306_TCP_ADDR"), "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", os.Getenv("DATABASE_USER"), "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", os.Getenv("DATABASE_PASSWORD"), "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "ToDo", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
