defmodule Solution do
    @spec two_sum(nums :: [integer], target :: integer) :: [integer]
    def two_sum(nums, target) do
        two_sum(nums, target, 0, %{})
    end

    defp two_sum([num | nums], target, index, map) do
        case Map.get(map, target - num) do
            nil -> two_sum(nums, target, index + 1, Map.put(map, num, index))
            val -> [index, val]
        end
    end
end
