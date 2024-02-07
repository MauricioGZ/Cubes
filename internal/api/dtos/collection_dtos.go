package dtos

type RemoveCubeFromCollection struct {
	CubeID int64 `query:"id"` //only works with GET or DELETE requests
}

type AddCubeToCollection struct {
	CubeID int64 `json:"cube_id"`
}
