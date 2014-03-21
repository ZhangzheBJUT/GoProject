package appkey

type TestAppkey struct {
}

func (this *TestAppkey) GenerateAppKey(args ...string) string {
	return "testAppKey"
}
