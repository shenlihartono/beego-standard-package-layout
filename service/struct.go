// Package service is for service/operation on domain types.
package service

import (
	groot "beego-standard-layout"
	"github.com/astaxie/beego/orm"
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

func (s StructService) Struct(ID string) Response {
	str, err := s.repo.Struct(ID)
	if err != nil {
		if err == orm.ErrNoRows {
			return NewResponse(StatusNotFound, "struct not found")
		}

		return NewResponse(StatusInternalServerError, "internal server error")
	}

	return NewResponse(StatusOK, str)
}

func (s StructService) Structs() Response {
	str, err := s.repo.Structs()
	if err != nil {
		if err == orm.ErrNoRows {
			return NewResponse(StatusNotFound, "struct not found")
		}

		return NewResponse(StatusInternalServerError, "internal server error")
	}

	return NewResponse(StatusOK, str)
}

func (s StructService) UpdateStruct(ID string, request groot.StructRequest) Response {
	str, err := s.repo.Struct(ID)
	if err != nil {
		if err == orm.ErrNoRows {
			return NewResponse(StatusNotFound, "struct not found")
		}

		return NewResponse(StatusInternalServerError, "internal server error")
	}

	str.Value = request.Value
	err = s.repo.Update(str)
	if err != nil {
		return NewResponse(StatusInternalServerError, "internal server error")
	}

	return NewResponse(StatusOK, str)
}

func (s StructService) Delete(ID string) Response {
	str, err := s.repo.Struct(ID)
	if err != nil {
		if err == orm.ErrNoRows {
			return NewResponse(StatusNotFound, "struct not found")
		}

		return NewResponse(StatusInternalServerError, "internal server error")
	}

	err = s.repo.Delete(str)
	if err != nil {
		return NewResponse(StatusInternalServerError, "internal server error")
	}

	return NewResponse(StatusOK, "success delete struct")
}
