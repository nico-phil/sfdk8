package mid

// Authenticate validates authentication via the auth service.
// func Authenticate(ctx context.Context, log *logger.Logger, client *authclient.Client, authorization string, handler Handler) error {
// 	resp, err := client.Authenticate(ctx, authorization)
// 	if err != nil {
// 		return errs.New(errs.Unauthenticated, err)
// 	}

// 	ctx = setUserID(ctx, resp.UserID)
// 	ctx = setClaims(ctx, resp.Claims)

// 	return handler(ctx)
// }

// Bearer processes JWT authentication logic.
// func Bearer(ctx context.Context, ath *auth.Auth, authorization string, handler Handler) error {
// 	claims, err := ath.Authenticate(ctx, authorization)
// 	if err != nil {
// 		return errs.New(errs.Unauthenticated, err)
// 	}

// 	if claims.Subject == "" {
// 		return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, no claims")
// 	}

// 	subjectID, err := uuid.Parse(claims.Subject)
// 	if err != nil {
// 		return errs.New(errs.Unauthenticated, fmt.Errorf("parsing subject: %w", err))
// 	}

// 	ctx = setUserID(ctx, subjectID)
// 	ctx = setClaims(ctx, claims)

// 	return handler(ctx)
// }
