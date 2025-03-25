package core

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"go.bug.st/serial"
)

type LinkyConnector struct {
	Mode      LinkyMode
	Device    string
	BaudRate  int
	FrameSize int
	Parity    serial.Parity
	StopBits  serial.StopBits
}

// Detect serial connection mode
func (connector *LinkyConnector) Detect() error {
	slog.Info("Trying to auto detect TIC mode...")

	if connector.trySerial(Standard) {
		slog.Info("Standard Mode detected !")
		connector.Mode = Standard
		connector.BaudRate = Standard.BaudRate
		connector.FrameSize = Standard.FrameSize
		connector.Parity = Standard.Parity
		connector.StopBits = Standard.StopBits
		return nil
	} else {
		slog.Debug("It's not standard mode !")
	}

	if connector.trySerial(Historical) {
		slog.Info("Historical Mode detected !")
		connector.Mode = Historical
		connector.BaudRate = Historical.BaudRate
		connector.FrameSize = Historical.FrameSize
		connector.Parity = Historical.Parity
		connector.StopBits = Historical.StopBits
		return nil
	} else {
		slog.Debug("It's not historical mode !")
	}

	return fmt.Errorf("impossible to auto detect TIC mode ")
}

// Try serial connection and reading
func (connector *LinkyConnector) trySerial(mode LinkyMode) bool {
	m := &serial.Mode{BaudRate: mode.BaudRate, DataBits: mode.FrameSize, Parity: mode.Parity, StopBits: mode.StopBits}
	stream, err := serial.Open(connector.Device, m)
	if err != nil {
		return false
	}

	reader := bufio.NewReader(stream)
	regex := regexp.MustCompile(`^[A-Z0-9\-+]+ +[a-zA-Z0-9 .\-]+ +.$`)

	slog.Debug("Read serial data...")
	for i := 1; i <= 5; i++ {
		bytes, _, _ := reader.ReadLine()
		line := string(bytes)
		slog.Debug("Try line", "number", i, "total", 5, "content", line)
		if regex.MatchString(line) {
			return true
		} else {
			slog.Debug("Regex not match")
		}
	}
	return false
}

// ASCII control characters used in serial communication
const (
	STX = 0x02 // Start of Text - marks the beginning of a data block
	ETX = 0x03 // End of Text - marks the end of a data block
)

// readSerial values
func (connector *LinkyConnector) readSerial() ([][]string, error) {
	slog.Debug("Read serial with config",
		"device", connector.Device,
		"baudrate", connector.BaudRate,
		"framesize", connector.FrameSize,
		"parity", connector.Parity,
		"stopbits", connector.StopBits)
	m := &serial.Mode{BaudRate: connector.BaudRate, DataBits: connector.FrameSize, Parity: connector.Parity, StopBits: connector.StopBits}
	stream, err := serial.Open(connector.Device, m)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(stream)
	started := false
	var values [][]string

	slog.Debug("Read serial data...")
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}

		line := string(bytes)

		// End loop when block ended
		if started && strings.ContainsRune(line, ETX) {
			err = stream.Close()
			if err != nil {
				slog.Error("Failed to close serial", "error", err)
			}
			break
		}

		// Collect data line by line
		if started {
			slog.Debug(line)
			values = append(values, strings.FieldsFunc(line, func(r rune) bool { return r == 0x09 || r == ' ' }))
		}

		// Start reading data when block started
		if strings.ContainsRune(line, STX) {
			started = true
		}
	}
	slog.Debug("Read serial data ended !")

	return values, nil
}

// GetLastHistoricalTicValue return last serial Historical TIC
func (connector *LinkyConnector) GetLastHistoricalTicValue() (*HistoricalTicValue, error) {
	lines, err := connector.readSerial()

	if err != nil {
		slog.Error("Failed to read historical serial", "error", err)
		return nil, err
	}

	values := HistoricalTicValue{}
	for _, line := range lines {
		values.ParseParam(line[0], line[1:])
	}

	return &values, nil
}

// GetLastStandardTicValue return last serial Standard TIC
func (connector *LinkyConnector) GetLastStandardTicValue() (*StandardTicValue, error) {
	lines, err := connector.readSerial()

	if err != nil {
		slog.Error("Failed to read standard serial", "error", err)
		return nil, err
	}

	values := StandardTicValue{}
	for _, line := range lines {
		values.ParseParam(line[0], line[1:])
	}

	return &values, nil
}

// ParseParity from string to serial object
func ParseParity(value string) (parity serial.Parity) {
	switch value {
	case "ParityNone", "N":
		parity = serial.NoParity

	case "ParityOdd", "O":
		parity = serial.OddParity

	case "ParityEven", "E":
		parity = serial.EvenParity

	case "ParityMark", "M":
		parity = serial.MarkParity

	case "ParitySpace", "S":
		parity = serial.SpaceParity

	default:
		slog.Error("Impossible to parse Parity", "error", value)
		os.Exit(3)
	}
	return
}

// ParseStopBits from string to serial object
func ParseStopBits(value string) (stopBits serial.StopBits) {
	switch value {
	case "Stop1", "1":
		stopBits = serial.OneStopBit
	case "Stop1Half", "15":
		stopBits = serial.OnePointFiveStopBits
	case "Stop2", "2":
		stopBits = serial.TwoStopBits
	default:
		slog.Error("Impossible to parse StopBits", "value", value)
		os.Exit(3)
	}
	return
}
