package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/timothygan/cafewhere/backend/internal/models"
)

var daysOfWeek = []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
var fullWeekNames = map[string]string{"Mo": "Monday", "Tu": "Tuesday", "We": "Wednesday", "Th": "Thursday", "Fr": "Friday", "Sa": "Saturday", "Su": "Sunday"}

func ParseOpeningHours(input string) (*models.WeekOpeningHours, error) {
	week := &models.WeekOpeningHours{}
	days := map[string]*models.DayHours{
		"Mo": &week.Monday,
		"Tu": &week.Tuesday,
		"We": &week.Wednesday,
		"Th": &week.Thursday,
		"Fr": &week.Friday,
		"Sa": &week.Saturday,
		"Su": &week.Sunday,
	}

	// Initialize all days as closed
	for dayAbbrev, day := range days {
		day.IsClosed = true
		day.Day = fullWeekNames[dayAbbrev]
		day.OpenTime = "n/a"
		day.CloseTime = "n/a"
	}

	// Split the input into individual day rules
	rules := strings.Split(input, ";")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}

		// Check for 24/7
		if rule == "24/7" {
			for _, day := range days {
				day.Is24Hours = true
				day.IsClosed = false
			}
			continue
		}

		// Parse day ranges and times
		parts := strings.SplitN(rule, " ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid rule format: %s", rule)
		}

		dayRange, times := parts[0], parts[1]
		affectedDays, err := parseDayRange(dayRange)
		if err != nil {
			return nil, err
		}

		// Parse times
		openTime, closeTime, err := parseTimes(times)
		if err != nil {
			return nil, err
		}

		// Apply times to affected days
		for _, day := range affectedDays {
			days[day].OpenTime = openTime
			days[day].CloseTime = closeTime
			days[day].IsClosed = false
		}
	}

	return week, nil
}

func parseDayRange(dayRange string) ([]string, error) {
	var result []string
	ranges := strings.Split(dayRange, ",")
	for _, r := range ranges {
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid day range: %s", r)
			}
			start, end := indexOf(daysOfWeek, parts[0]), indexOf(daysOfWeek, parts[1])
			if start == -1 || end == -1 {
				return nil, fmt.Errorf("invalid day in range: %s", r)
			}
			for i := start; i != (end+1)%7; i = (i + 1) % 7 {
				result = append(result, daysOfWeek[i])
			}
		} else {
			if indexOf(daysOfWeek, r) == -1 {
				return nil, fmt.Errorf("invalid day: %s", r)
			}
			result = append(result, r)
		}
	}
	return result, nil
}

func parseTimes(times string) (string, string, error) {
	if times == "off" {
		return "", "", nil
	}
	parts := strings.Split(times, "-")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid time format: %s", times)
	}
	openTime, err := parseTime(parts[0])
	if err != nil {
		return "", "", err
	}
	closeTime, err := parseTime(parts[1])
	if err != nil {
		return "", "", err
	}
	return openTime, closeTime, nil
}

func parseTime(t string) (string, error) {
	re := regexp.MustCompile(`^(\d{2}):(\d{2})$`)
	matches := re.FindStringSubmatch(t)
	if matches == nil {
		return "", fmt.Errorf("invalid time format: %s", t)
	}
	hour, minute := matches[1], matches[2]
	return fmt.Sprintf("%s:%s", hour, minute), nil
}

func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// Helper function to print the week's opening hours
func PrintWeekOpeningHours(week *models.WeekOpeningHours) {
	days := []struct {
		name  string
		hours models.DayHours
	}{
		{"Monday", week.Monday},
		{"Tuesday", week.Tuesday},
		{"Wednesday", week.Wednesday},
		{"Thursday", week.Thursday},
		{"Friday", week.Friday},
		{"Saturday", week.Saturday},
		{"Sunday", week.Sunday},
	}

	for _, day := range days {
		if day.hours.Is24Hours {
			fmt.Printf("%s: Open 24 hours\n", day.name)
		} else if day.hours.IsClosed {
			fmt.Printf("%s: Closed\n", day.name)
		} else {
			fmt.Printf("%s: %s - %s\n", day.name, day.hours.OpenTime, day.hours.CloseTime)
		}
	}
}

// Example usage
func main() {
	testCases := []string{
		"Mo-Fr 08:00-17:00; Sa 09:00-12:00; Su off",
		"Mo-Su 10:00-22:00",
		"24/7",
		"Mo,We,Fr 09:00-17:00; Tu,Th 09:00-19:00; Sa 10:00-15:00; Su off",
	}

	for _, tc := range testCases {
		fmt.Printf("Parsing: %s\n", tc)
		week, err := ParseOpeningHours(tc)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			PrintWeekOpeningHours(week)
		}
		fmt.Println()
	}
}
