package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/syberalexis/linky-exporter/pkg/core"
	"github.com/syberalexis/linky-exporter/pkg/prom"
)

var (
	// Default variables
	version          = "dev"
	defaultPort      = 9901
	defaultAddress   = "0.0.0.0"
	defaultBaudRate  = 1200
	defaultFrameSize = 7
	defaultParity    = "ParityNone"
	defaultStopBits  = "Stop1"

	// Flags
	debug      bool
	address    string
	port       int
	auto       bool
	historical bool
	standard   bool
	device     string
	baudRate   int
	size       int
	parity     string
	stopBits   string
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "linky-exporter",
		Version: version,
		Short:   "Prometheus exporter for Linky smart meters",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}

	// Define flags
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(
		&address,
		"address",
		"a",
		defaultAddress,
		"Listen address")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", defaultPort, "Listen port")
	rootCmd.PersistentFlags().BoolVar(&auto, "auto", false, "Automatique mode")
	rootCmd.PersistentFlags().BoolVar(&historical, "historical", false, "Historical mode")
	rootCmd.PersistentFlags().BoolVar(&standard, "standard", false, "Standard mode")
	rootCmd.PersistentFlags().StringVarP(&device, "device", "d", "", "Device to read")
	err := rootCmd.MarkPersistentFlagRequired("device")
	if err != nil {
		slog.Error("Error during flag parsing", "error", err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().IntVarP(&baudRate, "baud", "b", defaultBaudRate, "Baud rate")
	rootCmd.PersistentFlags().IntVar(&size, "size", defaultFrameSize, "Serial frame size")
	rootCmd.PersistentFlags().StringVar(
		&parity,
		"parity",
		defaultParity,
		"Serial parity (ParityNone, N, ParityOdd, O, ParityEven, E, ParityMark, M, ParitySpace, S)")
	rootCmd.PersistentFlags().StringVar(
		&stopBits,
		"stopbits",
		defaultStopBits,
		"Serial stopbits (Stop1, 1, Stop1Half, 15, Stop2, 2)")

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}

// Main run function
func run() {
	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Info("Debug mode enabled !")
	}

	// Checks before running
	_, err := os.Stat(device)
	if err != nil {
		slog.Error("Device not found", "error", err)
	}

	// Parse parameters
	connector := core.LinkyConnector{Device: device}
	detect := auto
	if !detect {
		if standard {
			connector.Mode = core.Standard
			connector.BaudRate = core.Standard.BaudRate
			connector.FrameSize = core.Standard.FrameSize
			connector.Parity = core.Standard.Parity
			connector.StopBits = core.Standard.StopBits
		} else if historical {
			connector.Mode = core.Historical
			connector.BaudRate = core.Historical.BaudRate
			connector.FrameSize = core.Historical.FrameSize
			connector.Parity = core.Historical.Parity
			connector.StopBits = core.Historical.StopBits
		} else {
			detect = true
		}
		if baudRate != 0 {
			connector.BaudRate = baudRate
		}
		if size != 0 {
			connector.FrameSize = size
		}
		if parity != "" {
			slog.Debug("Parsing parity", "parity", parity)
			connector.Parity = core.ParseParity(parity)
		}
		if stopBits != "" {
			slog.Debug("Parsing stop bits configuration", "stopBits", stopBits)
			connector.StopBits = core.ParseStopBits(stopBits)
		}
	}

	// Auto detection mode
	if detect {
		err := connector.Detect()
		slog.Debug("Connector configuration",
			"device", connector.Device,
			"mode", connector.Mode,
			"baudrate", connector.BaudRate,
			"framesize", connector.FrameSize,
			"parity", connector.Parity,
			"stopbits", connector.StopBits)
		if err != nil {
			slog.Error("Error during auto detection", "error", err)
		}
	}

	// Run exporter
	exporter := prom.LinkyExporter{Address: address, Port: port}
	exporter.Run(&connector)
}
