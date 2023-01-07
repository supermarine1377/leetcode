import unittest
import bestans

class Testbestans(unittest.TestCase):
  
  def test_contains_duplicate(self):
    self.assertTrue(bestans.Solution.containsDepulicate([1, 2, 4, 2]))
    self.assertFalse(bestans.Solution.containsDepulicate([0, 2, 4, 3]))