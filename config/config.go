package config

import (
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	HTTPPort       int    `json:"http_port" default:"7070" env:"HTTP_PORT"`
	GRPCPort       int    `json:"grpc_port" default:"7071" env:"GRPC_PORT"`
	PacketSizesStr string `json:"-" default:"250,500,1000,2000,5000" env:"PACKET_SIZES"`
	PacketSizes    []uint `json:"packet_sizes"`
}

// SetPacketSizes sets packet sizes via config file or env variable
func (c *Config) SetPacketSizes() {
	if len(c.PacketSizes) == 0 {
		sizes := strings.Split(c.PacketSizesStr, ",")
		for _, sizeStr := range sizes {
			size, _ := strconv.Atoi(sizeStr)
			c.PacketSizes = append(c.PacketSizes, uint(size))
		}
	}

	// sort packet sizes
	sort.Slice(c.PacketSizes, func(i, j int) bool {
		return c.PacketSizes[i] < c.PacketSizes[j]
	})
}

// LoadConfig loads the configuration from the given file or env variables
func LoadConfig() (*Config, error) {
	configFile := flag.String("config", "config.json", "Path to the config file")
	config := new(Config)
	if err := configor.Load(config, *configFile); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}
	config.SetPacketSizes()
	return config, nil
}
