.DEFAULT_GOAL := /usr/local/bin/stats

/usr/local/bin/stats: stats.go
	go build -o /usr/local/bin/stats stats.go
