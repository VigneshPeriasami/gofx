package cron

type CompaniesDownloader struct {
	Downloader
}

func NewCompaniesDownloader(downloader Downloader) Action {
	return &CompaniesDownloader{
		downloader,
	}
}

func (d *CompaniesDownloader) Execute() {
	companies, err := d.client.GetCompanies()
	if err != nil {
		d.logger.Println("Error downloading companies: ", err)
		return
	}
	d.logger.Printf("Total companies downloaded: %d\n", len(companies))
	err = d.repo.InsertCompanies(companies)
	if err != nil {
		d.logger.Fatalf("Error inserting companies: %s", err)
	}
}
