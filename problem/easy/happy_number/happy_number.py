class Solution:
    def isHappy(self, n: int) -> bool:
        s = set()
        while n != 1:
            temp = 0
            m = n
            while True:
                digit = m % 10
                temp += digit * digit
                m = m // 10
                if m < 1:
                  break
            if temp in s:
              return False
            s.add(temp)
            n = temp
        return True

