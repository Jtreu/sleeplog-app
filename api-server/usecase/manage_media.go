package usecase

import "github.com/jtreu/sleeplog-app/api-server/domain"

type MediaInteractor struct {
	MediaRepository domain.MediaRepository
}

func NewMediaInteractor() *MediaInteractor {
	return &MediaInteractor{}
}
