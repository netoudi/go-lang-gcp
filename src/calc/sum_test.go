package calc

import "testing"

func TestSum(test *testing.T) {
  result := Sum(5, 5)
  shouldBe := 10

  if result != shouldBe {
    test.Errorf("Should return 10, but returned %v", result);
  }
}
