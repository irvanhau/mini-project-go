package helper

func FormatResponse(message string, data any) map[string]any {
	var result = map[string]any{}

	result["message"] = message
	if data != nil {
		result["data"] = data
	}

	return result
}

func FormatPagination(message string, data any, current_page, total_page, total_data int) map[string]any {
	var meta = map[string]any{}
	var result = map[string]any{}

	result["message"] = message
	if data != nil {
		result["data"] = data
	}

	meta["current_page"] = current_page
	meta["total_page"] = total_page
	meta["total_data"] = total_data

	result["meta"] = meta

	return result
}
