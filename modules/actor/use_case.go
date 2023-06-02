package actor

import (
	"crm_service/entity"
	"crm_service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseActorInterface interface {
	CreateActor(actor ActorBody) (entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	GetAllActor() ([]entity.Actor, error)
	UpdateActorById(id uint, actor UpdateActorBody) (entity.Actor, error)
	DeleteActorById(id uint) error
}

type actorUseCaseStruct struct {
	actorRepository repository.ActorRepoInterface
}

func (uc actorUseCaseStruct) CreateActor(actor ActorBody) (entity.Actor, error) {
	var NewActor *entity.Actor

	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(actor.Password), 12)
	NewActor = &entity.Actor{
		Username: actor.Username,
		Password: string(hashingPassword),
	}

	_, err := uc.actorRepository.CreateActor(NewActor)
	if err != nil {
		return *NewActor, err
	}
	return *NewActor, nil
}

func (uc actorUseCaseStruct) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	actor, err := uc.actorRepository.GetActorById(id)
	return actor, err
}

func (uc actorUseCaseStruct) GetAllActor() ([]entity.Actor, error) {
	var actor []entity.Actor
	actor, err := uc.actorRepository.GetAllActor()
	return actor, err
}

func (uc actorUseCaseStruct) UpdateActorById(id uint, actor UpdateActorBody) (entity.Actor, error) {
	var NewActor *entity.Actor

	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(actor.Password), 12)
	NewActor = &entity.Actor{
		Username: actor.Username,
		Password: string(hashingPassword),
		Verified: actor.Verified,
		Active:   actor.Active,
	}
	_, err := uc.actorRepository.UpdateActorById(id, NewActor)
	if err != nil {
		return *NewActor, err
	}
	return *NewActor, nil
}

func (uc actorUseCaseStruct) DeleteActorById(id uint) error {
	err := uc.actorRepository.DeleteActorById(id)
	return err

}
