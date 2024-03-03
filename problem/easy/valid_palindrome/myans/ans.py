import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        if len(s) == 1:
            return True
        
        s = s.lower()

        regex = re.compile('[^a-z0-9]')
        s = regex.sub('', s)

        if len(s) == 1:
            return True
        

        l = 0
        r = len(s) - 1
        while r >= l:
            if s[l] != s[r]:
                return False
            r = r - 1
            l = l + 1
        return True
