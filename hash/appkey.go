package appkey

type IAppkey interface {
	GenerateAppKey(...string) string
}
