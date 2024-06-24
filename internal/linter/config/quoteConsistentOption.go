package config

import (
	"fmt"
	"strings"
)

// QuoteType is a type of quote for string.
type QuoteType int

// QuoteType constants.
const (
	DoubleQuote QuoteType = iota
	SingleQuote
)

// QuoteConsistentOption represents the option for the QUOTE_CONSISTENT rule.
type QuoteConsistentOption struct {
	CustomizableSeverityOption
	Quote QuoteType
}

// UnmarshalYAML implements yaml.v2 Unmarshaler interface.
func (r *QuoteConsistentOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var option struct {
		CustomizableSeverityOption `yaml:",inline"`
		Quote                      string `yaml:"quote"`
	}
	if err := unmarshal(&option); err != nil {
		return err
	}

	if 0 < len(option.Quote) {
		supportQuotes := map[string]QuoteType{
			"double": DoubleQuote,
			"single": SingleQuote,
		}
		quote, ok := supportQuotes[option.Quote]
		if !ok {
			var list []string
			for k := range supportQuotes {
				list = append(list, k)
			}
			return fmt.Errorf("%s is an invalid quote. valid options are [%s]",
				option.Quote, strings.Join(list, ","))
		}
		r.Quote = quote
	}
	r.CustomizableSeverityOption = option.CustomizableSeverityOption
	return nil
}

func (r *QuoteConsistentOption) UnmarshalToml(data interface{}) error {
	optionsMap := map[string]interface{}{}
	for k, v := range data.(map[string]interface{}) {
		optionsMap[k] = v.(string)
	}

	if quote, ok := optionsMap["quote"]; ok {
		quoteStr := quote.(string)
		if 0 < len(quoteStr) {
			supportQuotes := map[string]QuoteType{
				"double": DoubleQuote,
				"single": SingleQuote,
			}
			quote, ok := supportQuotes[quoteStr]
			if !ok {
				var list []string
				for k := range supportQuotes {
					list = append(list, k)
				}
				return fmt.Errorf("%s is an invalid quote. valid options are [%s]",
					quoteStr, strings.Join(list, ","))
			}
			r.Quote = quote
		}
	}
	return nil
}
