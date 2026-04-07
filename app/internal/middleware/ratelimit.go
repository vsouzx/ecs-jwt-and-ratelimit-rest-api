package middleware

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	redis "github.com/redis/go-redis/v9"
)

func RateLimiter(redisClient *redis.Client) fiber.Handler {
	maxRequests, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_COUNT"))
	windowMinutes, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TTL"))
	windowDuration := time.Duration(windowMinutes) * time.Minute

	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			return c.Next()
		}

		ctx := context.Background()
		rateLimitKey := fmt.Sprintf("rl:ip:%s", c.IP())

		requestCount, err := redisClient.Incr(ctx, rateLimitKey).Result()
		if err != nil {
			fmt.Printf("rate-limit: erro no INCR: %v\n", err)
			return c.Next()
		}

		if requestCount == 1 {
			if err := redisClient.Expire(ctx, rateLimitKey, windowDuration).Err(); err != nil {
				fmt.Printf("rate-limit: erro no EXPIRE: %s", err.Error())
			}
		}

		windowTimeLeft, err := redisClient.PTTL(ctx, rateLimitKey).Result()
		if err != nil || windowTimeLeft < 0 {
			windowTimeLeft = windowDuration
		}

		remainingRequests := maxRequests - int(requestCount)
		if remainingRequests < 0 {
			remainingRequests = 0
		}

		c.Set("X-RateLimit-Limit", strconv.Itoa(maxRequests))
		c.Set("X-RateLimit-Remaining", strconv.Itoa(remainingRequests))

		if int(requestCount) > maxRequests {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "rate limit exceeded",
			})
		}
		return c.Next()
	}
}
