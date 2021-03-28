package response

type RegLock struct {
	Response
}

func (response RegLock) Success() bool {
	return response.ResponseCode == 200
}
