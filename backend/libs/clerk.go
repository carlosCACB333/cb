package libs

import (
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func ClerkClient() clerk.Client {
	clint, _ := clerk.NewClient(os.Getenv("CLERK_SECRET_KEY"))
	return clint
}
