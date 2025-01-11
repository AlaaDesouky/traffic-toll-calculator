package types

var OBUKafkaTopic = "obudata"

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type OBUData struct {
	OBUID     int `json:"obuID"`
	RequestID int `json:"requestID"`
	Point
}
