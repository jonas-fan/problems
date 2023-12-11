defmodule Solution do
  @spec find_special_integer(arr :: [integer]) :: integer
  def find_special_integer(arr) do
    find_special_integer(arr, div(length(arr), 4), %{})
  end

  defp find_special_integer([first | rest], count, map) do
    map = Map.update(map, first, 1, fn val -> val + 1 end)

    case Map.get(map, first) do
      x when x > count -> first
      _ -> find_special_integer(rest, count, map)
    end
  end
end
