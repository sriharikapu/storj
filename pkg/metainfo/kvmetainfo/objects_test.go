// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kvmetainfo

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/pkg/storage/objects"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storage"
)

func TestGetObject(t *testing.T) {
	runTest(t, func(ctx context.Context, db *DB) {
		bucket, err := db.CreateBucket(ctx, TestBucket, nil)
		if !assert.NoError(t, err) {
			return
		}

		store, err := db.buckets.GetObjectStore(ctx, bucket.Name)
		if !assert.NoError(t, err) {
			return
		}

		var exp time.Time
		_, err = store.Put(ctx, "test-file", bytes.NewReader(nil), objects.SerializableMeta{}, exp)
		if !assert.NoError(t, err) {
			return
		}

		_, err = db.GetObject(ctx, "", "")
		assert.True(t, storj.ErrNoBucket.Has(err))

		_, err = db.GetObject(ctx, bucket.Name, "")
		assert.True(t, storage.ErrEmptyKey.Has(err))

		_, err = db.GetObject(ctx, "non-existing-bucket", "test-file")
		// TODO: Should return storj.ErrBucketNotFound
		assert.True(t, storage.ErrKeyNotFound.Has(err))

		_, err = db.GetObject(ctx, bucket.Name, "non-existing-file")
		assert.True(t, storage.ErrKeyNotFound.Has(err))

		object, err := db.GetObject(ctx, bucket.Name, "test-file")
		if assert.NoError(t, err) {
			assert.Equal(t, "test-file", object.Path)
		}
	})
}

func TestGetObjectStream(t *testing.T) {
	runTest(t, func(ctx context.Context, db *DB) {
		bucket, err := db.CreateBucket(ctx, TestBucket, nil)
		if !assert.NoError(t, err) {
			return
		}

		store, err := db.buckets.GetObjectStore(ctx, bucket.Name)
		if !assert.NoError(t, err) {
			return
		}

		var exp time.Time
		_, err = store.Put(ctx, "empty-file", bytes.NewReader(nil), objects.SerializableMeta{}, exp)
		if !assert.NoError(t, err) {
			return
		}

		_, err = store.Put(ctx, "test-file", bytes.NewReader([]byte("test")), objects.SerializableMeta{}, exp)
		if !assert.NoError(t, err) {
			return
		}

		_, err = db.GetObjectStream(ctx, "", "")
		assert.True(t, storj.ErrNoBucket.Has(err))

		_, err = db.GetObjectStream(ctx, bucket.Name, "")
		assert.True(t, storage.ErrEmptyKey.Has(err))

		_, err = db.GetObjectStream(ctx, "non-existing-bucket", "test-file")
		// TODO: Should return storj.ErrBucketNotFound
		assert.True(t, storage.ErrKeyNotFound.Has(err))

		_, err = db.GetObject(ctx, bucket.Name, "non-existing-file")
		assert.True(t, storage.ErrKeyNotFound.Has(err))

		stream, err := db.GetObjectStream(ctx, bucket.Name, "empty-file")
		if assert.NoError(t, err) {
			assertStream(ctx, t, stream, "empty-file", []byte(nil))
		}

		stream, err = db.GetObjectStream(ctx, bucket.Name, "test-file")
		if assert.NoError(t, err) {
			assertStream(ctx, t, stream, "test-file", []byte("test"))
		}
	})
}

func assertStream(ctx context.Context, t *testing.T, stream storj.ReadOnlyStream, path storj.Path, content []byte) bool {
	assert.Equal(t, path, stream.Info().Path)

	segments, more, err := stream.Segments(ctx, 0, 0)
	if !assert.NoError(t, err) {
		return false
	}

	assert.False(t, more)
	if !assert.Equal(t, 1, len(segments)) {
		return false

	}
	assert.EqualValues(t, 0, segments[0].Index)
	assert.EqualValues(t, len(content), segments[0].Size)

	// TODO: Currently Inline is always empty
	// assert.Equal(t, content, segments[0].Inline)

	return true
}

