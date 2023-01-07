class Solution:
  def isAnagram(self, s: str, t: str) -> bool:
    if len(s) != len(t):
      return False
    
    dicta = {}
    dictb = {}
    for ss in s:
      if ss in dicta:
        dicta[ss] = dicta[ss] + 1
      else:
        dicta[ss] = 1
    
    for tt in t:
      if tt in dictb:
        dictb[tt] = dictb[tt] + 1
      else:
        dictb[tt] = 1
      
    return dicta.__eq__(dictb)