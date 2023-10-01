import unittest
from solution import Valid_Parentheses


class TestSolution(unittest.TestCase):
    def test_valid_parentheses(self):
        solution = Valid_Parentheses()
        testcases = [
            ["()", True],
            ["()[]{}", True],
            ["(]", False],
            ["((", False],
            ["]", False],
            [")(){}", False]
        ]
        for s, expected in testcases:
            ans = solution.isValid(s)
            if ans:
                self.assertTrue(expected)
            else:
                self.assertFalse(expected)


if __name__ == "__main__":
    unittest.main()