func TestDeleteObject(t *testing.T) {
	runTest(t, func(ctx context.Context, db *DB) {
		bucket, err := db.CreateBucket(ctx, TestBucket, nil)
		if !assert.NoError(t, err) {
			return
		}

		store, err := db.buckets.GetObjectStore(ctx, bucket.Name)
		if !assert.NoError(t, err) {
			return
		}

		var exp time.Time
		_, err = store.Put(ctx, "test-file", bytes.NewReader(nil), objects.SerializableMeta{}, exp)
		if !assert.NoError(t, err) {
			return
		}

		err = db.DeleteObject(ctx, "", "")
		assert.True(t, storj.ErrNoBucket.Has(err))

		err = db.DeleteObject(ctx, bucket.Name, "")
		assert.True(t, storage.ErrEmptyKey.Has(err))

		_ = db.DeleteObject(ctx, "non-existing-bucket", "test-file")
		// TODO: Currently returns minio.BucketNotFound, should return storj.ErrBucketNotFound
		// assert.True(t, storj.ErrBucketNotFound.Has(err))

		err = db.DeleteObject(ctx, bucket.Name, "non-existing-file")
		assert.True(t, storage.ErrKeyNotFound.Has(err))

		err = db.DeleteObject(ctx, bucket.Name, "test-file")
		assert.NoError(t, err)
	})
}

func TestListObjectsEmpty(t *testing.T) {
	runTest(t, func(ctx context.Context, db *DB) {
		bucket, err := db.CreateBucket(ctx, TestBucket, nil)
		if !assert.NoError(t, err) {
			return
		}

		_, err = db.ListObjects(ctx, "", storj.ListOptions{})
		assert.True(t, storj.ErrNoBucket.Has(err))

		_, err = db.ListObjects(ctx, bucket.Name, storj.ListOptions{})
		assert.EqualError(t, err, "kvmetainfo: invalid direction 0")

		for _, direction := range []storj.ListDirection{
			storj.Before,
			storj.Backward,
			storj.Forward,
			storj.After,
		} {
			list, err := db.ListObjects(ctx, bucket.Name, storj.ListOptions{Direction: direction})
			if assert.NoError(t, err) {
				assert.False(t, list.More)
				assert.Equal(t, 0, len(list.Items))
			}
		}
	})
}

