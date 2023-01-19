package service

import (
	"douyin/biz/dao"
)

var Svr, err = New(dao.New())

type Service struct {
	dao dao.Dao
}

func New(d dao.Dao) (s *Service, err error) {
	s = &Service{
		dao: d,
	}
	return
}
