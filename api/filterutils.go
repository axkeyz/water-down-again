package api

import "strings"

// IsFilterableParam returns true if a (url) query parameter is filterable.
func IsFilterableParam(param string) bool {
	filterables := []string{
		"suburb", "street", "outage_type", "search",
		"before_start_date", "before_end_date", "after_end_date",
		"after_start_date", "start_date", "end_date",
		"location", "outage_id",
	}

	return isStringInArray(param, filterables)
}

// IsFilterableCountParam returns true if a query parameter is a count API
// parameter. It is intended to be an extension of IsFilterableParam.
func IsFilterableCountParam(param string) bool {
	filterables := []string{"total_hours", "total_outages"}

	return isStringInArray(param, filterables)
}

// IsCurrentOutageID checks if an outage id is in a list of current
// outage ids.
func IsCurrentOutageID(outage_id int, current_outage_ids []int) bool {
	return isIntInArray(outage_id, current_outage_ids)
}

// IsDateParam returns true if a parameter corresponds to a date
// column, and if so, it returns the actual column name.
func IsDateParam(param string) (isDate bool, column string) {
	if strings.Contains(param, "end_date") ||
		strings.Contains(param, "start_date") {
		isDate = true
		column = getEquationSignedColumn(param)
	}
	return
}

// getEquationSignedColumn returns the column name and the equation
// sign from the parameter string.
// For example:
//		"after_end_date" returns "end_date >="
//		"before_start_date" returns "start_date <="
func getEquationSignedColumn(param string) string {
	p := strings.Split(param, "_")
	return GetNWordsRemovedFromStart(param, "_", 1) + " " +
		getSQLEquationSigns(p[0])
}

// getSQLEquationSigns returns the corresponding equation sign
// to the given keys.
func getSQLEquationSigns(key string) string {
	signs := map[string]string{
		"after":  ">=",
		"before": "<=",
		"like":   "LIKE",
	}

	if sign, ok := signs[key]; ok {
		return sign
	} else {
		return "="
	}
}
