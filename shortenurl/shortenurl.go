package shortenurl

import (
    "fmt"
    "math/rand"
)

func ShortenURL(url string) string {
    // create a random string of size 6 from the alphabet
    s := ""
    for i := 0; i < 6; i++ {
        s += string(rand.Intn(26) + 97)
    }

    shortenedURL := fmt.Sprintf("http://localhost:8080/%s", s)
    return shortenedURL
}
