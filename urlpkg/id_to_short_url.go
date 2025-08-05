package urlpkg

import "fmt"

func idToShortUrl(baseUrl string, urlid int) string {
	fmt.Printf("DEBUG - IDTOSHORTURL id: %d\n", urlid)
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if urlid == 0 {
		return string(chars[0])
	}

	var result []byte
	for urlid > 0 {
		remainder := urlid % 62
		result = append(result, chars[remainder])
		urlid = urlid / 62
	}

	// reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	finalResult := baseUrl + string(result)

	fmt.Printf("DEBUG - IDTOSHORTURL finalResult: %s\n", finalResult)
	return finalResult
}
