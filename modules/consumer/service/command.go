package service

import (
	"mime/multipart"
	"xyz/modules/consumer/domain"
	"xyz/modules/consumer/repository"
	"xyz/packages/util"
	"xyz/packages/validator"
)

type consumerCommandService struct {
	consumerCommandRepository repository.ConsumerCommandRepositoryInterface
	consumerQueryRepository    repository.ConsumerQueryRepositoryInterface
}

func NewConsumerCommandService(
	consumerCommandRepository repository.ConsumerCommandRepositoryInterface,
	consumerQueryRepository repository.ConsumerQueryRepositoryInterface,
) ConsumerCommandServiceInterface {
	return &consumerCommandService{
		consumerCommandRepository: consumerCommandRepository,
		consumerQueryRepository:    consumerQueryRepository,
	}
}

func (s *consumerCommandService) CreateConsumer(consumer domain.Consumer, photoKTP, photoSelfie *multipart.FileHeader) (domain.Consumer, error) {

	consumerPhotoPath, err := util.SaveFile(photoKTP)
	if err != nil {
		return domain.Consumer{}, err
	}

	consumerSelfiePath, err := util.SaveFile(photoSelfie)
	if err != nil {
		return domain.Consumer{}, err
	}

	consumer.Photo_KTP = consumerPhotoPath
	consumer.Photo_Selfie = consumerSelfiePath

	consumerDomain, err := s.consumerCommandRepository.CreateConsumer(consumer)
	if err != nil {
		return domain.Consumer{}, err
	}

	return consumerDomain, nil
}

func (s *consumerCommandService) UpdateConsumer(consumer domain.Consumer, photoKTP, photoSelfie *multipart.FileHeader) (domain.Consumer, error) {
	validateEmpty := validator.CheckEmpty(consumer.Full_Name, consumer.Legal_Name, consumer.Place_Of_Birth, consumer.Date_Of_Birth, consumer.NIK, consumer.Salary, photoKTP, photoSelfie)
	if validateEmpty != nil {
		return domain.Consumer{}, validateEmpty
	}

	consumerDomain, err := s.consumerQueryRepository.GetConsumerByID(consumer.ID)
	if err != nil {
		return domain.Consumer{}, err
	}

	if photoKTP != nil {
		consumerPhotoPath, err := util.SaveFile(photoKTP)
		if err != nil {
			return domain.Consumer{}, err
		}
		consumerDomain.Photo_KTP = consumerPhotoPath
	}
	if photoSelfie != nil {
		consumerSelfiePath, err := util.SaveFile(photoSelfie)
		if err != nil {
			return domain.Consumer{}, err
		}
		consumerDomain.Photo_Selfie = consumerSelfiePath
	}

	consumerData, err := s.consumerCommandRepository.UpdateConsumer(consumerDomain)
	if err != nil {
		return domain.Consumer{}, err
	}

	return consumerData, nil
}