package sitegen

import "os"

type Service struct {
	hugo *os.Process
}

func New(hugo *os.Process) *Service {
	return &Service{}
}
