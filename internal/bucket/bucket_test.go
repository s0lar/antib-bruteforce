package bucket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBucket_Allow(t *testing.T) {
	interval := 2 * time.Millisecond
	ttl := interval * 2
	bucket := NewBucket(2, interval, ttl)

	require.True(t, bucket.Allow("test"))
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))

	// Check reset by time
	time.Sleep(interval * 2)
	require.True(t, bucket.Allow("test"))
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))
}

func TestBucket_Reset(t *testing.T) {
	interval := 2 * time.Millisecond
	ttl := interval * 2
	bucket := NewBucket(2, interval, ttl)

	require.True(t, bucket.Allow("test"))
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))

	bucket.Reset("test")

	require.True(t, bucket.Allow("test"))
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))

	//	Reset not exists key
	bucket.Reset("test2")
	require.True(t, bucket.Allow("test2"))
	require.True(t, bucket.Allow("test2"))
	require.False(t, bucket.Allow("test2"))
}

func TestBucket_BucketGC(t *testing.T) {
	interval := 5 * time.Minute
	ttl := 1 * time.Millisecond
	sleep := 2 * time.Millisecond
	bucket := NewBucket(1, interval, ttl)

	//	Clear bucket for key
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))

	//	Nothing to clear. TTL is not expired
	bucket.BucketGC()
	require.False(t, bucket.Allow("test"))

	//	Sleeping for TTL expiring
	time.Sleep(sleep)
	bucket.BucketGC()

	//	Clear bucket for key
	require.True(t, bucket.Allow("test"))
	require.False(t, bucket.Allow("test"))
}
