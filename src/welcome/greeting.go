package welcome

import "fmt"

func Greeting(message string) string {
  return fmt.Sprintf("<b>%s</b>", message)
}
