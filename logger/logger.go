package logger

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			start := time.Now()

			fields := []zapcore.Field{
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
				zap.String("referer", req.Referer()),
			}

			err := next(c)

			latency := time.Since(start)

			fields = append(fields,
				zap.Int("status", res.Status),
				zap.String("latency", latency.String()),
			)

			if err != nil {
				fields = append(fields, zap.Error(err))
				logger.Error("HTTP Request", fields...)
			} else {
				switch {
				case res.Status >= 500:
					logger.Error("HTTP Request", fields...)
				case res.Status >= 400:
					logger.Warn("HTTP Request", fields...)
				default:
					logger.Info("HTTP Request", fields...)
				}
			}

			return err
		}
	})
}
