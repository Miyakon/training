IO.puts "No.1"

cond do
  2 + 2 == 5 ->
    IO.puts "This will not be true"
  2 * 2 == 3 ->
    IO.puts "Nor this"
  true ->
    IO.puts "But this will"
end
