class Solution:
    def isValid(self, s: str) -> bool:
        st = []
        for ss in s:
            if ss == "(" or ss == "{" or ss == "[":
                st.append(ss)
                continue
            if len(st) != 0:
                e = st.pop()
                if e == "(" and ss != ")":
                    return False
                if e == "{" and ss != "}":
                    return False
                if e == "[" and ss != "]":
                    return False
            else:
                return False
        return len(st) == 0
