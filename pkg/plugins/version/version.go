package version

import (
	"github.com/quarkslab/kdigger/pkg/bucket"
)

const (
	bucketName        = "version"
	bucketDescription = "Version dumps the API server version informations."
)

var bucketAliases = []string{"versions", "v"}

type Bucket struct {
	config bucket.Config
}

func (n Bucket) Run() (bucket.Results, error) {
	res := bucket.NewResults(bucketName)
	v, err := n.config.Client.Discovery().ServerVersion()
	if err != nil {
		return bucket.Results{}, err
	}
	res.SetHeaders([]string{"version", "buildDate", "platform", "goVersion"})
	res.AddContent([]interface{}{v.GitVersion, v.BuildDate, v.Platform, v.GoVersion})
	return *res, nil
}

func Register(b *bucket.Buckets) {
	b.Register(bucket.Bucket{
		Name:        bucketName,
		Description: bucketDescription,
		Aliases:     bucketAliases,
		Factory: func(config bucket.Config) (bucket.Interface, error) {
			return NewVersionBucket(config)
		},
		SideEffects:   false,
		RequireClient: true, // TODO change that to false
	})
}

func NewVersionBucket(config bucket.Config) (*Bucket, error) {
	if config.Client == nil {
		return nil, bucket.ErrMissingClient
	}
	return &Bucket{
		config: config,
	}, nil
}
