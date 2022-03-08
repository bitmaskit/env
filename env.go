package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadEnvVars loads environment variables from 1 or many files
func LoadEnvVars(first string, filenames ...string) (err error) {
	params := []string{first}
	if len(filenames) > 0 {
		params = append(params, filenames...)
	}
	vars, err := readEnv(params...)
	if err != nil {
		return err
	}

	return SetEnvVars(vars)
}

// PrintKeys prints the keys that were loaded from files
func PrintKeys() {
	fmt.Print(keys[0])
	for i := 1; i < len(keys); i++ {
		fmt.Print(", ", keys[i])
	}
	fmt.Println()
}

var keys []string

func readEnv(filenames ...string) (res map[string]string, err error) {
	res = map[string]string{}
	for _, f := range filenames {
		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), "=")
			res[split[0]] = split[1]
			keys = append(keys, split[0])
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SetEnvVars sets environment variables from map
func SetEnvVars(vars map[string]string) error {
	for k, v := range vars {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}
