package mon_test

import (
	"github.com/applike/gosoline/pkg/mon"
	"github.com/applike/gosoline/pkg/mon/mocks"
	"github.com/jonboulle/clockwork"
	"testing"
	"time"
)

func TestSamplingLogger_Infof(t *testing.T) {
	mock := new(mocks.Logger)
	mock.On("Info", "this should be logged").Once()
	mock.On("Infof", "log msg", "a", 4).Twice()

	clock := clockwork.NewFakeClock()
	logger := mon.NewSamplingLoggerWithInterfaces(mock, clock, time.Minute)

	logger.Infof("log msg", "a", 4)
	logger.Infof("log msg", "a", 4)
	logger.Info("this should be logged")

	clock.Advance(time.Second)
	logger.Infof("log msg", "a", 4)

	clock.Advance(time.Hour)
	logger.Infof("log msg", "a", 4)

	mock.AssertExpectations(t)
}
