package url

func (u *Url) idToShortUrl(baseUrl string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := u.ID

	if id == 0 {
		return string(chars[0])
	}

	var result []byte
	for id > 0 {
		remainder := id % 62
		result = append(result, chars[remainder])
		id = id / 62
	}

	// reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	finalResult := baseUrl + string(result)

	return finalResult
}
