IO.puts("No.1")
case {1, 2, 3} do
  {4, 5, 6} ->
    IO.puts("this clause won't match")
  {1, _, 3} ->
    IO.puts("This clause will match and bind x to 2 in this clause")
  _ ->
    IO.puts("this clause would match any value")
end

IO.puts("\nNo.2")
case {1, 2, 3} do
  {1, x, 3} when x > 0 ->
    IO.puts "Will match"
  _ ->
    IO. puts "Would match, if guard condition were not satisfied"
end

IO.puts "\nNo3"
case 1 do
  x when x < 0 -> IO.puts "Won't match"
  x -> IO.puts "Got #{x}"
end

IO.puts "\nNo4"
f = fn
  x, y when x > 0 -> x + y
  x, y -> x * y
end

IO.puts f.(1, 3)
IO.puts f.(-1, 5)
