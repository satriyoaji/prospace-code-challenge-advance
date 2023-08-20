package app

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *IntergalacticConverter) ProcessInput(input string) string {
	if strings.HasSuffix(input, "Credits") {
		if err := c.processMetalValue(input); err != nil {
			return "Error: " + err.Error()
		}
		return ""
	} else if strings.HasPrefix(input, "how much") || strings.HasPrefix(input, "how many") || strings.HasPrefix(input, "Does") || strings.HasPrefix(input, "Is") {
		response, err := c.processQuery(input)
		if err != nil {
			return "Error: " + err.Error()
		}
		return response
	} else {
		parts := strings.Split(input, " is ")
		if len(parts) == 2 {
			c.NumeralMappings[parts[0]] = parts[1]
		}
		return ""
	}
}

func (c *IntergalacticConverter) intergalacticToRoman(intergalactic string) (string, error) {
	words := strings.Fields(intergalactic)
	var romanBuilder strings.Builder

	for _, word := range words {
		romanNumeral, found := c.NumeralMappings[word]
		if !found {
			return "", fmt.Errorf("unknown intergalactic numeral: %s", word)
		}
		romanBuilder.WriteString(romanNumeral)
	}

	return romanBuilder.String(), nil
}

func (c *IntergalacticConverter) romanToArabic(roman string) (int, error) {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var lastValue, total int
	for i := len(roman) - 1; i >= 0; i-- {
		value, exists := values[rune(roman[i])]
		if !exists {
			return 0, fmt.Errorf("invalid Roman numeral symbol: %v", roman[i])
		}

		if value < lastValue {
			total -= value
		} else {
			total += value
		}
		lastValue = value
	}

	return total, nil
}

func (c *IntergalacticConverter) processMetalValue(definition string) error {
	parts := strings.Split(definition, " is ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid metal value definition")
	}

	words := strings.Fields(parts[0])
	metal := words[len(words)-1]
	intergalactic := strings.Join(words[:len(words)-2], " ")

	roman, err := c.intergalacticToRoman(intergalactic)
	if err != nil {
		return err
	}

	arabic, err := c.romanToArabic(roman)
	if err != nil {
		return err
	}

	credits, err := strconv.ParseFloat(parts[1][:len(parts[1])-8], 64)
	if err != nil {
		return err
	}

	c.MetalValues[metal] = credits / float64(arabic)
	return nil
}

func (c *IntergalacticConverter) processQuery(query string) (string, error) {
	// need to handle different types of queries here, such as "how much" and "how many".
	// The exact implementation will depend on the specific requirements and format of the queries.
	if strings.HasPrefix(query, "how much is") {
		intergalactic := strings.TrimSuffix(strings.TrimPrefix(query, "how much is "), " ?")
		roman, err := c.intergalacticToRoman(intergalactic)
		if err != nil {
			return "", err
		}

		arabic, err := c.romanToArabic(roman)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s is %d", intergalactic, arabic), nil
	} else if strings.HasPrefix(query, "how many Credits is") {
		intergalacticAndMetal := strings.TrimSuffix(strings.TrimPrefix(query, "how many Credits is "), " ?")
		words := strings.Fields(intergalacticAndMetal)
		metal := words[len(words)-1]
		intergalactic := strings.Join(words[:len(words)-1], " ")

		roman, err := c.intergalacticToRoman(intergalactic)
		if err != nil {
			return "", err
		}

		arabic, err := c.romanToArabic(roman)
		if err != nil {
			return "", err
		}

		metalValue, found := c.MetalValues[metal]
		if !found {
			return "", fmt.Errorf("Unknown metal: %s", metal)
		}

		credits := float64(arabic) * metalValue
		return fmt.Sprintf("%s %s is %.0f Credits", intergalactic, metal, credits), nil
	}

	return "I have no idea what you are talking about", nil
}
