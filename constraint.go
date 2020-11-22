package main

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Constraints struct {
		ProcessGroups []struct {
			Process string
			Match   string
			Cpu     float64
			Mem     float64
		} `yaml:"processGroups"`
	}
}

type Constraint struct {
	Name          string
	CPUPercent    float64
	MemoryPercent float64
	Processes     []ProcessInfo
}

type ConstraintList []Constraint

// NewConstraintList builds a list of constraint from configuration file
func NewConstraintList(yamlText string) ConstraintList {
	config := Configuration{}
	err := yaml.Unmarshal([]byte(yamlText), &config)
	if err != nil {
		return nil
	}
	constraints := make([]Constraint, 0)
	for _, constraint := range config.Constraints.ProcessGroups {
		processes := make([]ProcessInfo, 0)
		re, err := regexp.Compile(constraint.Match)
		if err == nil {
			processes = append(processes, ProcessInfoSliceFromRegexp(*re)...)
		}
		constraints = append(constraints, Constraint{
			Name:          constraint.Process,
			CPUPercent:    constraint.Cpu,
			MemoryPercent: constraint.Mem,
			Processes:     processes,
		})
	}
	return constraints
}

// Sanitize checks and sanitizes constraint values and produces messages when violated
func (c ConstraintList) Sanitize(options Options) []string {
	messages := make([]string, 0)
	for i := 0; i < len(c); i++ {
		if c[i].CPUPercent < 0 {
			c[i].CPUPercent = defCPUPercent
			messages = append(messages,
				fmt.Sprintf("CPU constraint for %s invalid, setted to default (%.2f%%)",
					c[i].Name, options.CPUPercent))
		}
		if c[i].MemoryPercent < 0 {
			c[i].MemoryPercent = defMemoryPercent
			messages = append(messages,
				fmt.Sprintf("Memory constraint for %s invalid, setted to default (%.2f%%)",
					c[i].Name, options.MemoryPercent))
		}
	}
	return messages
}

// HasProcesses checks if at least one process is available for polling
func (c ConstraintList) HasProcesses() bool {
	for _, constr := range c {
		if len(constr.Processes) > 0 {
			return true
		}
	}
	return false
}
