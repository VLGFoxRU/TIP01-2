package order

import ( 
	"errors"
	"fmt"
)

var ErrOrderNotFound = errors.New("order not found")

type Repo struct {
	data map[int64]Order
}

func NewRepo() *Repo {
	return &Repo{
		data: map[int64]Order{
			101: {ID: 101, UserID: 1, Item: "Ноутбук", Price: 79990},
			102: {ID: 102, UserID: 2, Item: "Мышь", Price: 2490},
			103: {ID: 103, UserID: 1, Item: "Клавиатура", Price: 5990},
		},
	}
}

func (r *Repo) GetByID(id int64) (Order, error) {
	o, ok := r.data[id]
	if !ok {
		return Order{}, ErrOrderNotFound
	}
	return o, nil
}

// GetAllByUserID возвращает все заказы указанного пользователя.
func (r *Repo) GetAllByUserID(userID int64) ([]Order, error) {
	var result []Order
	for _, o := range r.data {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("У пользователя %d нет заказов", userID)
	}
	return result, nil
}