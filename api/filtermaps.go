package api

var FilterableParams = []string{
	"suburb", "street", "outage_type", "search",
	"before_start_date", "before_end_date", "after_end_date",
	"after_start_date", "location", "outage_id",
}

var FilterableCountParams = []string{
	"total_hours", "total_outages",
}

var SQLSigns = map[string]string{
	"after":  ">=",
	"before": "<=",
	"like":   "LIKE",
}
