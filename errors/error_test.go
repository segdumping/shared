package errors

import "testing"

func TestError(t *testing.T) {
	var err error = New("test error").WithCode(403)

	//error
	t.Logf("error: %s", err.Error())

	//convert
	e, ok := err.(*Error)
	if !ok {
		t.Log("convert error")
		return
	}

	//errCode
	t.Logf("error code: %d", e.errCode)

	//error string
	t.Logf("error string: %s", e.Error())
}
