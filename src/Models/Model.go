package Models

type AstarData struct {
	NumNodes       int    `json:"numnodes"`
	SourceNode     int    `json:"sourcenode"`
	DestNode       int    `json:"destnode"`
	MatrixRelation string `json:"matrixrelation"`
	Coordinates    string `json:"coordinates"`
}
