package inventory

import "github.com/pedritoelcabra/projectx/src/core/defs"

type ResourceRequest struct {
	Type   string
	Amount int
}

func (r *ResourceRequest) Clone() *ResourceRequest {
	newCopy := r
	return newCopy
}

func FromRequirement(requirement *defs.ResourceRequirement) *ResourceRequest {
	request := &ResourceRequest{}
	request.Type = requirement.Type
	request.Amount = requirement.Amount
	return request
}

func CopyRequestList(template []*ResourceRequest, target []*ResourceRequest) {
	for _, request := range template {
		target = append(target, request.Clone())
	}
}

func CopyResourceRequirements(template []*defs.ResourceRequirement, target []*ResourceRequest) []*ResourceRequest {
	for _, requirement := range template {
		target = append(target, FromRequirement(requirement))
	}
	return target
}
