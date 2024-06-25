package middleware

import (
	"net/http"
)

// AuthorizeAccess is a middleware to enforce function's access control of the requester
//
//	this method should be called after SessionCtx()
func (rs *Resource) AuthorizeAccess(resourceCode, act string) func(next http.Handler) http.Handler {
	return rs.utility.AuthorizeAccess(resourceCode, act)
}
