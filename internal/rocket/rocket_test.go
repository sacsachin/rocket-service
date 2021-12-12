package rocket

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Test get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		id := "UUID-1"
		rocketStoreMock.EXPECT().GetRocketByID(id).Return(Rocket{
			ID: id,
		}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("Test insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(Rocket{
				ID: id,
			}).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)

		rkt, err := rocketService.InsertRocket(
			context.Background(),
			Rocket{
				ID: id,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("test delete rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)

		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocket(id).
			Return(nil)

		rocketService := New(rocketStoreMock)

		err := rocketService.DeleteRocket(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		ioutil.ReadFile("abc.txt")
	})
}
