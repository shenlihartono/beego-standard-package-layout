// Package service is for service/operation on domain types.
package service

import (
	groot "beego-standard-layout"
)

type StructService struct {
	repo groot.StructRepository
}

func NewStructService(repo groot.StructRepository) StructService {
	return StructService{repo: repo}
}

func (s StructService) CreateStruct(request groot.StructRequest) Response {
	str, err := s.repo.Create(request)
	if err != nil {
		return NewResponse(StatusInternalServerError, "internal server error")
	}

	return NewResponse(StatusOK, str)
}

//func (s StructService) Struct() Response {
//
//}