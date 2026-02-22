package rbac

import (
	"encoding/json"
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

func (s *Service) HasPermission(roleNames []string, required string) (bool, error) {

	for _, role := range roleNames {

		cacheKey := "role:" + role

		// 1️⃣ cek cache
		cached, err := s.cache.Get(cacheKey)
		if err == nil && cached != "" {

			var perms []string
			if err := json.Unmarshal([]byte(cached), &perms); err == nil {
				for _, p := range perms {
					if p == required {
						return true, nil
					}
				}
			}

			// kalau tidak ketemu permission di cache, lanjut role berikutnya
			continue
		}

		// 2️⃣ kalau tidak ada di cache → ambil per role
		perms, err := s.repo.GetPermissionsByRoleNames([]string{role})
		if err != nil {
			return false, err
		}

		// simpan ke cache
		bytes, _ := json.Marshal(perms)
		_ = s.cache.Set(cacheKey, string(bytes), 10*time.Minute)

		for _, p := range perms {
			if p == required {
				return true, nil
			}
		}
	}

	return false, nil
}