package vector

import (
	"slices"
	"time"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/risk"
)

const (
	maxAmount            = 10000
	maxInstallments      = 12.0
	amountVsAvgRatio     = 10
	maxMinutes           = 1440
	maxKm                = 1000
	maxTxCount24h        = 20
	maxMerchantAvgAmount = 10000
)

type Builder struct {
	mccRisk risk.MCCRisk
}

func NewBuilder(mccRisk risk.MCCRisk) *Builder {
	return &Builder{
		mccRisk: mccRisk,
	}
}

func (b *Builder) Build(request domain.FraudScoreRequest) Vector {
	var v Vector

	// D1
	v[0] = Clamp(
		request.Transaction.Amount / maxAmount,
	)

	// D2
	v[1] = Clamp(
		float64(request.Transaction.Installments) / maxInstallments,
	)

	// D3
	if request.Customer.AvgAmount == 0 {
		v[2] = 1
	} else {
		v[2] = Clamp(
			request.Transaction.Amount /
				request.Customer.AvgAmount /
				amountVsAvgRatio,
		)
	}

	// D4 and D5
	t, _ := time.Parse(time.RFC3339, request.Transaction.RequestedAt)

	hour := t.Hour()
	v[3] = float32(hour) / 23

	dayOfWeek := (int(t.Weekday()) + 6) % 7
	v[4] = float32(dayOfWeek) / 6

	// D6 and D7
	if request.LastTransaction == nil {
		v[5] = -1
		v[6] = -1
	} else {
		lastTransactionTime, _ := time.Parse(time.RFC3339, request.LastTransaction.Timestamp)
		minutes := t.Sub(lastTransactionTime).Minutes()
		v[5] = Clamp(
			minutes / maxMinutes,
		)
		v[6] = Clamp(
			float64(request.LastTransaction.KmFromCurrent) / maxKm,
		)
	}

	// D8
	v[7] = Clamp(
		float64(request.Terminal.KmFromHome) / maxKm,
	)

	// D9
	v[8] = Clamp(
		float64(request.Customer.TxCount24h) / maxTxCount24h,
	)

	// D10
	isOnline := 0
	if request.Terminal.IsOnline {
		isOnline = 1
	}
	v[9] = float32(isOnline)

	// D11
	cardPresent := 0
	if request.Terminal.CardPresent {
		cardPresent = 1
	}
	v[10] = float32(cardPresent)

	// D12
	unknownMerchant := 1
	if slices.Contains(request.Customer.KnownMerchants, request.Merchant.ID) {
		unknownMerchant = 0
	}
	v[11] = float32(unknownMerchant)

	// D13
	riskScore, found := b.mccRisk[request.Merchant.MCC]
	if !found {
		riskScore = 0.5
	}
	v[12] = riskScore

	// D14
	v[13] = Clamp(
		request.Merchant.AvgAmount / maxMerchantAvgAmount,
	)

	return v
}
