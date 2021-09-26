/**
    package: sco_tracers
    filename: middleware
    author: diogo@gmail.com
    time: 2021/9/14 11:32
**/
package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"strconv"
	"sync"
	"time"
)

// Config ...
type LimiterConfig struct {
	// Environment dev or prod
	// Default dev
	Environment string
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(request *ghttp.Request) bool
	// Timeout in seconds on how long to keep records of requests in memory
	// Default: 60
	Timeout int
	// Max number of recent connections during `Timeout` seconds before sending a 429 response
	// Default: 10
	Max int
	// Message
	// default: "Too many requests, please try again later."
	Message string
	// StatusCode
	// Default: 429 Too Many Requests
	StatusCode int
	// Key allows to use a custom handler to create custom keys
	// Default: func(c *ghttp.Request) string {
	//   return c.GetClientIp()
	// }
	Key func(request *ghttp.Request) string
	// Handler is called when a request hits the limit
	// Default: func(c *ghttp.Request) {
	//   c.Status(cfg.StatusCode).SendString(cfg.Message)
	// }
	Handler func(request *ghttp.Request)
}

// 初始化
func NewLimiter(config ...LimiterConfig) func(r *ghttp.Request) {
	// mutex for parallel read and write access
	mux := &sync.Mutex{}
	// Init config
	var cfg LimiterConfig
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Environment == ""{
		cfg.Environment = "dev"
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 60
	}
	if cfg.Max == 0 {
		cfg.Max = 10
	}

	if cfg.Message == "" {
		cfg.Message = "Too many requests, please try again later."
	}
	if cfg.StatusCode == 0 {
		cfg.StatusCode = 429
	}
	if cfg.Key == nil {
		cfg.Key = func(c *ghttp.Request) string {
			return c.GetClientIp()
		}
	}

	if cfg.Handler == nil {
		cfg.Handler = func(c *ghttp.Request) {
			// c.Status(cfg.StatusCode).SendString(cfg.Message)
			// c.Response.Status(cfg.StatusCode)
			c.Response.WriteStatus(cfg.StatusCode, cfg.Message)
			glog.Println("限制访问==")
		}
	}
	// Limiter settings
	var hits = make(map[string]int)
	var reset = make(map[string]int)
	var timestamp = int(time.Now().Unix())
	// Update timestamp every second
	go func() {
		for {
			timestamp = int(time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}()

	// Reset hits every cfg.Timeout
	go func() {
		for {
			// For every key in reset
			for key := range reset {
				// If resetTime exist and current time is equal or bigger
				if reset[key] != 0 && timestamp >= reset[key] {
					// Reset hits and resetTime
					mux.Lock()
					hits[key] = 0
					reset[key] = 0
					mux.Unlock()
				}
			}
			// Wait cfg.Timeout
			time.Sleep(time.Duration(cfg.Timeout) * time.Second)
		}
	}()

	return func(c *ghttp.Request) {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Middleware.Next()
			return
		}
		// Get key (default is the remote IP)
		key := cfg.Key(c)
		mux.Lock()
		// Increment key hits
		hits[key]++
		// Set unix timestamp if not exist
		if reset[key] == 0 {
			reset[key] = timestamp + cfg.Timeout
		}
		// Get current hits
		hitCount := hits[key]
		// Calculate when it resets in seconds
		resetTime := reset[key] - timestamp
		mux.Unlock()
		// Set how many hits we have left
		remaining := cfg.Max - hitCount
		// Check if hits exceed the cfg.Max
		if remaining < 0 {
			// Call Handler func
			cfg.Handler(c)
			// Return response with Retry-After header
			// https://tools.ietf.org/html/rfc6584
			c.Response.Header().Set("Retry-After", strconv.Itoa(resetTime))
			return
		}
		// We can continue, update RateLimit headers
		// FIXME: 开发的时候.可开启. 正常上线关闭
		if cfg.Environment == "dev" {
			c.Response.Header().Set("X-RateLimit-Limit", strconv.Itoa(cfg.Max))
			c.Response.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
			c.Response.Header().Set("X-RateLimit-Reset", strconv.Itoa(resetTime))
		}
		// Bye!
		c.Middleware.Next()
	}

}
