package conf

import (
	"github.com/stretchr/testify/require"
	"path"
	"testing"
)

func TestFromYaml(t *testing.T) {
	conf, err := FromYaml(path.Join("testdata/config.yml"))
	require.NoError(t, err)

	expected := Conf{
		Keys: Keys{
			Youtube: "youtube-key-xoaowusdqkjdndasdfiuuuum",
		},
		PlayListId: "PLqlu7ZxfTBeiDDYv-a0NTwO6lvc4C3kBv",
		ChunkSize:  50,
		Aws: Aws{
			AccessKey: "AKIApppppp",
			SecretKey: "AKIAssssssssss",
			S3: S3{
				Bucket: "super-s3-bucket",
			},
		},
	}
	require.Equal(t, expected, conf)
}

func TestFromYamlInvalidFile(t *testing.T) {
	conf, err := FromYaml(path.Join("testdata/invalid_config.txt"))
	require.Error(t, err)
	require.Equal(t, Conf{}, conf)
}
