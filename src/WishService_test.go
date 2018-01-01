package src

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/vehsamrak/resttest"
    "net/http"
    "testing"
    "net/url"
)

func TestWishService(test *testing.T) {
    suite.Run(test, new(WishServiceTest))
}

type WishServiceTest struct {
    suite.Suite
}

func (suite *WishServiceTest) TestIndexHandler_httpGetNonexistentRoute_404CodeReturned() {
    service := &WishService{}

    response := resttest.RequestGet(service.IndexHandler, "/nonexistent-route", nil)

    assert.Equal(suite.T(), http.StatusNotFound, response.Code)
    assert.Equal(suite.T(), "{\"data\":\"Url not found\"}", response.Body.String())
}

func (suite *WishServiceTest) TestIndexHandler_httpGetIndexRoute_404CodeReturned() {
    service := &WishService{}

    response := resttest.RequestGet(service.IndexHandler, "/", nil)

    assert.Equal(suite.T(), http.StatusNotFound, response.Code)
    assert.Equal(suite.T(), "{\"data\":\"Url not found\"}", response.Body.String())
}

func (suite *WishServiceTest) TestWishesHandler_httpGetWishesRouteWithoutUserId_400() {
    service := &WishService{}

    response := resttest.RequestGet(service.WishesHandler, "/wishes", nil)

    assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
}

func (suite *WishServiceTest) TestWishesHandler_httpGetWishesRouteWithUserId_200() {
    service := &WishService{}
    parameters := url.Values{}
    parameters.Set("userId", "1")

    response := resttest.RequestGet(service.WishesHandler, "/wishes?" + parameters.Encode(), nil)

    assert.Equal(suite.T(), http.StatusOK, response.Code)
}
