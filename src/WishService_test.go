package src

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/vehsamrak/resttest"
    "net/http"
    "testing"
)

func TestWishService(test *testing.T) {
    suite.Run(test, new(WishServiceTest))
}

type WishServiceTest struct {
    suite.Suite
}

func (suite *WishServiceTest) SetupTest() {}

func (suite *WishServiceTest) TestIndexHandler_httpGetNonexistentRoute_404CodeReturned() {
    service := &WishService{}

    response := resttest.RequestGet(service.IndexHandler, "/nonexistent-route", nil)

    assert.Equal(suite.T(), http.StatusNotFound, response.Code)
    assert.Equal(suite.T(), "{\"message\":\"Url not found\"}", response.Body.String())
}

func (suite *WishServiceTest) TestIndexHandler_httpGetIndexRoute_200CodeReturned() {
    service := &WishService{}

    response := resttest.RequestGet(service.IndexHandler, "/", nil)

    assert.Equal(suite.T(), http.StatusOK, response.Code)
    assert.Equal(suite.T(), "{\"message\":\"Url found\"}", response.Body.String())
}

func (suite *WishServiceTest) TestWishesHandler_httpGetIndexRoute_200() {
    service := &WishService{}

    response := resttest.RequestGet(service.WishesHandler, "/wishes", nil)

    assert.Equal(suite.T(), http.StatusNotFound, response.Code)
}
