package domain

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config is an abstraction around the map that holds the config values
type Config struct {
	config map[string]interface{}
	lock   sync.RWMutex
}

// SetFromBytes sets the internal config based on a byte array of YAML
func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("config is not a map")
	}

	config, err := convertKeysToStrings(untypedConfig)
	if err != nil {
		return err
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.config = config
	return nil
}

// Get returns the config for a particular service
func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	a, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config is not a map")
	}

	// If no config is defined for the service
	if _, ok = c.config[serviceName]; !ok {
		// Return the base config
		return a, nil
	}

	b, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q config is not a map", serviceName)
	}

	// Merge the maps with the service config taking precedence
	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}

	return config, nil
}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToStrings(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil
}
