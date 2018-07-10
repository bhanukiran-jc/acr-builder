// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package builder

import (
	"context"
	"fmt"
	"os"
	"time"
)

const (
	maxPushRetries = 3
)

func (b *Builder) pushWithRetries(ctx context.Context, images []string) error {
	if len(images) <= 0 {
		return nil
	}

	for _, img := range images {
		args := []string{"docker", "push", img}

		retry := 0
		for retry < maxPushRetries {
			fmt.Printf("Pushing image: %s, attempt %d\n", img, retry+1)

			if err := b.cmder.Run(ctx, args, nil, os.Stdout, os.Stderr, ""); err != nil {
				time.Sleep(500 * time.Millisecond)
				retry++
			} else {
				break
			}
		}

		if retry == maxPushRetries {
			return fmt.Errorf("Failed to push images successfully")
		}
	}

	return nil
}