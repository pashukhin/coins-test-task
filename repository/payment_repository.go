package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pashukhin/coins-test-task/entity"
)

type PaymentRepository interface{
	GetAll() (all []*entity.Payment, err error)
	GetIncomingFor(acc *entity.Account) (list []*entity.Payment, err error)
	GetOutgoingFor(acc *entity.Account) (list []*entity.Payment, err error)
	GetAllFor(acc *entity.Account) (list []*entity.Payment, err error)
	Get(id int64) (payment *entity.Payment, err error)
	Create(from, to *entity.Account, amount float64) (payment *entity.Payment, err error)
	Delete(payment *entity.Payment) error
}

func NewPaymentRepository(db *sqlx.DB) PaymentRepository {
	return &paymentRepository{&repository{db}}
}

type paymentRepository struct {
	*repository
}

func (p paymentRepository) GetAll() (all []*entity.Payment, err error) {
	err = p.db.Select(&all, "select * from payment order by created")
	return
}

func (p paymentRepository) GetIncomingFor(acc *entity.Account) (list []*entity.Payment, err error) {
	err = p.db.Select(&list, "select * from payment where account_to_id = $1 order by created", acc.ID)
	return
}

func (p paymentRepository) GetOutgoingFor(acc *entity.Account) (list []*entity.Payment, err error) {
	err = p.db.Select(&list, "select * from payment where account_from_id = $1 order by created", acc.ID)
	return
}

func (p paymentRepository) GetAllFor(acc *entity.Account) (list []*entity.Payment, err error) {
	err = p.db.Select(&list, "select * from payment where account_from_id = $1 or account_to_id = $1 order by created", acc.ID)
	return
}

func (p paymentRepository) Get(id int64) (payment *entity.Payment, err error) {
	payment = &entity.Payment{}
	err = p.db.Get(payment, "select * from payment where id = $1", id)
	return
}

func (p paymentRepository) Create(from, to *entity.Account, amount float64) (payment *entity.Payment, err error) {
	payment = &entity.Payment{}
	sql := "insert into payment (account_from_id, account_to_id, amount) values ($1, $2, $3) returning *"
	err = p.db.Get(payment, sql, from.ID, to.ID, amount)
	return
}

func (p paymentRepository) Delete(payment *entity.Payment) error {
	return p.ExecForOne("delete from payment where id = $1", payment.ID)
}
