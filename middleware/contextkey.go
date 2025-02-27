package middleware

// for using context value, saved in utils so every key come from 1 place
type ContextKey string

const (
	UserKey        = ContextKey("user")
	CompanyCodeKey = ContextKey("companyCode")
	SessionKey     = ContextKey("session")
	StoreKey       = ContextKey("store")
)
