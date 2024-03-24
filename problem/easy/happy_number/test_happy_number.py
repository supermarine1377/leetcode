import unittest
import happy_number

class TestHappyNumber(unittest.TestCase):
  def test_happy_number(self):
    self.assertEqual(happy_number.Solution.isHappy(self, 19), True)
    self.assertEqual(happy_number.Solution.isHappy(self, 2), False)
    
if __name__ == '__main__':
    unittest.main()