package config

import (
	"fmt"
	"strings"
)

// ConventionType is a type of name case convention.
type ConventionType int

// ConventionType constants.
const (
	ConventionLowerCamel ConventionType = iota + 1
	ConventionUpperSnake
	ConventionLowerSnake
)

// RPCNamesCaseOption represents the option for the RPC_NAMES_CASE rule.
type RPCNamesCaseOption struct {
	CustomizableSeverityOption
	Convention ConventionType
}

// UnmarshalYAML implements yaml.v2 Unmarshaler interface.
func (r *RPCNamesCaseOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var option struct {
		CustomizableSeverityOption `yaml:",inline"`
		Convention                 string `yaml:"convention"`
	}
	if err := unmarshal(&option); err != nil {
		return err
	}

	if 0 < len(option.Convention) {
		supportConventions := map[string]ConventionType{
			"lower_camel_case": ConventionLowerCamel,
			"upper_snake_case": ConventionUpperSnake,
			"lower_snake_case": ConventionLowerSnake,
		}
		convention, ok := supportConventions[option.Convention]
		if !ok {
			var list []string
			for k := range supportConventions {
				list = append(list, k)
			}
			return fmt.Errorf("%s is an invalid name convention. valid options are [%s]",
				option.Convention, strings.Join(list, ","))
		}
		r.Convention = convention
	}
	r.CustomizableSeverityOption = option.CustomizableSeverityOption
	return nil
}

// UnmarshalTOML implements toml Unmarshaler interface.
func (r *RPCNamesCaseOption) UnmarshalTOML(data interface{}) error {
	optionsMap := map[string]interface{}{}
	for k, v := range data.(map[string]interface{}) {
		optionsMap[k] = v.(string)
	}

	if convention, ok := optionsMap["convention"]; ok {
		conventionStr := convention.(string)
		if 0 < len(conventionStr) {
			supportConventions := map[string]ConventionType{
				"lower_camel_case": ConventionLowerCamel,
				"upper_snake_case": ConventionUpperSnake,
				"lower_snake_case": ConventionLowerSnake,
			}
			convention, ok := supportConventions[conventionStr]
			if !ok {
				var list []string
				for k := range supportConventions {
					list = append(list, k)
				}
				return fmt.Errorf("%s is an invalid name convention. valid options are [%s]",
					conventionStr, strings.Join(list, ","))
			}
			r.Convention = convention
		}
	}

	return nil
}
