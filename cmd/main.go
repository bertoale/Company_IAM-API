package main

import (
	"company_iam/internal/application"
	"company_iam/internal/auth"
	"company_iam/internal/permission"
	"company_iam/internal/rbac"
	"company_iam/internal/role"
	"company_iam/internal/role_permission"
	"company_iam/internal/user"
	"company_iam/internal/user_application"
	"company_iam/internal/user_role"

	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"
	"company_iam/pkg/redis"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %d - %s %s %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.StatusCode,
			param.Method,
			param.Path,
			param.Latency,
		)
	}))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.CorsOrigin},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}, AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	rateLimiter := middlewares.NewRateLimiter(20, time.Minute) // 20 request per menit
	r.Use(rateLimiter.Middleware())

	// === Static Files untuk Upload ===
	r.Static("/uploads", "./uploads")

	// === Database ===
	if err := config.Connect(cfg); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// === Migrate Database ===
	db := config.GetDB()
	tables := []interface{}{
		&user.User{},
		&permission.Permission{},
		&role.Role{},
		&application.Application{},
		&user_role.UserRole{},
		&role_permission.RolePermission{},
		&application.Application{},
		&user_application.UserApplication{},
	}
	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("✅ Migrasi database berhasil.")

	// === Home Route ===
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "Welcome to the REST API",
			"version":   "1.0.0",
			"timestamp": time.Now(),
		})
	})

	r.Use(middlewares.GinErrorHandler())
	redisClient, err := redis.NewRedis(cfg.RedisAddr, cfg.RedisPass, cfg.RedisDB)
if err != nil {
	log.Fatalf("Redis not available: %v", err)
}

	//seeder
	user.SeedAdminUser()

	rbacRepo := rbac.NewRepository(db)
	rbacService := rbac.NewService(rbacRepo, redisClient)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userController := user.NewController(userService)
	user.SetupRoutes(r, userController, cfg, rbacService)

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg)
	authController := auth.NewController(authService, cfg)
	auth.SetupRoutes(r, authController, cfg)


	roleRepo := role.NewRepository(db)
	roleService := role.NewService(roleRepo)
	roleController := role.NewController(roleService)
	role.SetupRoutes(r, roleController, cfg, rbacService)

	permissionRepo := permission.NewRepository(db)
	permissionService := permission.NewService(permissionRepo)
	permissionController := permission.NewController(permissionService)
	permission.SetupRoutes(r, permissionController, cfg, rbacService)

	applicationRepo := application.NewRepository(db)
	applicationService := application.NewService(applicationRepo)
	applicationController := application.NewController(applicationService)
	application.SetupRoutes(r, applicationController, cfg, rbacService)

	userRoleRepo := user_role.NewRepository(db)
	userRoleService := user_role.NewService(userRoleRepo)
	userRoleController := user_role.NewController(userRoleService)
	user_role.SetupRoutes(r, userRoleController, cfg, rbacService)

	rolePermissionRepo := role_permission.NewRepository(db)
	rolePermissionService := role_permission.NewService(rolePermissionRepo)
	rolePermissionController := role_permission.NewController(rolePermissionService)
	role_permission.SetupRoutes(r, rolePermissionController, cfg, rbacService)

	userApplicationRepo := user_application.NewRepository(db)
	userApplicationService := user_application.NewService(userApplicationRepo)
	userApplicationController := user_application.NewController(userApplicationService)
	user_application.SetupRoutes(r, userApplicationController, cfg, rbacService)

	// 404 Not Found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})

	// === Start Server ===
	log.Printf("Server running on port %s", cfg.Port)
	log.Printf("Local: http://localhost:%s", cfg.Port)
	log.Printf("Environment: %s", cfg.NodeEnv)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}

}
