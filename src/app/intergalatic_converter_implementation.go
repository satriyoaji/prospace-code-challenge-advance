package app

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *IntergalacticConverter) ProcessInput(input string) error {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handling intergalactic to Roman numeral mapping
		if strings.Count(line, " ") == 2 && strings.Contains(line, " is ") && !strings.Contains(line, " Credits") {
			if err := c.processIntergalacticMapping(line); err != nil {
				return err
			}
			continue
		}

		// Handling metal value definitions
		if strings.Contains(line, " Credits") && !strings.HasPrefix(line, "how many") {
			if err := c.processMetalValue(line); err != nil {
				return err
			}
			continue
		}

		// Handling queries
		if strings.HasPrefix(line, "how much is") || strings.HasPrefix(line, "how many Credits is") {
			result, err := c.processQuery(line)
			if err != nil {
				return err
			}
			fmt.Println(result)
			continue
		}

		// Todo: if any Handling other types of lines or returning an error
	}
	return nil
}

func (c *IntergalacticConverter) processIntergalacticMapping(line string) error {
	parts := strings.Split(line, " is ")
	intergalactic := parts[0]
	roman := parts[1]

	if len(roman) != 1 || !strings.Contains("IVXLCDM", roman) {
		return fmt.Errorf("invalid Roman numeral: %s", roman)
	}

	c.NumeralMappings[intergalactic] = roman
	return nil
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

func (c *IntergalacticConverter) processMetalValue(statement string) error {
	parts := strings.Split(statement, " is ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid metal value definition")
	}
	intergalacticAndMetal := strings.Fields(parts[0])

	metal := intergalacticAndMetal[len(intergalacticAndMetal)-1]
	intergalactic := strings.Join(intergalacticAndMetal[:len(intergalacticAndMetal)-1], " ")

	roman, err := c.intergalacticToRoman(intergalactic)
	if err != nil {
		return err
	}
	arabic, err := c.romanToArabic(roman)
	if err != nil {
		return err
	}

	// Extracting the numeric part of the string (without " Credits")
	creditsPart := strings.Split(parts[1], " ")
	creditsValue, err := strconv.ParseFloat(creditsPart[0], 64) // Parsing only the numeric part
	if err != nil {
		return err
	}

	c.MetalValues[metal] = creditsValue / float64(arabic)
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
			return "", fmt.Errorf("unknown metal: %s", metal)
		}

		credits := float64(arabic) * metalValue
		return fmt.Sprintf("%s %s is %.0f Credits", intergalactic, metal, credits), nil
	}

	return "I have no idea what you are talking about", nil
}
