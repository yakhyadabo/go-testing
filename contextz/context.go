package contextz

import (
	"context"
	"time"
)

type UserRequest struct {

}

func (request UserRequest) fitchThirdPartyStuffThatCanBeVerySlow() (int, error){
	time.Sleep(time.Second*10)
	return 0, nil
}

func (request UserRequest) FetchUserData(ctx context.Context)(int, error){
	val, err := request.fitchThirdPartyStuffThatCanBeVerySlow()
	if err != nil {
		return 0, err
	}
	return val,nil
}