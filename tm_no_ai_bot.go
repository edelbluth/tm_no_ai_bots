package tm_no_ai_bots

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type TmNoAiBot struct {
	next    http.Handler
	name    string
	config  *Config
	Matcher *regexp.Regexp
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config == nil || len(config.BotPatterns) == 0 {
		return nil, ErrConfigurationFailed
	}
	quotedBotPatterns := []string{}
	for _, pattern := range config.BotPatterns {
		quotedBotPatterns = append(quotedBotPatterns, regexp.QuoteMeta(pattern))
	}
	matcherCode := fmt.Sprintf("(?i)(%s)", strings.Join(quotedBotPatterns, "|"))
	matcher, err := regexp.Compile(matcherCode)
	if err != nil {
		return nil, err
	}
	return &TmNoAiBot{
		next:    next,
		name:    name,
		config:  config,
		Matcher: matcher,
	}, nil
}
