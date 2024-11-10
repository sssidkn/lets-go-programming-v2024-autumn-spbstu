package parser

import (
	"gopkg.in/yaml.v3"

	"anastasiya.soboleva/task-3/internal/models"
	"anastasiya.soboleva/task-3/internal/utils"
)

func ParseConfig(path string) (*models.Configs, error) {
	file, err := utils.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer utils.CloseFile(file)
	var cfg models.Configs
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
