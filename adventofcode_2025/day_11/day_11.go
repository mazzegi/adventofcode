package day_11

import (
	"fmt"
	"strings"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

type dacFFTResult struct {
	valid            bool
	numWithout       int
	numWithDAC       int
	numWithFFT       int
	numWithDACAndFFT int
}

type device struct {
	name         string
	outputs      []string
	dacFFTResult dacFFTResult
}

func parseDevice(s string) (device, error) {
	name, outs, ok := strings.Cut(s, ":")
	if !ok {
		return device{}, fmt.Errorf("invalid device string %q", s)
	}
	devName := strings.TrimSpace(name)
	if devName == "" {
		return device{}, fmt.Errorf("invalid device string %q", s)
	}

	dev := device{
		name: devName,
	}
	outFields := strings.Fields(outs)
	for _, out := range outFields {
		out = strings.TrimSpace(out)
		if out == "" {
			continue
		}
		dev.outputs = append(dev.outputs, out)
	}
	return dev, nil
}

func part1MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	devices := map[string]device{}
	for _, line := range lines {
		dev, err := parseDevice(line)
		if err != nil {
			return 0, fmt.Errorf("parse_device %q: %w", line, err)
		}
		if _, ok := devices[dev.name]; ok {
			return 0, fmt.Errorf("duplicate device name %q", dev.name)
		}
		devices[dev.name] = dev
	}

	//
	var numPathsToOut func(devName string) (int, error)
	numPathsToOut = func(devName string) (int, error) {
		if devName == "out" {
			return 0, nil
		}
		dev, ok := devices[devName]
		if !ok {
			return 0, fmt.Errorf("didnt find device %q", devName)
		}
		var np int
		for _, output := range dev.outputs {
			if output == "out" {
				np++
				continue
			}
			npSub, err := numPathsToOut(output)
			if err != nil {
				return 0, fmt.Errorf("num_paths_to_out: %w", err)
			}
			np += npSub
		}
		return np, nil
	}

	//
	res, err := numPathsToOut("you")
	if err != nil {
		return 0, fmt.Errorf("num_paths_to_out: %w", err)
	}

	return res, nil
}

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	devices := map[string]*device{}
	for _, line := range lines {
		dev, err := parseDevice(line)
		if err != nil {
			return 0, fmt.Errorf("parse_device %q: %w", line, err)
		}
		if _, ok := devices[dev.name]; ok {
			return 0, fmt.Errorf("duplicate device name %q", dev.name)
		}
		devices[dev.name] = &dev
	}

	var numPathsToOut func(devName string) (numWithout, numWithDAC, numWithFFT, numWithDACAndFFT int, err error)
	numPathsToOut = func(devName string) (numWithout, numWithDAC, numWithFFT, numWithDACAndFFT int, err error) {
		dev, ok := devices[devName]
		if !ok {
			return 0, 0, 0, 0, fmt.Errorf("didnt find device %q", devName)
		}
		if dev.dacFFTResult.valid {
			return dev.dacFFTResult.numWithout, dev.dacFFTResult.numWithDAC, dev.dacFFTResult.numWithFFT, dev.dacFFTResult.numWithDACAndFFT, nil
		}

		var (
			thisNumWithout       int
			thisNumWithDAC       int
			thisNumWithFFT       int
			thisNumWithDACAndFFT int
		)

		for _, output := range dev.outputs {
			if output == "out" {
				thisNumWithout++
				continue
			}
			npSubWithout, npSubDAC, npSubFFT, npSubDACAndFFT, err := numPathsToOut(output)
			if err != nil {
				return 0, 0, 0, 0, fmt.Errorf("num_paths_to_out: %w", err)
			}
			thisNumWithout += npSubWithout
			thisNumWithDAC += npSubDAC
			thisNumWithFFT += npSubFFT
			thisNumWithDACAndFFT += npSubDACAndFFT
		}

		switch devName {
		case "dac":
			// num with dac stays
			numWithout = 0
			numWithDAC = thisNumWithDAC + thisNumWithout
			numWithDACAndFFT = thisNumWithDACAndFFT + thisNumWithFFT
			numWithFFT = 0

		case "fft":
			numWithout = 0
			numWithDAC = 0
			numWithDACAndFFT = thisNumWithDACAndFFT + thisNumWithDAC
			numWithFFT = thisNumWithFFT + thisNumWithout

		default:
			numWithout = thisNumWithout
			numWithDAC = thisNumWithDAC
			numWithFFT = thisNumWithFFT
			numWithDACAndFFT = thisNumWithDACAndFFT
		}
		dev.dacFFTResult.valid = true
		dev.dacFFTResult.numWithout = numWithout
		dev.dacFFTResult.numWithDAC = numWithDAC
		dev.dacFFTResult.numWithFFT = numWithFFT
		dev.dacFFTResult.numWithDACAndFFT = numWithDACAndFFT

		return
	}

	//
	nWithout, nDAC, nFFT, nDACAndFFT, err := numPathsToOut("svr")
	if err != nil {
		return 0, fmt.Errorf("num_paths_to_out: %w", err)
	}
	log("without=%d, dac=%d, fft=%d, dac+fft=%d", nWithout, nDAC, nFFT, nDACAndFFT)

	return nDACAndFFT, nil
}
