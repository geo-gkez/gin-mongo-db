package service

import (
	"geo.org/gin-members/models"
	"geo.org/gin-members/repository"
)

func GetMembers() (members []models.Member) {
	return repository.FindAllMembers()
}

func CreateMember(member models.Member) (newMember models.Member) {
	return repository.InsertMember(member)
}
