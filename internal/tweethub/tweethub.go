// Package tweethub provides a simple API for interacting with Twitter using headless browsing.
//
// This package leverages the chromedp library for automated interactions with the Twitter web interface.
package tweethub

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

var twitterURL string = "https://twitter.com"

// TweetHub represents the main interface for interacting with Twitter.
type TweetHub struct {
	username string
	password string
}

// chromeContext returns a new Chrome context and associated cancel function.
// It is used for setting up the headless browser environment.
func chromeContext() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)

	// create chrome instance
	ctx, cancelCtx := chromedp.NewContext(allocCtx)

	// create a timeout
	ctx, cancelTimeout := context.WithTimeout(ctx, 120*time.Second)

	cancel := func() {
		cancelAlloc()
		cancelCtx()
		cancelTimeout()
	}

	return ctx, cancel
}

// New creates a new instance of TweetHub.
func New() *TweetHub {
	return &TweetHub{}
}

// SetUsername sets the Twitter username for the TweetHub instance.
func (t *TweetHub) SetUsername(username string) {
	if username != "" {
		t.username = username
	}
}

// SetPassword sets the Twitter password for the TweetHub instance.
func (t *TweetHub) SetPassword(password string) {
	if password != "" {
		t.password = password
	}
}

// Login performs the login to Twitter with the provided credentials.
// It returns the Chrome context and associated cancel function for further interactions.
func (t TweetHub) Login() (context.Context, context.CancelFunc) {
	twitterLoginURL, _ := url.JoinPath(twitterURL, "login")

	inputUsernameSelector := `//div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/label/div/div[2]/div/input[@autocomplete="username"]`
	inputPasswordSelector := `//div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div/div[3]/div/label/div/div[2]/div[1]/input[@name="password"]`

	cellInnerSelector := `//div/div/div[2]/main/div/div/div/div/div/div[5]/div/section/div/div/div[@data-testid="cellInnerDiv"]`

	ctx, cancel := chromeContext()

	err := chromedp.Run(ctx,
		chromedp.Navigate(twitterLoginURL),

		chromedp.WaitVisible(inputUsernameSelector, chromedp.BySearch),
		chromedp.SendKeys(inputUsernameSelector, t.username+kb.Enter, chromedp.BySearch),

		chromedp.WaitVisible(inputPasswordSelector, chromedp.BySearch),
		chromedp.SendKeys(inputPasswordSelector, t.password+kb.Enter, chromedp.BySearch),

		chromedp.WaitVisible(cellInnerSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to login for user %s: %v", t.username, err)
	} else {
		fmt.Printf("Successful login for user: %s\n", t.username)
	}

	return ctx, cancel
}

// Like performs the "like" action on a given tweet URL.
func (t TweetHub) Like(tweetURL string) context.CancelFunc {
	likeButtonSelector := `//div[3]/div[@data-testid="like"]`
	unlikeButtonSelector := `//div[3]/div[@data-testid="unlike"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(tweetURL),

		chromedp.WaitVisible(likeButtonSelector),
		chromedp.Click(likeButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(unlikeButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to like the content at %s: %v", tweetURL, err)
		return cancel
	} else {
		fmt.Printf("Successfully liked content at: %s\n", tweetURL)
	}

	return cancel
}

// UnLike performs the "unlike" action on a given tweet URL.
func (t TweetHub) UnLike(tweetURL string) context.CancelFunc {
	likeButtonSelector := `//div[3]/div[@data-testid="like"]`
	unlikeButtonSelector := `//div[3]/div[@data-testid="unlike"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(tweetURL),

		chromedp.WaitVisible(unlikeButtonSelector),
		chromedp.Click(unlikeButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(likeButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to unlike the content at %s: %v", tweetURL, err)
		return cancel
	} else {
		fmt.Printf("Successfully unliked content at: %s\n", tweetURL)
	}

	return cancel
}

// Tweet creates a new tweet with the provided message.
func (t TweetHub) Tweet(message string) context.CancelFunc {
	tweetTextareaSelector := `//div/div/div[2]/main/div/div/div/div/div/div[3]/div/div[2]/div[1]/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/div/label/div[1]/div/div/div/div/div/div[2]/div[@data-testid="tweetTextarea_0"]`
	alertSelector := `//div[2]/div/div/div/div[@role="alert"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.SendKeys(tweetTextareaSelector, message),
		chromedp.KeyEvent(kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Enter),

		chromedp.WaitVisible(alertSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to create tweet: %v\n", err)
		return cancel
	} else {
		fmt.Println("Tweet created successfully")
	}

	return cancel
}

// UnTweet deletes an existing tweet identified by its URL.
func (t TweetHub) UnTweet(tweetURL string) context.CancelFunc {
	moreSelector := `//div/div/div[2]/main/div/div/div/div/div/section/div/div/div[1]/div/div/article/div/div/div[2]/div[2]/div/div/div[2]/div/div/div/div/div[@aria-label="More"]`
	alertSelector := `//div[2]/div/div/div/div[@role="alert"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(tweetURL),

		chromedp.WaitVisible(moreSelector, chromedp.BySearch),
		chromedp.Click(moreSelector, chromedp.BySearch, chromedp.NodeVisible),
		chromedp.KeyEvent(kb.Enter),
		chromedp.KeyEvent(kb.Enter),

		chromedp.WaitVisible(alertSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to delete tweet at URL %s: %v\n", tweetURL, err)
		return cancel
	} else {
		fmt.Println("Tweet deleted successfully")
	}

	return cancel
}

// Repost performs the "repost" action on a given post URL.
func (t TweetHub) Repost(postURL string) context.CancelFunc {

	retweetButtonSelector := `//div[2]/div[@data-testid="retweet"]`
	unretweetButtonSelector := `//div[2]/div[@data-testid="unretweet"]`
	repostButtonSelector := `//div[2]/div/div/div/div[2]/div/div[3]/div/div/div/div[@data-testid="retweetConfirm"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(postURL),

		chromedp.WaitVisible(retweetButtonSelector, chromedp.BySearch),
		chromedp.Click(retweetButtonSelector, chromedp.BySearch, chromedp.NodeVisible),
		chromedp.WaitVisible(repostButtonSelector, chromedp.BySearch),
		chromedp.Click(repostButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(unretweetButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to repost: %v\n", err)
	} else {
		fmt.Println("Repost successful")
	}

	return cancel
}

// UnRepost performs the "unrepost" action on a given post URL.
func (t TweetHub) UnRepost(postURL string) context.CancelFunc {

	retweetButtonSelector := `//div[2]/div[@data-testid="retweet"]`
	unretweetButtonSelector := `//div[2]/div[@data-testid="unretweet"]`
	unrepostButtonSelector := `///div[2]/div/div/div/div[2]/div/div[3]/div/div/div/div[@data-testid="unretweetConfirm"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(postURL),

		chromedp.WaitVisible(unretweetButtonSelector, chromedp.BySearch),
		chromedp.Click(unretweetButtonSelector, chromedp.BySearch, chromedp.NodeVisible),
		chromedp.WaitVisible(unrepostButtonSelector, chromedp.BySearch),
		chromedp.Click(unrepostButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(retweetButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to unrepost: %v\n", err)
		return cancel
	} else {
		fmt.Println("Unrepost successful")
	}

	return cancel
}

// Quote performs the "quote" action on a given post URL with an optional custom message.
func (t TweetHub) Quote(postURL string, message ...string) context.CancelFunc {
	retweetButtonSelector := `//div[2]/div[@data-testid="retweet"]`

	tweetTextareaSelector := `//div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div/div[3]/div[2]/div[1]/div/div/div/div[1]/div[2]/div/div/div/div/div/div/div/div/div/div/div[1]/label/div[1]/div/div/div/div/div/div[2]/div[@data-testid="tweetTextarea_0"]`
	tweetPostButtonSelector := `//div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div/div[3]/div[2]/div[1]/div/div/div/div[2]/div[2]/div/div/div/div[@data-testid="tweetButton"]`
	alertSelector := `//div[2]/div/div/div/div[@role="alert"]`

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(postURL),

		chromedp.WaitVisible(retweetButtonSelector, chromedp.BySearch),
		chromedp.Click(retweetButtonSelector, chromedp.BySearch, chromedp.NodeVisible),
		chromedp.KeyEvent(kb.ArrowDown),
		chromedp.KeyEvent(kb.Enter),

		chromedp.WaitVisible(tweetTextareaSelector, chromedp.BySearch),
		chromedp.SendKeys(tweetTextareaSelector, message[0], chromedp.BySearch),
		chromedp.Click(tweetPostButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(alertSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to quote: %v\n", err)
	} else {
		fmt.Println("Quote successful")
	}

	return cancel
}

// Follow performs the "follow" action on a specified Twitter username.
func (t TweetHub) Follow(username string) context.CancelFunc {
	profileURL, _ := url.JoinPath(twitterURL, username)

	followButtonSelector := fmt.Sprintf(`//div/div/div[2]/main/div/div/div/div/div/div[3]/div/div/div/div/div[1]/div[2]/div[2]/div[1]/div[@aria-label="Follow @%s"]`, username)
	followingButtonSelector := fmt.Sprintf(`//div/div/div[2]/main/div/div/div/div/div/div[3]/div/div/div/div/div[1]/div[2]/div[3]/div[1]/div[@aria-label="Following @%s"]`, username)

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(profileURL),

		chromedp.WaitVisible(followButtonSelector, chromedp.BySearch),
		chromedp.Click(followButtonSelector, chromedp.BySearch, chromedp.NodeVisible),

		chromedp.WaitVisible(followingButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to follow @%s: %v\n", username, err)
	} else {
		fmt.Printf("Followed @%s successfully\n", username)
	}

	return cancel
}

// UnFollow performs the "unfollow" action on a specified Twitter username.
func (t TweetHub) UnFollow(username string) context.CancelFunc {
	profileURL, _ := url.JoinPath(twitterURL, username)

	followButtonSelector := fmt.Sprintf(`//div/div/div[2]/main/div/div/div/div/div/div[3]/div/div/div/div/div[1]/div[2]/div[2]/div[1]/div[@aria-label="Follow @%s"]`, username)
	followingButtonSelector := fmt.Sprintf(`//div/div/div[2]/main/div/div/div/div/div/div[3]/div/div/div/div/div[1]/div[2]/div[3]/div[1]/div[@aria-label="Following @%s"]`, username)

	ctx, cancel := t.Login()

	err := chromedp.Run(ctx,
		chromedp.Navigate(profileURL),

		chromedp.WaitVisible(followingButtonSelector, chromedp.BySearch),
		chromedp.Click(followingButtonSelector, chromedp.BySearch, chromedp.NodeVisible),
		chromedp.KeyEvent(kb.Enter),

		chromedp.WaitVisible(followButtonSelector, chromedp.BySearch),
	)

	if err != nil {
		fmt.Printf("Failed to unfollow @%s: %v\n", username, err)
	} else {
		fmt.Printf("Unfollowed @%s successfully\n", username)
	}

	return cancel
}
