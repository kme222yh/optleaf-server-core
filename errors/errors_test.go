package errors


import (
    "testing"
)


func TestError(t *testing.T) {
    testCases := [...] string {
        "testtest",
        "",
        "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        "This is Error",
        "12345678",
    }

    for i := 0; i < len(testCases); i++ {
        val := testCases[i]
        result := New(val)
        if(result.Error() != val) {
            t.Errorf("errors.New(\"%v\").Error() = \"%v\", want %v", val, result.Error(), val)
        }
    }
}
