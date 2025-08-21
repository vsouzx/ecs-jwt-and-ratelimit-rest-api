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
	skip := map[string]struct{}{
		"/login": {},
	}

	limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_COUNT"))
	ttlInt, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TTL"))

	ttl := time.Duration(ttlInt) * time.Minute

	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		if c.Method() == fiber.MethodOptions {
			return c.Next()
		}
		if _, ok := skip[c.Path()]; ok {
			return c.Next()
		}

		clientIP := c.IP()
		key := fmt.Sprintf("%s:%s", "rl:ip", clientIP)

		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			// Fail-open: se o Redis falhar, não bloqueia a requisição
			fmt.Printf("rate-limit: erro no INCR: %v\n", err)
			return c.Next()
		}

		if count == 1 {
			if err := redisClient.Expire(ctx, key, ttl).Err(); err != nil {
				fmt.Printf("rate-limit: erro no EXPIRE: %s", err.Error())
			}
		}

		pttl, err := redisClient.PTTL(ctx, key).Result()
		if err != nil || pttl < 0 {
			pttl = ttl
		}

		resetUnix := time.Now().Add(pttl).Unix()
		remaining := limit - int(count)
		if remaining < 0 {
			remaining = 0
		}

		c.Set("X-RateLimit-Limit", strconv.Itoa(limit))
		c.Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Set("X-RateLimit-Reset", strconv.FormatInt(resetUnix, 10))

		if int(count) > limit {
			retryAfter := int((pttl + (time.Second - 1)) / time.Second)
			c.Set("Retry-After", strconv.Itoa(retryAfter))

			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "rate limit exceeded",
				"limit":       limit,
				"remaining":   remaining,
				"reset_unix":  resetUnix,
				"retry_after": retryAfter,
			})
		}

		return c.Next()
	}
}
