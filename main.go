package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lily/commands"
	"github.com/filecoin-project/lily/commands/job"
	"github.com/filecoin-project/lily/version"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
	"github.com/filecoin-project/lotus/node/repo"
)

var log = logging.Logger("lily/main")

type UpSchedule struct {
	Height    int64
	Network   uint
	Expensive bool
}

func (u *UpSchedule) String() string {
	return fmt.Sprintf("Height: %d, Network: %d, Expensive: %t", u.Height, u.Network, u.Expensive)
}

type UpScheduleList []*UpSchedule

func (ul UpScheduleList) String() string {
	var sb strings.Builder
	for _, u := range ul {
		sb.WriteString(fmt.Sprintln("\t\t" + u.String()))
	}
	return sb.String()
}

func main() {
	// Set up a context that is canceled when the command is interrupted
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a signal handler to cancel the context
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-interrupt:
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}

	var up UpScheduleList
	for _, u := range filcns.DefaultUpgradeSchedule() {
		up = append(up, &UpSchedule{
			Height:    int64(u.Height),
			Network:   uint(u.Network),
			Expensive: false,
		})
	}

	app := &cli.App{
		Name:    "lily",
		Usage:   "a tool for capturing on-chain state from the filecoin network",
		Version: fmt.Sprintf("Lily Version: \t%s\n   NewestNetworkVersion: \t%d\n   GenesisFile: \t%s\n   DevNet: \t%t\n   UserVersion: \t%s\n   UpgradeSchedule: \n%s", version.String(), build.TestNetworkVersion, build.GenesisFile, build.Devnet, string(build.NodeUserVersion()), up.String()),
		Flags: []cli.Flag{
			commands.ClientAPIFlag,
			commands.ClientTokenFlag,
			&cli.StringFlag{
				Name:        "log-level",
				EnvVars:     []string{"GOLOG_LOG_LEVEL"},
				Value:       "info",
				Usage:       "Set the default log level for all loggers to `LEVEL`",
				Destination: &commands.LilyLogFlags.LogLevel,
			},
			&cli.StringFlag{
				Name:        "log-level-named",
				EnvVars:     []string{"LILY_LOG_LEVEL_NAMED"},
				Value:       "",
				Usage:       "A comma delimited list of named loggers and log levels formatted as name:level, for example 'logger1:debug,logger2:info'",
				Destination: &commands.LilyLogFlags.LogLevelNamed,
			},
			&cli.BoolFlag{
				Name:        "jaeger-tracing",
				EnvVars:     []string{"LILY_JAEGER_TRACING"},
				Value:       false,
				Destination: &commands.LilyTracingFlags.Enabled,
			},
			&cli.StringFlag{
				Name:        "jaeger-service-name",
				EnvVars:     []string{"LILY_JAEGER_SERVICE_NAME"},
				Value:       "lily",
				Destination: &commands.LilyTracingFlags.ServiceName,
			},
			&cli.StringFlag{
				Name:        "jaeger-provider-url",
				EnvVars:     []string{"LILY_JAEGER_PROVIDER_URL"},
				Value:       "http://localhost:14268/api/traces",
				Destination: &commands.LilyTracingFlags.ProviderURL,
			},
			&cli.Float64Flag{
				Name:        "jaeger-sampler-ratio",
				EnvVars:     []string{"LILY_JAEGER_SAMPLER_RATIO"},
				Usage:       "If less than 1 probabilistic metrics will be used.",
				Value:       1,
				Destination: &commands.LilyTracingFlags.JaegerSamplerParam,
			},
			&cli.StringFlag{
				Name:        "prometheus-port",
				EnvVars:     []string{"LILY_PROMETHEUS_PORT"},
				Value:       ":9991",
				Destination: &commands.LilyMetricFlags.PrometheusPort,
			},
			&cli.StringFlag{
				Name:        "redis-addr",
				EnvVars:     []string{"LILY_REDIS_ADDR"},
				Usage:       `Redis server address in "host:port" format`,
				Value:       "",
				Destination: &commands.LilyMetricFlags.RedisAddr,
			},

			&cli.StringFlag{
				Name:        "redis-username",
				EnvVars:     []string{"LILY_REDIS_USERNAME"},
				Usage:       `Username to authenticate the current connection when redis ACLs are used.`,
				Value:       "",
				Destination: &commands.LilyMetricFlags.RedisUsername,
			},

			&cli.StringFlag{
				Name:        "redis-password",
				EnvVars:     []string{"LILY_REDIS_PASSWORD"},
				Usage:       `Password to authenticate the current connection`,
				Value:       "",
				Destination: &commands.LilyMetricFlags.RedisPassword,
			},

			&cli.IntFlag{
				Name:        "redis-db",
				EnvVars:     []string{"LILY_REDIS_DB"},
				Usage:       `Redis DB to select after connection to server`,
				Value:       0,
				Destination: &commands.LilyMetricFlags.RedisDB,
			},
		},
		Commands: []*cli.Command{
			commands.ChainCmd,
			commands.DaemonCmd,
			commands.ExportChainCmd,
			commands.InitCmd,
			commands.LogCmd,
			commands.MigrateCmd,
			commands.NetCmd,
			commands.ShedCmd,
			commands.StopCmd,
			commands.SyncCmd,
			commands.WaitAPICmd,
			commands.FileToCARCmd,
			commands.ImportFromCIDCmd,
			job.JobCmd,
		},
	}
	app.Setup()
	app.Metadata["repoType"] = repo.FullNode
	app.Metadata["traceContext"] = ctx

	if err := app.RunContext(ctx, os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stdout, err.Error())
		if err != nil {
			fmt.Printf("got error in main: %v", err)
		}
		os.Exit(1)
	}
}
