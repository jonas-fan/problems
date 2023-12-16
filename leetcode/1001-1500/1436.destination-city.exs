defmodule Solution do
  @spec dest_city(paths :: [[String.t]]) :: String.t
  def dest_city(paths) do
    paths
    |> Enum.reduce(%{}, &arrange(&2, &1))
    |> follow(hd(hd(paths)))
  end

  defp arrange(map, [first | rest]) do
    Map.put(map, first, hd(rest))
  end

  defp follow(map, current) do
    case Map.get(map, current) do
      nil -> current
      next -> follow(map, next)
    end
  end
end
