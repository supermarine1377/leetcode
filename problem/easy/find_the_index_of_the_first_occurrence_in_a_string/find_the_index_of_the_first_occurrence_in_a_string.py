class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if len(needle) == 0:
            return 0
        l, r = 0, len(haystack)

        while l < r:
            if haystack[l] != needle[0]:
                l = l + 1
            elif haystack[r - 1] != needle[-1]: 
                r = r - 1
            else:
                if haystack[l:l + len(needle)] == needle:
                    return l
                else:
                    l = l + 1
        return -1
        