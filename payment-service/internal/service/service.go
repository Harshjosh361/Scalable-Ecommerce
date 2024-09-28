package service

import (
	"context"
	"harsh/internal/data"
	model "harsh/internal/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// load env variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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

	// 	stripe api client
	stripe.Key = os.Getenv("STRIPE_API")

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
		ps.paymentData.SavePayment(ctx, payment)
		return "", err
	}

	// Payment successful, update status to Success
	payment.Status = "Success"
	payment.UpdatedAt = time.Now()

	// Save the payment to the database
	err = ps.paymentData.SavePayment(ctx, payment)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func (ps *PaymentService) GetPaymentById(ctx context.Context, id primitive.ObjectID) (*model.Payment, error) {

	return ps.paymentData.GetPaymentById(ctx, id)
}
