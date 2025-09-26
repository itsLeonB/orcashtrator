package config

type ServiceClient struct {
	BillsplittrHost string `split_words:"true" required:"true"`
	CocoonHost      string `split_words:"true" required:"true"`
	DrexHost        string `split_words:"true" required:"true"`
	StortrHost      string `split_words:"true" required:"true"`
}
