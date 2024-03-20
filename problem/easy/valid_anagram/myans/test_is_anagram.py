import unittest
import is_anagram as is_anagram

class Testmyans(unittest.TestCase):
  
  def test_is_anagram(self):
    self.assertFalse(is_anagram.Solution.isAnagram(self, 'a', 'aa'))
    self.assertFalse(is_anagram.Solution.isAnagram(self, 'cat', 'rat'))
    self.assertTrue(is_anagram.Solution.isAnagram(self, 'dog', 'god'))
    self.assertFalse(is_anagram.Solution.isAnagram(self, 'optp', 'optt'))


if __name__ == "__main__":
    unittest.main()
