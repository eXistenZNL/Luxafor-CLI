package main

import (
	"errors"
	"fmt"
	"github.com/leosunmo/goluxafor"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	message, err := handleLuxafor(args)

	if err != nil {
		fmt.Println("ERROR: " + err.Error() + "\n")
		fmt.Println("Usage: " + os.Args[0] + " <mode> <args>")
		os.Exit(1)
	}

	fmt.Println(message)
}

func handleLuxafor(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("expected at least 1 argument")
	}

	switch args[0] {

	case "static":
		if len(args) != 4 {
			return "", errors.New("args: static <red> <green> <blue>")
		}
		uint64red, err := strconv.ParseUint(args[1], 10, 8)
		if err != nil {
			return "", err
		}
		red := uint8(uint64red)
		uint64green, err := strconv.ParseUint(args[2], 10, 8)
		if err != nil {
			return "", err
		}
		green := uint8(uint64green)
		uint64blue, err := strconv.ParseUint(args[3], 10, 8)
		if err != nil {
			return "", err
		}
		blue := uint8(uint64blue)
		luxafor := goluxafor.NewLuxafor()
		err = luxafor.Colour(goluxafor.LedAll, red, green, blue, 30)
		luxafor.Close()
		if err != nil {
			return "", err
		}

	case "strobe":
		if len(args) != 6 {
			return "", errors.New("args: strobe <red> <green> <blue> <speed> <repeat>")
		}
		uint64red, err := strconv.ParseUint(args[1], 10, 8)
		if err != nil {
			return "", err
		}
		red := uint8(uint64red)
		uint64green, err := strconv.ParseUint(args[2], 10, 8)
		if err != nil {
			return "", err
		}
		green := uint8(uint64green)
		uint64blue, err := strconv.ParseUint(args[3], 10, 8)
		if err != nil {
			return "", err
		}
		blue := uint8(uint64blue)
		uint64speed, err := strconv.ParseUint(args[4], 10, 8)
		if err != nil {
			return "", err
		}
		speed := uint8(uint64speed)
		uint64amount, err := strconv.ParseUint(args[5], 10, 8)
		if err != nil {
			return "", err
		}
		amount := uint8(uint64amount)
		luxafor := goluxafor.NewLuxafor()
		err = luxafor.Strobe(goluxafor.LedAll, red, green, blue, speed, amount)
		luxafor.Close()
		if err != nil {
			return "", err
		}

	case "wave":
		if len(args) != 7 {
			return "", errors.New("args: wave <pattern> <red> <green> <blue> <speed> <repeat>")
		}
		var wave goluxafor.Wave = 0x00
		switch args[1] {
		case "1":
			wave = goluxafor.Wave1
			break
		case "2":
			wave = goluxafor.Wave2
			break
		case "3":
			wave = goluxafor.Wave3
			break
		case "4":
			wave = goluxafor.Wave4
			break
		case "5":
			wave = goluxafor.Wave5
			break
		default:
			return "", errors.New("valid waves are: 1-5")
		}
		uint64red, err := strconv.ParseUint(args[2], 10, 8)
		if err != nil {
			return "", err
		}
		red := uint8(uint64red)
		uint64green, err := strconv.ParseUint(args[3], 10, 8)
		if err != nil {
			return "", err
		}
		green := uint8(uint64green)
		uint64blue, err := strconv.ParseUint(args[4], 10, 8)
		if err != nil {
			return "", err
		}
		blue := uint8(uint64blue)
		uint64speed, err := strconv.ParseUint(args[5], 10, 8)
		if err != nil {
			return "", err
		}
		speed := uint8(uint64speed)
		uint64repeat, err := strconv.ParseUint(args[6], 10, 8)
		if err != nil {
			return "", err
		}
		repeat := uint8(uint64repeat)

		luxafor := goluxafor.NewLuxafor()
		err = luxafor.Wave(wave, red, green, blue, speed, repeat)
		luxafor.Close()
		if err != nil {
			return "", err
		}

	case "pattern":
		if len(args) != 2 {
			return "", errors.New("args: wave <pattern>")
		}
		var pattern goluxafor.Pattern = 0x00
		switch args[1] {
		case "1":
			pattern = goluxafor.Pattern1
			break
		case "2":
			pattern = goluxafor.Pattern2
			break
		case "3":
			pattern = goluxafor.Pattern3
			break
		case "4":
			pattern = goluxafor.Pattern4
			break
		case "5":
			pattern = goluxafor.Pattern5
			break
		case "6":
			pattern = goluxafor.Pattern6
			break
		case "7":
			pattern = goluxafor.Pattern7
			break
		case "8":
			pattern = goluxafor.Pattern8
			break
		default:
			return "", errors.New("valid patterns are: 1-8")
		}
		luxafor := goluxafor.NewLuxafor()
		err := luxafor.Pattern(pattern, 0)
		luxafor.Close()
		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("valid modes are: static, strobe, wave, pattern")
	}
	return "Set the mode successfully", nil
}
