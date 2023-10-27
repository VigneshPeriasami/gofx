package cron

type TransactionsDownloader struct {
	Downloader
}

func NewTransactionsDownloader(downloader Downloader) Action {
	return &TransactionsDownloader{
		downloader,
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
