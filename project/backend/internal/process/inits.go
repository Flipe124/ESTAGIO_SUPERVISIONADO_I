package process

import (
	// Imports of all packages project.
	_ "backend/internal/core/address"
	_ "backend/internal/core/auth"
	_ "backend/internal/core/client"
	_ "backend/internal/core/order"
	_ "backend/internal/core/service"
	_ "backend/internal/core/user"
	_ "backend/internal/infra/db"
	_ "backend/internal/models"
	_ "backend/internal/server"
	_ "backend/pkg/docs"
)
