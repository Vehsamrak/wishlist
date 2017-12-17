package src

type WishService struct {

}

func (wishService *WishService) Start() (chan bool, error) {
    return make(chan bool), nil
}

func (wishService *WishService) Stop() bool {
    return true
}
