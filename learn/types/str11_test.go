package cache_sdk

//
//import (
//	"awesomeProject/zues-learn/logger"
//	bigcache "github.com/allegro/bigcache"
//	"testing"
//	"time"
//)
//
//func Test_creat11(t *testing.T) {
//
//
//}
//
//
//
//func ZeusNewBigCache(shard int, lifeWindows time.Duration, hardMaxCacheSize int, maxEntriesInWindow int) *bigcache.BigCache {
//	var config = bigcache.Config{
//		// number of shards (must be a power of 2)
//		Shards: shard,
//
//		// time after which entry can be evicted
//		LifeWindow: lifeWindows,
//
//		// Interval between removing expired entries (clean up).
//		// If set to <= 0 then no action is performed.
//		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
//		CleanWindow: 10 * time.Minute,
//
//		// rps * lifeWindow, used only in initial memory allocation
//		MaxEntriesInWindow: maxEntriesInWindow,
//
//		// max entry size in bytes, used only in initial memory allocation
//		MaxEntrySize: 10,
//
//		// prints information about additional memory allocation
//		Verbose: false,
//
//		// cache will not allocate more memory than this limit, value in MB
//		// if value is reached then the oldest entries can be overridden for the new ones
//		// 0 value means no size limit
//		HardMaxCacheSize: hardMaxCacheSize,
//
//		// callback fired when the oldest entry is removed because of its expiration time or no space left
//		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
//		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
//		OnRemove: nil,
//
//		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
//		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
//		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
//		// Ignored if OnRemove is specified.
//		OnRemoveWithReason: nil,
//	}
//	cache, initErr := bigcache.NewBigCache(config)
//	if initErr != nil {
//		logger.Logger.WithField("init bigcache error", initErr).Warn("_undef")
//	}
//	return cache
//}
