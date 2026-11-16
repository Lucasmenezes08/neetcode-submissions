class Solution:
    def isPalindrome(self, s: str) -> bool:
        cleanWord = re.sub(r'[^a-z0-9]' ,'', s.lower())
        reverseWord = ""

        for i in range(len(cleanWord) - 1, -1, -1):
            reverseWord += cleanWord[i]

        if cleanWord == reverseWord:
            return True

        return False

        