import unittest
from solution import Solution


class TestSolution(unittest.TestCase):
    def test_twoSum(self):
        solution = Solution()

        test_cases = [
            ([2, 7, 11, 15], 9, {0, 1}),
            ([3, 2, 4], 6, {1, 2}),
            ([3, 3], 6, {0, 1}),
            ([2, 5, 5, 11], 10, {1, 2}),
            ([1, 2, 3, 4, 5], 10, None)
        ]

        for nums, target, expected in test_cases:
            result_N2 = solution.twoSumN2(nums, target)
            result_N1 = solution.twoSumN1(nums, target)
            self.assertEqual(
                set(result_N2) if result_N2 is not None else None, expected)
            self.assertEqual(
                set(result_N1) if result_N1 is not None else None, expected)

if __name__ == "__main__":
    unittest.main()
