package galer

import (
	"context"
	"errors"
	"time"

	// "github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Config declare its configurations
type Config struct {
	Timeout int
	Proxy   string
	// Headers network.Headers
	Context context.Context
	Cancel  context.CancelFunc
}

// New defines context for the configurations
func New(cfg *Config) *Config {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer(cfg.Proxy))

	execAllocator, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := chromedp.NewContext(execAllocator)

	cfg.Context, cfg.Cancel = context.WithTimeout(ctx, time.Duration(cfg.Timeout)*time.Second)

	return cfg
}

// Crawl to navigate to the URL & dump URLs on it
func (cfg *Config) Crawl(URL string) ([]string, error) {
	var res []string

	if !IsURI(URL) {
		return nil, errors.New("cannot parse URL")
	}

	err := chromedp.Run(
		cfg.Context,
		// network.Enable(),
		// network.SetExtraHTTPHeaders(network.Headers(cfg.Headers)),
		chromedp.Navigate(URL),
		chromedp.Evaluate(script, &res),
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
