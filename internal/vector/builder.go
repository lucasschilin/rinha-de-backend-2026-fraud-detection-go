package vector

import (
	"fmt"
	"log"
	"time"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
)

const (
	maxAmount            = 10000
	maxInstallments      = 12.0
	amountVsAvgRation    = 10
	maxMinutes           = 1440
	maxKm                = 1000
	maxTxCount24h        = 20
	maxMerchantAvgAmount = 10000
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(request domain.FraudScoreRequest) Vector {
	fmt.Println("")
	var v Vector

	// D1
	log.Printf("amount=%f", request.Transaction.Amount)
	v[0] = Clamp(
		request.Transaction.Amount / maxAmount,
	)

	// D2
	log.Printf("installments=%d", request.Transaction.Installments)
	v[1] = Clamp(
		float64(request.Transaction.Installments) / maxInstallments,
	)

	// D3
	log.Printf("avg_amount=%f", request.Customer.AvgAmount)
	v[2] = Clamp(
		request.Transaction.Amount / request.Customer.AvgAmount / amountVsAvgRation,
	)

	// D4 and D5
	t, _ := time.Parse(time.RFC3339, request.Transaction.RequestedAt)

	hour := t.Hour()
	log.Printf("hour=%d", hour)
	v[3] = float32(hour) / 23

	dayOfWeek := (int(t.Weekday()) + 6) % 7
	log.Printf("day_of_week=%d", dayOfWeek)
	v[4] = float32(dayOfWeek) / 6

	// D6 and D7
	if request.LastTransaction == nil {
		log.Print("minutes_since_last_transaction=nil")
		v[5] = -1
		log.Print("km_from_current=nil")
		v[6] = -1
	} else {
		lastTransactionTime, _ := time.Parse(time.RFC3339, request.LastTransaction.Timestamp)
		minutes := t.Sub(lastTransactionTime).Minutes()
		log.Printf("minutes_since_last_transaction=%f", minutes)
		v[5] = Clamp(
			minutes / maxMinutes,
		)
		log.Printf("km_from_current=%f", request.LastTransaction.KmFromCurrent)
		v[6] = Clamp(
			float64(request.LastTransaction.KmFromCurrent) / maxKm,
		)
	}

	// D8
	log.Printf("km_from_home=%f", request.Terminal.KmFromHome)
	v[7] = Clamp(
		float64(request.Terminal.KmFromHome) / maxKm,
	)

	// D9
	log.Printf("tx_count_24h=%v", request.Customer.TxCount24h)
	v[8] = Clamp(
		float64(request.Customer.TxCount24h) / maxTxCount24h,
	)

	// D10
	log.Printf("is_online=%v", request.Terminal.IsOnline)
	isOnline := 0
	if request.Terminal.IsOnline {
		isOnline = 1
	}
	v[9] = float32(isOnline)

	// D11
	log.Printf("card_present=%v", request.Terminal.CardPresent)
	cardPresent := 0
	if request.Terminal.CardPresent {
		cardPresent = 1
	}
	v[10] = float32(cardPresent)

	// D12
	v[11] = 999 // TODO: Verify if the merchant is known

	// D13
	v[12] = 999 // TODO: Verify if the merchant MCC is known

	// D14
	log.Printf("merchant_avg_amount=%f", request.Merchant.AvgAmount)
	v[13] = Clamp(
		request.Merchant.AvgAmount / maxMerchantAvgAmount,
	)

	return v
}
