package application

func ToApplicationResponse(app Application) ApplicationResponse {
	return ApplicationResponse{
		ID:       app.ID,
		Code:     app.Code,
		Name:     app.Name,
		IsActive: app.IsActive,
	}
}
