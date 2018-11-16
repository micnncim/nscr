package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	finder "github.com/b4b4r07/go-finder"
	"github.com/pkg/errors"
)

const (
	packageJSONFile = "package.json"
)

type (
	// PackageJSON extracts npm scripts from package.json
	PackageJSON struct {
		Scripts map[string]string `json:"scripts,omitempty"`
	}

	// Script expresses map for npm scripts commands and scripts themselves
	Script struct {
		Exec    string
		Command string
	}
)

func main() {
	scripts, err := parseScript()
	if err != nil {
		log.Fatal(err)
	}

	script, err := selectScript(scripts)
	if err != nil {
		log.Fatal(err)
	}

	if err := runScript(script); err != nil {
		log.Fatal(err)
	}
}

func parseScript() ([]Script, error) {
	bytes, err := ioutil.ReadFile(packageJSONFile)
	if err != nil {
		return nil, err
	}

	var p PackageJSON
	if err := json.Unmarshal(bytes, &p); err != nil {
		return nil, err
	}

	scripts := make([]Script, 0, len(p.Scripts))
	for k, v := range p.Scripts {
		scripts = append(scripts, Script{
			Exec:    k,
			Command: v,
		})
	}

	return scripts, nil
}

func selectScript(scripts []Script) (string, error) {
	filter, err := finder.New()
	if err != nil {
		return "", err
	}

	lines := make([]string, 0, len(scripts))
	for _, s := range scripts {
		lines = append(lines, fmt.Sprintf("%s\t%s", s.Exec, s.Command))
	}

	selected, err := filter.Select(lines)
	if err != nil {
		return "", err
	}
	if len(selected) == 0 {
		return "", errors.Wrap(err, "must select 1 script")
	}
	if len(selected) > 1 {
		return "", errors.Wrap(err, "cannot select >1 script")
	}

	s, ok := selected[0].(string)
	if !ok {
		return "", errors.Wrap(err, "")
	}

	return strings.Split(s, "\t")[0], nil
}

func runScript(script string) error {
	var cmd string
	if _, err := exec.LookPath("yarn"); err == nil {
		cmd = "yarn"
	} else if exec.LookPath("npm"); err == nil {
		cmd = "npm"
	} else {
		return errors.Wrap(err, "cannot find yarn or npm. need to install one of them")
	}

	bytes, err := exec.Command(cmd, script).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))

	return nil
}
