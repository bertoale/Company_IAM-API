package permission

func ToPermissionResponse(p Permission) PermissionResponse {
	return PermissionResponse{
		ID:          p.ID,
		Code:        p.Code,
		Description: p.Description,
	}
}
