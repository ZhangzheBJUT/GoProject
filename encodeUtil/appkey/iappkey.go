package appkey

type IAppKey interface {
	GenerateAppKey(args ...string) string
}
