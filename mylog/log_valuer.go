package mylog

import (
	"log/slog"
	"os"
)

type User struct {
	Username string
	Email    string
	Age      int
	Password string
}

type UserLogValuer struct {
	User User
}

func (u UserLogValuer) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("username", u.User.Username),
		slog.Int("age", u.User.Age),
	)
}

// LogValuer shows how to use the LogValue method to customize how a type is logged
func LogValuer() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("User info", slog.Any(
		"user",
		User{Username: "gopher", Email: "gopher@example.com", Age: 5, Password: "secret"}),
	)

	logger.Info("User info with valuer", slog.Any(
		"user",
		UserLogValuer{User: User{Username: "gopher", Email: "gopher@example.com", Age: 5, Password: "secret"}}),
	)
}
