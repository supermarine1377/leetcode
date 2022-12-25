import unittest
import myans

class Testmyans(unittest.TestCase):
  
  def test_is_anagram(self):
    self.assertFalse(myans.Solution.isAnagram(self, 'a', 'aa'))
    self.assertFalse(myans.Solution.isAnagram(self, 'cat', 'rat'))
    self.assertTrue(myans.Solution.isAnagram(self, 'dog', 'god'))
    self.assertFalse(myans.Solution.isAnagram(self, 'optp', 'optt'))
