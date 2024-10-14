package service

import (
	"context"
	"harsh/internal/data"
	model "harsh/internal/models"
	"log"
	"os"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentService struct {
	paymentData *data.PaymentData
}

func NewPaymentService(paymentData *data.PaymentData) *PaymentService {
	return &PaymentService{
		paymentData: paymentData,
	}
}

func (ps *PaymentService) ProcessPayment(ctx context.Context, payment *model.Payment) (string, error) {
	payment.Status = "Pending"
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	STRIPE_API := os.Getenv("STRIPE_API")
	if STRIPE_API == "" {
		log.Fatal("failed to get stripe api env")
	}
	// 	stripe api client
	stripe.Key = STRIPE_API

	// Creating a stripe payment intent
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(int64(payment.Amount * 100)),
		Currency:      stripe.String(payment.Currency),
		PaymentMethod: stripe.String("pm_card_visa"), //testing purpose
		Confirm:       stripe.Bool(true),
	}

	// triggering the payment
	result, err := paymentintent.New(params)
	if err != nil {
		payment.Status = "Failed"
		payment.UpdatedAt = time.Now() // Update updated_at when the status changes
		if saveErr := ps.paymentData.SavePayment(ctx, payment); saveErr != nil {
			log.Fatal("failed to save payment data")
		}
		return "", err
	}

	// Payment successful, update status to Success
	payment.Status = "Success"
	payment.UpdatedAt = time.Now()

	// Save the payment to the database
	Newerr := ps.paymentData.SavePayment(ctx, payment)
	if Newerr != nil {
		return "", Newerr
	}

	return result.ID, nil
}

func (ps *PaymentService) GetPaymentById(ctx context.Context, id primitive.ObjectID) (*model.Payment, error) {

	return ps.paymentData.GetPaymentById(ctx, id)
}
