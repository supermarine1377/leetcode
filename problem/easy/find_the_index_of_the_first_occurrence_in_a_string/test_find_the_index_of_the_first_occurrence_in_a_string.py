import unittest
import find_the_index_of_the_first_occurrence_in_a_string as ans

class TestStrstr(unittest.TestCase):
  def test(self):
    self.assertEqual(ans.Solution.strStr(self, "sadbutsad", "sad"), 0)
    self.assertEqual(ans.Solution.strStr(self, "leetcode", "leeto"), -1)
    self.assertEqual(ans.Solution.strStr(self, "mississippi", "issip"), 4)

if __name__ == '__main__':
    unittest.main()