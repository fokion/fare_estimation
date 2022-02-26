package converters

import (
	"errors"
	"fare_estimation/models"
	"fmt"
	"strconv"
	"strings"
)

func ConvertLineToPoint(line string) (string, *models.Point, error) {
	parts := strings.Split(strings.Trim(line, " "), ",")
	lat, err := strconv.ParseFloat(strings.Trim(parts[1], " "), 64)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("line %s does not have a valid latitude", line))
	}
	lon, errLon := strconv.ParseFloat(strings.Trim(parts[2], " "), 64)
	if errLon != nil {
		return "", nil, errors.New(fmt.Sprintf("line %s does not have a valid longitude", line))
	}

	timestamp, errTimestamp := strconv.ParseInt(strings.Trim(parts[3], " "), 10, 64)
	if errTimestamp != nil {
		return "", nil, errors.New(fmt.Sprintf("line %s does not have a valid timestamp", line))
	}
	return strings.Trim(parts[0], " "), &models.Point{Latitude: lat, Longitude: lon, Timestamp: timestamp}, nil
}
