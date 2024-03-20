import unittest
import myans

class Testbestans(unittest.TestCase):
  
    def test_remove_duplicates_from_sorted_array(self):
        self.assertEqual(myans.Solution.removeDuplicates(self, [1, 1, 2]), [1,2])
        self.assertEqual(myans.Solution.removeDuplicates(self, [1, 1, 2, 2, 2, 3, 3, 4]), [1,2,3,4])
        self.assertEqual(myans.Solution.removeDuplicates(self, [1, 1, 1]), [1])

if __name__ == "__main__":
    unittest.main()