func TestListObjects(t *testing.T) {
	runTest(t, func(ctx context.Context, db *DB) {
		var exp time.Time
		bucket, err := db.CreateBucket(ctx, TestBucket, &storj.Bucket{PathCipher: storj.Unencrypted})
		if !assert.NoError(t, err) {
			return
		}

		store, err := db.buckets.GetObjectStore(ctx, bucket.Name)
		if !assert.NoError(t, err) {
			return
		}

		filePaths := []string{
			"a", "aa", "b", "bb", "c",
			"a/xa", "a/xaa", "a/xb", "a/xbb", "a/xc",
			"b/ya", "b/yaa", "b/yb", "b/ybb", "b/yc",
		}
		for _, path := range filePaths {
			_, err = store.Put(ctx, path, bytes.NewReader(nil), objects.SerializableMeta{}, exp)
			if !assert.NoError(t, err) {
				return
			}
		}

		otherBucket, err := db.CreateBucket(ctx, "otherbucket", nil)
		if !assert.NoError(t, err) {
			return
		}

		otherStore, err := db.buckets.GetObjectStore(ctx, otherBucket.Name)
		if !assert.NoError(t, err) {
			return
		}

		_, err = otherStore.Put(ctx, "file-in-other-bucket", bytes.NewReader(nil), objects.SerializableMeta{}, exp)
		if !assert.NoError(t, err) {
			return
		}

		for i, tt := range []struct {
			options storj.ListOptions
			more    bool
			result  []string
		}{
			{
				options: options("", "", storj.After, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "`", storj.After, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "b", storj.After, 0),
				result:  []string{"b/", "bb", "c"},
			}, {
				options: options("", "c", storj.After, 0),
				result:  []string{},
			}, {
				options: options("", "ca", storj.After, 0),
				result:  []string{},
			}, {
				options: options("", "", storj.After, 1),
				more:    true,
				result:  []string{"a"},
			}, {
				options: options("", "`", storj.After, 1),
				more:    true,
				result:  []string{"a"},
			}, {
				options: options("", "aa", storj.After, 1),
				more:    true,
				result:  []string{"b"},
			}, {
				options: options("", "c", storj.After, 1),
				result:  []string{},
			}, {
				options: options("", "ca", storj.After, 1),
				result:  []string{},
			}, {
				options: options("", "", storj.After, 2),
				more:    true,
				result:  []string{"a", "a/"},
			}, {
				options: options("", "1", storj.After, 2),
				more:    true,
				result:  []string{"a", "a/"},
			}, {
				options: options("", "aa", storj.After, 2),
				more:    true,
				result:  []string{"b", "b/"},
			}, {
				options: options("", "bb", storj.After, 2),
				result:  []string{"c"},
			}, {
				options: options("", "c", storj.After, 2),
				result:  []string{},
			}, {
				options: options("", "ca", storj.After, 2),
				result:  []string{},
			}, {
				options: optionsRecursive("", "", storj.After, 0),
				result:  []string{"a", "a/xa", "a/xaa", "a/xb", "a/xbb", "a/xc", "aa", "b", "b/ya", "b/yaa", "b/yb", "b/ybb", "b/yc", "bb", "c"},
			}, {
				options: options("a", "", storj.After, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "", storj.After, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "xb", storj.After, 0),
				result:  []string{"xbb", "xc"},
			}, {
				options: optionsRecursive("", "a/xbb", storj.After, 5),
				more:    true,
				result:  []string{"a/xc", "aa", "b", "b/ya", "b/yaa"},
			}, {
				options: options("a/", "xaa", storj.After, 2),
				more:    true,
				result:  []string{"xb", "xbb"},
			}, {
				options: options("", "", storj.Forward, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "`", storj.Forward, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "b", storj.Forward, 0),
				result:  []string{"b", "b/", "bb", "c"},
			}, {
				options: options("", "c", storj.Forward, 0),
				result:  []string{"c"},
			}, {
				options: options("", "ca", storj.Forward, 0),
				result:  []string{},
			}, {
				options: options("", "", storj.Forward, 1),
				more:    true,
				result:  []string{"a"},
			}, {
				options: options("", "`", storj.Forward, 1),
				more:    true,
				result:  []string{"a"},
			}, {
				options: options("", "aa", storj.Forward, 1),
				more:    true,
				result:  []string{"aa"},
			}, {
				options: options("", "c", storj.Forward, 1),
				result:  []string{"c"},
			}, {
				options: options("", "ca", storj.Forward, 1),
				result:  []string{},
			}, {
				options: options("", "", storj.Forward, 2),
				more:    true,
				result:  []string{"a", "a/"},
			}, {
				options: options("", "`", storj.Forward, 2),
				more:    true,
				result:  []string{"a", "a/"},
			}, {
				options: options("", "aa", storj.Forward, 2),
				more:    true,
				result:  []string{"aa", "b"},
			}, {
				options: options("", "bb", storj.Forward, 2),
				result:  []string{"bb", "c"},
			}, {
				options: options("", "c", storj.Forward, 2),
				result:  []string{"c"},
			}, {
				options: options("", "ca", storj.Forward, 2),
				result:  []string{},
			}, {
				options: optionsRecursive("", "", storj.Forward, 0),
				result:  []string{"a", "a/xa", "a/xaa", "a/xb", "a/xbb", "a/xc", "aa", "b", "b/ya", "b/yaa", "b/yb", "b/ybb", "b/yc", "bb", "c"},
			}, {
				options: options("a", "", storj.Forward, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "", storj.Forward, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "xb", storj.Forward, 0),
				result:  []string{"xb", "xbb", "xc"},
			}, {
				options: optionsRecursive("", "a/xbb", storj.Forward, 5),
				more:    true,
				result:  []string{"a/xbb", "a/xc", "aa", "b", "b/ya"},
			}, {
				options: options("a/", "xaa", storj.Forward, 2),
				more:    true,
				result:  []string{"xaa", "xb"},
			}, {
				options: options("", "", storj.Backward, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "1", storj.Backward, 0),
				result:  []string{},
			}, {
				options: options("", "b", storj.Backward, 0),
				result:  []string{"a", "a/", "aa", "b"},
			}, {
				options: options("", "c", storj.Backward, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "ca", storj.Backward, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "", storj.Backward, 1),
				more:    true,
				result:  []string{"c"},
			}, {
				options: options("", "1", storj.Backward, 1),
				result:  []string{},
			}, {
				options: options("", "aa", storj.Backward, 1),
				more:    true,
				result:  []string{"aa"},
			}, {
				options: options("", "c", storj.Backward, 1),
				more:    true,
				result:  []string{"c"},
			}, {
				options: options("", "ca", storj.Backward, 1),
				more:    true,
				result:  []string{"c"},
			}, {
				options: options("", "", storj.Backward, 2),
				more:    true,
				result:  []string{"bb", "c"},
			}, {
				options: options("", "`", storj.Backward, 2),
				result:  []string{},
			}, {
				options: options("", "a/", storj.Backward, 2),
				result:  []string{"a"},
			}, {
				options: options("", "bb", storj.Backward, 2),
				more:    true,
				result:  []string{"b/", "bb"},
			}, {
				options: options("", "c", storj.Backward, 2),
				more:    true,
				result:  []string{"bb", "c"},
			}, {
				options: options("", "ca", storj.Backward, 2),
				more:    true,
				result:  []string{"bb", "c"},
			}, {
				options: optionsRecursive("", "", storj.Backward, 0),
				result:  []string{"a", "a/xa", "a/xaa", "a/xb", "a/xbb", "a/xc", "aa", "b", "b/ya", "b/yaa", "b/yb", "b/ybb", "b/yc", "bb", "c"},
			}, {
				options: options("a", "", storj.Backward, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "", storj.Backward, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "xb", storj.Backward, 0),
				result:  []string{"xa", "xaa", "xb"},
			}, {
				options: optionsRecursive("", "b/yaa", storj.Backward, 5),
				more:    true,
				result:  []string{"a/xc", "aa", "b", "b/ya", "b/yaa"},
			}, {
				options: options("a/", "xbb", storj.Backward, 2),
				more:    true,
				result:  []string{"xb", "xbb"},
			}, {
				options: options("", "", storj.Before, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "`", storj.Before, 0),
				result:  []string{},
			}, {
				options: options("", "a", storj.Before, 0),
				result:  []string{},
			}, {
				options: options("", "b", storj.Before, 0),
				result:  []string{"a", "a/", "aa"},
			}, {
				options: options("", "c", storj.Before, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb"},
			}, {
				options: options("", "ca", storj.Before, 0),
				result:  []string{"a", "a/", "aa", "b", "b/", "bb", "c"},
			}, {
				options: options("", "", storj.Before, 1),
				more:    true,
				result:  []string{"c"},
			}, {
				options: options("", "`", storj.Before, 1),
				result:  []string{},
			}, {
				options: options("", "a/", storj.Before, 1),
				result:  []string{"a"},
			}, {
				options: options("", "c", storj.Before, 1),
				more:    true,
				result:  []string{"bb"},
			}, {
				options: options("", "ca", storj.Before, 1),
				more:    true,
				result:  []string{"c"},
			}, {
				options: options("", "", storj.Before, 2),
				more:    true,
				result:  []string{"bb", "c"},
			}, {
				options: options("", "`", storj.Before, 2),
				result:  []string{},
			}, {
				options: options("", "a/", storj.Before, 2),
				result:  []string{"a"},
			}, {
				options: options("", "bb", storj.Before, 2),
				more:    true,
				result:  []string{"b", "b/"},
			}, {
				options: options("", "c", storj.Before, 2),
				more:    true,
				result:  []string{"b/", "bb"},
			}, {
				options: options("", "ca", storj.Before, 2),
				more:    true,
				result:  []string{"bb", "c"},
			}, {
				options: optionsRecursive("", "", storj.Before, 0),
				result:  []string{"a", "a/xa", "a/xaa", "a/xb", "a/xbb", "a/xc", "aa", "b", "b/ya", "b/yaa", "b/yb", "b/ybb", "b/yc", "bb", "c"},
			}, {
				options: options("a", "", storj.Before, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "", storj.Before, 0),
				result:  []string{"xa", "xaa", "xb", "xbb", "xc"},
			}, {
				options: options("a/", "xb", storj.Before, 0),
				result:  []string{"xa", "xaa"},
			}, {
				options: optionsRecursive("", "b/yaa", storj.Before, 5),
				more:    true,
				result:  []string{"a/xbb", "a/xc", "aa", "b", "b/ya"},
			}, {
				options: options("a/", "xbb", storj.Before, 2),
				more:    true,
				result:  []string{"xaa", "xb"},
			},
		} {
			errTag := fmt.Sprintf("%d. %+v", i, tt)

			list, err := db.ListObjects(ctx, bucket.Name, tt.options)

			if assert.NoError(t, err, errTag) {
				assert.Equal(t, tt.more, list.More, errTag)
				assert.Equal(t, tt.result, getObjectPaths(list), errTag)
			}
		}
	})
}

func options(prefix, cursor string, direction storj.ListDirection, limit int) storj.ListOptions {
	return storj.ListOptions{
		Prefix:    prefix,
		Cursor:    cursor,
		Direction: direction,
		Limit:     limit,
	}
}

func optionsRecursive(prefix, cursor string, direction storj.ListDirection, limit int) storj.ListOptions {
	return storj.ListOptions{
		Prefix:    prefix,
		Cursor:    cursor,
		Direction: direction,
		Limit:     limit,
		Recursive: true,
	}
}

func getObjectPaths(list storj.ObjectList) []string {
	names := make([]string, len(list.Items))

	for i, item := range list.Items {
		names[i] = item.Path
	}

	return names
}
