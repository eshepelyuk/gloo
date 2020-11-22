// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/failover.proto

package v1

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
func (m *Failover) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Failover)
	if !ok {
		that2, ok := that.(Failover)
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

	if len(m.GetPrioritizedLocalities()) != len(target.GetPrioritizedLocalities()) {
		return false
	}
	for idx, v := range m.GetPrioritizedLocalities() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPrioritizedLocalities()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPrioritizedLocalities()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *LocalityLbEndpoints) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LocalityLbEndpoints)
	if !ok {
		that2, ok := that.(LocalityLbEndpoints)
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

	if h, ok := interface{}(m.GetLocality()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLocality()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLocality(), target.GetLocality()) {
			return false
		}
	}

	if len(m.GetLbEndpoints()) != len(target.GetLbEndpoints()) {
		return false
	}
	for idx, v := range m.GetLbEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLbEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLbEndpoints()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLoadBalancingWeight()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLoadBalancingWeight(), target.GetLoadBalancingWeight()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *LbEndpoint) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LbEndpoint)
	if !ok {
		that2, ok := that.(LbEndpoint)
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

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if h, ok := interface{}(m.GetHealthCheckConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthCheckConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthCheckConfig(), target.GetHealthCheckConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUpstreamSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamSslConfig(), target.GetUpstreamSslConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLoadBalancingWeight()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLoadBalancingWeight(), target.GetLoadBalancingWeight()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Locality) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Locality)
	if !ok {
		that2, ok := that.(Locality)
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

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetZone(), target.GetZone()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubZone(), target.GetSubZone()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Failover_PrioritizedLocality) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Failover_PrioritizedLocality)
	if !ok {
		that2, ok := that.(Failover_PrioritizedLocality)
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

	if len(m.GetLocalityEndpoints()) != len(target.GetLocalityEndpoints()) {
		return false
	}
	for idx, v := range m.GetLocalityEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalityEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLocalityEndpoints()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *LbEndpoint_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LbEndpoint_HealthCheckConfig)
	if !ok {
		that2, ok := that.(LbEndpoint_HealthCheckConfig)
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

	if m.GetPortValue() != target.GetPortValue() {
		return false
	}

	if strings.Compare(m.GetHostname(), target.GetHostname()) != 0 {
		return false
	}

	return true
}
