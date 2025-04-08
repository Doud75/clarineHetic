package dto

type ProfileSearchRequest struct {
    SearchTerm string `json:"search_term" binding:"required"`
}
