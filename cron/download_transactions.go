package cron

import (
	"log"

	"github.com/vigneshperiasami/analytics/repository"
	"github.com/vigneshperiasami/analytics/upstream"
)

type TransactionsDownloader struct {
	client upstream.UpstreamClient
	repo   *repository.CompanyClient
	logger *log.Logger
}

func NewTransactionsDownloader(
	client upstream.UpstreamClient,
	repo *repository.CompanyClient,
	logger *log.Logger) Action {
	return &TransactionsDownloader{
		client: client,
		repo:   repo,
		logger: logger,
	}
}

func (t *TransactionsDownloader) Execute() {
	transactions, err := t.client.GetTransactions()

	if err != nil {
		t.logger.Fatalf("Error downloading transactions: %v", err)
	}
	t.logger.Printf("Downloaded transactions: %v", len(transactions))
	err = t.repo.InsertTransactions(transactions)

	if err != nil {
		t.logger.Fatalf("Error inserting transactions: %s", err)
	}
}
