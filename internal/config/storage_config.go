package config

type Storage struct {
	BucketNameExpenseBill string `split_words:"true" required:"true"`
}
