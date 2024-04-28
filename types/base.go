package types

type BaseFilterRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type PaginationDTO struct {
	TotalPages       int   `json:"totalPages"`
	TotalElements    int64 `json:"totalElements"`
	NumberOfElements int   `json:"numberOfElements"`
}

type BaseRelationDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
