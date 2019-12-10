package con

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := NewContext(context.Background(), &User{
		Name: "Demon",
	})

	u, _ := FromContext(ctx)

	fmt.Println(u.Name)
}
