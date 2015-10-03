package redis

type FakeRedis struct {
	bytes []byte
	err   error
}

func NewFakeRedis(bytes []byte, err error) *FakeRedis {
	return &FakeRedis{
		bytes: bytes,
		err:   err,
	}
}

func (f *FakeRedis) Info() ([]byte, error) {
	return f.bytes, f.err
}
