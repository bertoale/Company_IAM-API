package rbac

import (
	"encoding/json"
	"fmt"
	"time"
)

type Cache interface {
	Set(key, value string, ttl time.Duration) error 
	Get (key string) (string,error)
}

type Service struct {
	repo Repository
	cache Cache
}

func NewService(repo Repository, cache Cache) *Service {
	return &Service{
		repo: repo,
		cache: cache,
	}
}

func (s *Service) GetUserPermissions(userID uint) (map[string]struct{}, error) {

	cacheKey := fmt.Sprintf("user:%d:permissions", userID)

	// 1️⃣ cek cache
	cached, err := s.cache.Get(cacheKey)
	if err == nil && cached != "" {

		var perms []string
		if err := json.Unmarshal([]byte(cached), &perms); err == nil {

			permMap := make(map[string]struct{}, len(perms))
			for _, p := range perms {
				permMap[p] = struct{}{}
			}

			return permMap, nil
		}
	}

	// 2️⃣ ambil dari DB (SEMUA SEKALIGUS)
	perms, err := s.repo.GetPermissionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 3️⃣ simpan ke cache
	bytes, _ := json.Marshal(perms)
	_ = s.cache.Set(cacheKey, string(bytes), 10*time.Minute)

	// 4️⃣ convert ke map
	permMap := make(map[string]struct{}, len(perms))
	for _, p := range perms {
		permMap[p] = struct{}{}
	}

	return permMap, nil
}
func (s *Service) HasPermission(userID uint, required string) (bool, error) {
	perms, err := s.GetUserPermissions(userID)
	if err != nil {
		return false, err
	}

	_, ok := perms[required]
	return ok, nil
}