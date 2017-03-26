package usecase

import "github.com/jtreu/sleeplog-app/server/domain"

type MediaInteractor struct {
	MediaRepository domain.MediaRepository
}

func NewMediaInteractor() *MediaInteractor {
	return &MediaInteractor{}
}
