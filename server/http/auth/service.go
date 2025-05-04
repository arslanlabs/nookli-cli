// server/http/auth/service.go
package auth

import (
	internal "nookli/internal/auth"
	svc "nookli/pkg/service/auth"
)

// svcAuth is the singleton auth service used by handler & middleware.
var authSvc svc.Service = internal.NewSupabaseService()
