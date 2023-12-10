defmodule Solution do
    @spec is_palindrome(x :: integer) :: boolean
    def is_palindrome(x) do
        case x do
            x when x < 0 -> false
            x when x == 0 -> true
            x -> is_palindrome([], x)
        end
    end

    defp is_palindrome(nums, x) do
        case x do
            0 -> nums == Enum.reverse(nums)
            _ -> is_palindrome([rem(x, 10) | nums], div(x, 10))
        end
    end
end
