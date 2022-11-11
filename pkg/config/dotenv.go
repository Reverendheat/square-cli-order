package dotenv

import (
	"errors"
	"os"
	"strings"
)

func setEnv(s string) error {
	equalSplit := strings.Split(s, "=")
	if len(equalSplit) < 2 {
		return errors.New("environment file improperly formatted")
	}
	err := os.Setenv(equalSplit[0], equalSplit[1])
	if err != nil {
		return err
	}
	return nil
}

func loadEnvFromFile() ([]byte, error) {
	dat, err := os.ReadFile(".env")
	if err != nil {
		return nil, errors.New("unable to load .env file")
	}
	return dat, nil
}

func Load() error {
	file, err := loadEnvFromFile()
	if err != nil {
		return err
	}
	newLineSplit := strings.Split(string(file), "\n")
	for _, s := range newLineSplit {
		if len(s) != 0 {
			ts := strings.TrimSuffix(s, "\n")
			tr := strings.TrimSuffix(ts, "\r")
			err := setEnv(tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
