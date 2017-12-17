package src

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"
)

func TestWishService(test *testing.T) {
    suite.Run(test, new(HelpWishService))
}

type HelpWishService struct {
    suite.Suite
}

func (suite *HelpWishService) SetupTest() {}

func (suite *HelpWishService) TestGetNames_emptyParameters_stringReturned() {
    wishService := &WishService{}

    _, status := wishService.Start()
    wishService.Stop()

    assert.Nil(suite.T(), status)
}
