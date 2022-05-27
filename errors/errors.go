/*  Usage

err := errors.New("An error occuerd.")
err.Error() => "An error occuerd."

*/


package errors


type err struct {
    text string
}


func New(text string) *err {
    return &err{text: text}
}


func (e *err) Error() string {
    return e.text
}
