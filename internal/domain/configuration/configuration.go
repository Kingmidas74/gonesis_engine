package configuration

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"strconv"
	"time"
)

const (
	defaultMaxEnergy              = 100
	defaultBrainSize              = 64
	defaultEnergy                 = 20
	defaultCount                  = 0
	defaultDailyCommands          = 1
	defaultReproductionSystemType = enum.ReproductionSystemTypeBudding
	defaultReproductionEnergyCost = 20
	defaultReproductionChance     = .5
	defaultMutationChance         = .1
)

type Ratio struct {
	Width  int `json:"Width"`
	Height int `json:"Height"`
}

type WorldConfiguration struct {
	MazeType enum.MazeType     `json:"MazeType"`
	Topology enum.TopologyType `json:"Topology"`
	Seed     string            `json:"Seed"`
	Ratio    Ratio             `json:"Ratio"`
}

type AgentConfiguration struct {
	MaxEnergy              int                         `json:"MaxEnergy"`
	InitialCount           int                         `json:"InitialCount"`
	MaxDailyCommandCount   int                         `json:"MaxDailyCommandCount"`
	InitialEnergy          int                         `json:"InitialEnergy"`
	BrainVolume            int                         `json:"BrainVolume"`
	ReproductionType       enum.ReproductionSystemType `json:"ReproductionType"`
	ReproductionEnergyCost int                         `json:"ReproductionEnergyCost"`
	ReproductionChance     float64                     `json:"ReproductionChance"`
	MutationChance         float64                     `json:"MutationChance"`
}

type Configuration struct {
	WorldConfiguration      WorldConfiguration `json:"WorldConfiguration"`
	PlantConfiguration      AgentConfiguration `json:"PlantConfiguration"`
	HerbivoreConfiguration  AgentConfiguration `json:"HerbivoreConfiguration"`
	CarnivoreConfiguration  AgentConfiguration `json:"CarnivoreConfiguration"`
	DecomposerConfiguration AgentConfiguration `json:"DecomposerConfiguration"`
	OmnivoreConfiguration   AgentConfiguration `json:"OmnivoreConfiguration"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		WorldConfiguration: WorldConfiguration{
			MazeType: enum.MazeTypeEmpty,
			Topology: enum.TopologyTypeNeumann,
			Seed:     strconv.FormatInt(time.Now().UnixNano(), 10),
		},
		PlantConfiguration: AgentConfiguration{
			MaxEnergy:              defaultMaxEnergy,
			InitialCount:           defaultCount,
			MaxDailyCommandCount:   defaultDailyCommands,
			InitialEnergy:          defaultEnergy,
			BrainVolume:            defaultBrainSize,
			ReproductionType:       defaultReproductionSystemType,
			ReproductionEnergyCost: defaultReproductionEnergyCost,
			ReproductionChance:     defaultReproductionChance,
			MutationChance:         defaultMutationChance,
		},
		HerbivoreConfiguration: AgentConfiguration{
			MaxEnergy:              defaultMaxEnergy,
			InitialCount:           defaultCount,
			MaxDailyCommandCount:   defaultDailyCommands,
			InitialEnergy:          defaultEnergy,
			BrainVolume:            defaultBrainSize,
			ReproductionType:       defaultReproductionSystemType,
			ReproductionEnergyCost: defaultReproductionEnergyCost,
			ReproductionChance:     defaultReproductionChance,
			MutationChance:         defaultMutationChance,
		},
		CarnivoreConfiguration: AgentConfiguration{
			MaxEnergy:              defaultMaxEnergy,
			InitialCount:           defaultCount,
			MaxDailyCommandCount:   defaultDailyCommands,
			InitialEnergy:          defaultEnergy,
			BrainVolume:            defaultBrainSize,
			ReproductionType:       defaultReproductionSystemType,
			ReproductionEnergyCost: defaultReproductionEnergyCost,
			ReproductionChance:     defaultReproductionChance,
			MutationChance:         defaultMutationChance,
		},
		DecomposerConfiguration: AgentConfiguration{
			MaxEnergy:              defaultMaxEnergy,
			InitialCount:           defaultCount,
			MaxDailyCommandCount:   defaultDailyCommands,
			InitialEnergy:          defaultEnergy,
			BrainVolume:            defaultBrainSize,
			ReproductionType:       defaultReproductionSystemType,
			ReproductionEnergyCost: defaultReproductionEnergyCost,
			ReproductionChance:     defaultReproductionChance,
			MutationChance:         defaultMutationChance,
		},
		OmnivoreConfiguration: AgentConfiguration{
			MaxEnergy:              defaultMaxEnergy,
			InitialCount:           defaultCount,
			MaxDailyCommandCount:   defaultDailyCommands,
			InitialEnergy:          defaultEnergy,
			BrainVolume:            defaultBrainSize,
			ReproductionType:       defaultReproductionSystemType,
			ReproductionEnergyCost: defaultReproductionEnergyCost,
			ReproductionChance:     defaultReproductionChance,
			MutationChance:         defaultMutationChance,
		},
	}
}

func NewAgentConfiguration() *AgentConfiguration {
	return &AgentConfiguration{
		MaxEnergy:              defaultMaxEnergy,
		InitialCount:           defaultCount,
		MaxDailyCommandCount:   defaultDailyCommands,
		InitialEnergy:          defaultEnergy,
		BrainVolume:            defaultBrainSize,
		ReproductionType:       defaultReproductionSystemType,
		ReproductionEnergyCost: defaultReproductionEnergyCost,
		ReproductionChance:     defaultReproductionChance,
		MutationChance:         defaultMutationChance,
	}
}

func (c *AgentConfiguration) FromJSON(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), c)
}

func (c *Configuration) FromJSON(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), c)
}
