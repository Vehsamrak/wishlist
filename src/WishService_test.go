package src

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
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

    result := wishService.Start()

    assert.Equal(suite.T(), result, true)
}
