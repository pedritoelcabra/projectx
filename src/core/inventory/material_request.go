package inventory

import "github.com/pedritoelcabra/projectx/src/core/defs"

type RequestList []*ResourceRequest

type ResourceRequest struct {
	Type   int
	Amount int
}

func (r *ResourceRequest) Clone() *ResourceRequest {
	newCopy := r
	return newCopy
}

func NewRequest(amount, material int) *ResourceRequest {
	request := &ResourceRequest{}
	request.Type = material
	request.Amount = amount
	return request
}

func FromRequirement(requirement *defs.ResourceRequirement) *ResourceRequest {
	request := &ResourceRequest{}
	request.Type = defs.GetMaterialKeyByName(requirement.Type)
	request.Amount = requirement.Amount
	return request
}

func CopyRequestList(template []*ResourceRequest, target []*ResourceRequest) {
	for _, request := range template {
		target = append(target, request.Clone())
	}
}

func CopyResourceRequirements(template []*defs.ResourceRequirement, target RequestList) RequestList {
	for _, requirement := range template {
		target = append(target, FromRequirement(requirement))
	}
	return target
}

func FulfillResourceRequests(inventory *Inventory, target RequestList) RequestList {
	revisedList := RequestList{}
	for _, request := range target {
		amountFulfilled := inventory.RemoveItems(request.Type, request.Amount)
		if amountFulfilled == request.Amount {
			continue
		}
		revisedList = append(revisedList, NewRequest(request.Amount-amountFulfilled, request.Type))
	}
	return revisedList
}
