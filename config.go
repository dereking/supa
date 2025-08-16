package supabase

type Config struct {
	ApiURL          string
	ApiKey          string
	Bucket          string
	ProjectRef      string
	Debug           bool
	PostgresOptions []PostgresOption
	AuthOptions     []AuthOption
	StorageOptions  []StorageOption
}
