// クラウドストレージサービスへのアクセス

import (
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

bucket, _ := blob.Openbucket(ctx, "gs://my-bucket")
defer bucket.Close()

reader, err := bucket.NewReader(ctx, "my-file", nil)
defer reader.Close()