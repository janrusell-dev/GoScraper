package main

import (
	"github.com/janrusell-dev/goscraper/pkg/cron"
	"github.com/janrusell-dev/goscraper/pkg/scraper"
)

func main() {
	scraper.Scraper()
	cron.Cron()
}

