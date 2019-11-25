package welcome

import "testing"

func TestGreeting(test *testing.T) {
  result := Greeting("Hello World!")
  shouldBe := "<b>Hello World!</b>"

  if result != shouldBe {
    test.Errorf("Should return '<b>Hello World!</b>', but returned '%v'", result)
  }
}
