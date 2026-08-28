class Solution:
    def search(self, nums: List[int], target: int) -> int:
        inicio = 0
        fim = len(nums) - 1

        while inicio <= fim:
            meio = (inicio + fim) // 2
            if target == nums[meio]:
                return meio

            elif target > nums[meio]:
                inicio = meio + 1

            else:
                fim  = meio - 1
                
        return -1
