package flags

import (
	"flag"
	"os"
	"strconv"

	"github.com/zalgonoise/cloaki/cmd/config"
)

// ParseFlags will consume the CLI flags as the app is executed
func ParseFlags() *config.Config {
	var conf = &config.Default

	httpPort := flag.Int("port", conf.HTTPPort, "port to use for the HTTP server")
	boltDBPath := flag.String("bolt-path", conf.BoltDBPath, "path to the Bolt database file")
	sqliteDBPath := flag.String("sqlite-path", conf.SQLiteDBPath, "path to the SQLite database file")
	signingKeyPath := flag.String("jwt-key", conf.SigningKeyPath, "path to the JWT signing key file")
	logfilePath := flag.String("logfile-path", conf.LogFilePath, "path to the logfile stored in the service")
	tracefilePath := flag.String("tracefile-path", conf.TraceFilePath, "path to the tracefile stored in the service")

	flag.Parse()
	osFlags := ParseOSEnv()

	conf.Apply(
		config.Port(*httpPort),
		config.BoltDB(*boltDBPath),
		config.SQLiteDB(*sqliteDBPath),
		config.JWTKey(*signingKeyPath),
		config.Logfile(*logfilePath),
		config.Tracefile(*tracefilePath),
	)

	return conf.Merge(osFlags)
}

// ParseOSEnv will consume the OS environment variables associated with this app, when executed
func ParseOSEnv() *config.Config {
	portStr := os.Getenv("CLOAKI_PORT")

	var port int = 0
	var err error
	if portStr != "" {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			port = 0
		}
	}

	return &config.Config{
		HTTPPort:       port,
		BoltDBPath:     os.Getenv("CLOAKI_BOLT_PATH"),
		SQLiteDBPath:   os.Getenv("CLOAKI_SQLITE_PATH"),
		SigningKeyPath: os.Getenv("CLOAKI_JWT_KEY_PATH"),
		LogFilePath:    os.Getenv("CLOAKI_LOGFILE_PATH"),
		TraceFilePath:  os.Getenv("CLOAKI_TRACEFILE_PATH"),
	}
}
