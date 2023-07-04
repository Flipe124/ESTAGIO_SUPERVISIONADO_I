package process

import (
	// Internal imports.
	_ "backend/internal/core/auth"
	_ "backend/internal/core/user"
	_ "backend/internal/infra/db"
	_ "backend/internal/models"
	_ "backend/internal/server"

	// Public imports.
	_ "backend/pkg/docs"
)
