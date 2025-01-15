package types

type MapStorage map[int]float64
type StoreType string

var (
	OBUKafkaTopic    string    = "obudata"
	MemeoryStoreType StoreType = "memory"
)

type Invoice struct {
	OBUID         int     `json:"obuID"`
	TotalAmount   float64 `json:"totalAmount"`
	TotalDistance float64 `json:"totalDistance"`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type OBUData struct {
	OBUID     int `json:"obuID"`
	RequestID int `json:"requestID"`
	Point
}
