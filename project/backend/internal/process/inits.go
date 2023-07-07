package process

import (
	// Imports of all packages project.
	_ "backend/internal/core/account"
	_ "backend/internal/core/auth"
	_ "backend/internal/core/category"
	_ "backend/internal/core/status"
	_ "backend/internal/core/type"
	_ "backend/internal/core/user"
	_ "backend/internal/infra/db"
	_ "backend/internal/models"
	_ "backend/internal/server"
	_ "backend/pkg/docs"
)
