package services

import (
"github.com/cooladdr/xiuxiubeidanci/datamodels"
"github.com/cooladdr/xiuxiubeidanci/repositories"
)
	

type WordService interface{
	Find(word string) datamodels.Word
}


type wordService struct{
	repo repositories.WordRepository
}

func NewWordService(repo repositories.WordRepository) WordService {
	return &wordService{
		repo:repo,
	}
}


func (ws *wordService) Find(word string) (w datamodels.Word) {
	return ws.repo.Find(word)
}


