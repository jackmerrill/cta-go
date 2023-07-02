package ctago

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var timestampString string
	err := json.Unmarshal(data, &timestampString)
	if err != nil {
		return err
	}

	// Define the layout of the timestamp based on the format you have.
	layout := "2006-01-02T15:04:05"

	// Parse the timestamp string into a time.Time value.
	parsedTime, err := time.Parse(layout, timestampString)
	if err != nil {
		return err
	}

	// Assign the parsed time to the CustomTime field.
	ct.Time = parsedTime
	return nil
}

type CustomFloat64 float64

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	var value interface{}
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if string(data) == "null" {
		*cf = 0
		return nil
	}

	switch v := value.(type) {
	case string:
		parsedValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*cf = CustomFloat64(parsedValue)
	case float64:
		*cf = CustomFloat64(v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}

type ArrivalsService struct {
	http *NetworkClient
}

type ArrivalResponse struct {
	ArrivalBase ArrivalBase `json:"ctatt"`
}

type ArrivalBase struct {
	Timestamp CustomTime  `json:"tmst"`
	ErrorCode string      `json:"errCd"`
	ErrorMsg  interface{} `json:"errNm"`
	Trains    []Eta       `json:"eta"`
}

type Eta struct {
	StationID       string        `json:"staId"`
	StopID          string        `json:"stpId"`
	StationName     string        `json:"staNm"`
	StopDescription string        `json:"stpDe"`
	Run             string        `json:"rn"`
	RouteName       string        `json:"rt"`
	DestinationStop string        `json:"destSt"`
	DestinationName string        `json:"destNm"`
	TrainDirection  string        `json:"trDr"`
	Generated       CustomTime    `json:"prdt"`
	ArriveDepart    CustomTime    `json:"arrT"`
	Approaching     string        `json:"isApp"`
	Scheduled       string        `json:"isSch"`
	Delayed         string        `json:"isDly"`
	Fault           string        `json:"isFlt"`
	Flags           interface{}   `json:"flags"`
	Latitude        CustomFloat64 `json:"lat"`
	Longitude       CustomFloat64 `json:"lon"`
	Heading         string        `json:"heading"`
}

func (s *ArrivalsService) Get(StationID int) (*ArrivalBase, error) {
	var out ArrivalResponse

	uri := s.http.baseURL

	uri.Path = fmt.Sprintf("%s/ttarrivals.aspx", uri.Path)
	q := uri.Query()
	q.Add("mapid", fmt.Sprintf("%d", StationID))

	uri.RawQuery = q.Encode()

	err := s.http.Get(uri, &out)

	if err != nil {
		return nil, err
	}

	return &out.ArrivalBase, nil
}
