package gomock

import (
	"encoding/json"
	"fmt"
	. "github.com/golang/mock/gomock"
	"testing"
)

func testRepository(repository Repository) {
	_, err := repository.Retrieve("123")
	fmt.Println(err)
	repository.Create("123", []byte("456"))
	resp, err := repository.Retrieve("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func TestMockRepository(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	obj := struct {
		X, Y int
	}{}
	objBytes, _ := json.Marshal(obj)

	// 假设有这样一个场景：先Retrieve领域对象失败，然后Create领域对象成功，再次Retrieve领域对象就能成功。
	// 这个场景对应的mock对象的行为注入代码如下所示
	mockRepo.EXPECT().Retrieve(Any()).Return(nil, fmt.Errorf("can not Retrieve"))
	mockRepo.EXPECT().Create(Any(), Any()).Return(nil)
	mockRepo.EXPECT().Retrieve(Any()).Return(objBytes, nil)

	testRepository(mockRepo)
}
