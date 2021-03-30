package main

import "fmt"

const (
	FACEBOOK = "facebook"
	GOOGLE   = "google"
	TWITTER  = "twitter"
)

func main() {
	socialSharing := NewSocialSharing(FACEBOOK)
	result := socialSharing.Share("https://google.com.vn")
	fmt.Println(result)
}

type SocialSharing interface {
	Share(url string) string
}

func NewSocialSharing(shareType string) SocialSharing {
	switch shareType {
	case FACEBOOK:
		return &facebookSharing{}
	case GOOGLE:
		return &googleSharing{}
	case TWITTER:
		return &twitterSharing{}
	default:
		return nil
	}
}

/*--------------------- facebook implement ---------------------*/
type facebookSharing struct{}
func (*facebookSharing) Share(url string) string {
	return fmt.Sprintf("Share on facebook with url: %s", url)
}

/*---------------------- google implement ----------------------*/
type googleSharing struct{}
func (*googleSharing) Share(url string) string {
	return fmt.Sprintf("Share on google with url: %s", url)
}

/*--------------------- twitter implement ----------------------*/
type twitterSharing struct{}
func (*twitterSharing) Share(url string) string {
	return fmt.Sprintf("Share on twitter with url: %s", url)
}
