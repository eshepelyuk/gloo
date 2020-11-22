// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/rbac/rbac.proto

package rbac

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
)

// Equal function
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetRequireRbac() != target.GetRequireRbac() {
		return false
	}

	return true
}

// Equal function
func (m *ExtensionSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtensionSettings)
	if !ok {
		that2, ok := that.(ExtensionSettings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetDisable() != target.GetDisable() {
		return false
	}

	if len(m.GetPolicies()) != len(target.GetPolicies()) {
		return false
	}
	for k, v := range m.GetPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicies()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicies()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Policy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPrincipals()) != len(target.GetPrincipals()) {
		return false
	}
	for idx, v := range m.GetPrincipals() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPrincipals()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPrincipals()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetPermissions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPermissions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPermissions(), target.GetPermissions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Principal) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Principal)
	if !ok {
		that2, ok := that.(Principal)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetJwtPrincipal()).(equality.Equalizer); ok {
		if !h.Equal(target.GetJwtPrincipal()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetJwtPrincipal(), target.GetJwtPrincipal()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *JWTPrincipal) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JWTPrincipal)
	if !ok {
		that2, ok := that.(JWTPrincipal)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetClaims()) != len(target.GetClaims()) {
		return false
	}
	for k, v := range m.GetClaims() {

		if strings.Compare(v, target.GetClaims()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetProvider(), target.GetProvider()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Permissions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Permissions)
	if !ok {
		that2, ok := that.(Permissions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetPathPrefix(), target.GetPathPrefix()) != 0 {
		return false
	}

	if len(m.GetMethods()) != len(target.GetMethods()) {
		return false
	}
	for idx, v := range m.GetMethods() {

		if strings.Compare(v, target.GetMethods()[idx]) != 0 {
			return false
		}

	}

	return true
}
