package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/domarcio/bexs/config"
	commonLog "github.com/domarcio/bexs/src/common/log"
	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/infra/repository"
	"github.com/domarcio/bexs/src/service/connection"
	"github.com/domarcio/bexs/src/service/cost"
)

func main() {
	log := commonLog.NewLogfile(config.Logfile, "[BEXS] ", log.LstdFlags|log.Lmicroseconds|log.Llongfile)

	log.Info("Running cli interface on `%s` environment", config.Env)

	filename, err := handleFilename()
	if err != nil {
		log.Error(err.Error())
		fmt.Fprintf(os.Stdout, "Error on handle filename: %s\n", err.Error())
		os.Exit(1)
	}

	repo, err := repository.NewRouteCSVFile(filename)
	if err != nil {
		log.Error(err.Error())
		fmt.Fprintf(os.Stdout, "Error on repository: %s\n", err.Error())
		os.Exit(1)
	}
	connService := connection.NewService(repo)
	costService := cost.NewService(connService)

	// Waiting for CTRL+C
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, os.Interrupt)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		i := 0
		for {
			i++
			if i == 1 {
				fmt.Fprintf(os.Stdout, "\n$ OK! Your filename is: %s\n\n", filename)
			}
			fmt.Fprintf(os.Stdout, "$ Enter a source and target to check THE BEST route! eg: AAA-BBB: ")

			cmdString, err := reader.ReadString('\n')
			if err != nil {
				log.Error(err.Error())
				fmt.Fprintln(os.Stderr, err)
			}

			source, target, err := formatSourceTarget(cmdString)
			if err != nil {
				log.Error(err.Error())
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			route, err := costService.LowCost(&entity.Airport{Code: source}, &entity.Airport{Code: target})
			if err != nil {
				log.Error(err.Error())
				fmt.Fprintln(os.Stderr, err)
			}
			if route == "" {
				route = "route not found"
			}
			fmt.Fprintf(os.Stdout, "$    >> %s\n", route)
		}
	}()

	<-sg
	log.Info("Finish cli")
	fmt.Fprintln(os.Stdout, "$ bye!")
}

func handleFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("invalid filename")
	}

	filename := strings.TrimSpace(os.Args[1])
	return filename, nil
}

func formatSourceTarget(input string) (source string, target string, err error) {
	il := len(input)
	if il < 7 || il > 8 {
		return "", "", errors.New("invalid input length")
	}

	rg, err := regexp.Match("[A-Z]{3}-[A-Z]{3}", []byte(input))
	if err != nil {
		fmt.Errorf("Error on regexp: %s", err.Error())
		os.Exit(1)
	}
	if !rg {
		return "", "", errors.New("no match found")
	}

	split := strings.Split(input, "-")
	source = strings.TrimSpace(split[0])
	target = strings.TrimSpace(split[1])

	return source, target, nil
}
