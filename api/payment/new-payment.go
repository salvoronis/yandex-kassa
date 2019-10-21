package payment

import (
	"github.com/dbzer0/yandex-kassa/api/client"
)

type NewPayment struct {
	APIClient    *client.APIClient `json:"-"`
	Amount       NewAmount         `json:"amount"`                        // сумма платежа
	Description  *string           `json:"description,omitempty"`         // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient    *NewRecipient     `json:"recipient,omitempty"`           // получатель платежа
	MethodData   *MethodData       `json:"payment_method_data,omitempty"` // данные для оплаты конкретным способом  (payment_method)
	Confirmation *Confirmation     `json:"confirmation,omitempty"`        // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
}

type NewRecipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

type NewAmount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

// Create создает платеж. Он содержит всю необходимую информацию для проведения о
// платы (сумму, валюту и статус).
func (p *NewPayment) Create() (*Payment, error) {
	return &Payment{
		APIClient:   p.APIClient,
		ID:          "",
		Status:      "",
		Amount:      Amount{},
		Description: nil,
		Recipient:   Recipient{},
		Requestor:   Requestor{},
		Method:      nil,
		CreatedAt:   "",
		Test:        false,
		Paid:        false,
		Refundable:  false,
	}, nil
}
