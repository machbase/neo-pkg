package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	var writer io.Writer
	if ghOut := os.Getenv("GITHUB_OUTPUT"); ghOut != "" {
		f, _ := os.OpenFile(ghOut, os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		writer = f
	} else {
		writer = os.Stdout
	}

	names := strings.Fields(os.Args[1])
	plans := []*BuildPlan{}
	for _, pkgName := range names {
		meta, err := parsePackageMetaFile(filepath.Join("projects", pkgName, "package.yml"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if meta.InjectRecipe.Type == "web" {
			plan := &BuildPlan{
				Platform: BuildPlatform{
					OS:        []string{"ubuntu-latest"},
					Name:      "linux+x86-64",
					Container: "debian:buster-slim",
					TinyName:  "*nix64",
				},
				Pkg: pkgName,
			}
			plans = append(plans, plan)
		} else {
			fmt.Println("package does not support web injection")
			os.Exit(1)
		}
	}
	sb := &strings.Builder{}
	if err := json.NewEncoder(sb).Encode(plans); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(writer, "matrix=%s", sb.String())
}

type BuildPlan struct {
	Platform BuildPlatform `json:"platform"`
	Pkg      string        `json:"pkg"`
}

type BuildPlatform struct {
	OS        []string `json:"os"`
	Name      string   `json:"name"`
	Container string   `json:"container,omitempty"`
	TinyName  string   `json:"tinyname,omitempty"`
}

type PackageMeta struct {
	InjectRecipe    InjectRecipe     `yaml:"inject" json:"inject"`
	Distributable   Distributable    `yaml:"distributable" json:"distributable"`
	Description     string           `yaml:"description" json:"description"`
	BuildRecipe     BuildRecipe      `yaml:"build" json:"build"`
	Provides        []string         `yaml:"provides" json:"provides"`
	TestRecipe      *TestRecipe      `yaml:"test,omitempty" json:"test,omitempty"`
	InstallRecipe   *InstallRecipe   `yaml:"install,omitempty" json:"install,omitempty"`
	UninstallRecipe *UninstallRecipe `yaml:"uninstall,omitempty" json:"uninstall,omitempty"`
}

type InjectRecipe struct {
	Type string `yaml:"type"`
}

type Distributable struct {
	Github          string `yaml:"github"`
	Url             string `yaml:"url"`
	StripComponents int    `yaml:"strip_components"`
}

type BuildRecipe struct {
	Script []string `yaml:"script"`
	Env    []string `yaml:"env"`
}

type TestRecipe struct {
	Script string `yaml:"script"`
}

type InstallRecipe struct {
	Script string `yaml:"script"`
}

type UninstallRecipe struct {
	Script string `yaml:"script"`
}

func parsePackageMetaFile(path string) (*PackageMeta, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ret := &PackageMeta{}
	if err := yaml.Unmarshal(content, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
