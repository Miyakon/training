IO.puts "# case"
case {1, 2, 3} do
  {4, 5, 6} ->
    IO.puts "No"
  {1, x, 3} ->
    IO.puts "Yes"
  _ ->
    IO.puts "___"
end

x = 1
case 10 do
  ^x -> "no"
  __ -> IO.puts("Yes")
end

case {1, 2, 3} do
  {1, x, 3} when x > 0 -> # when ガード．この句はエラーとして処理されない．
    IO.puts("Yes")
  _ ->
    IO.puts("no")
end

IO.puts("# cond")
cond do
  2 + 3 == 5 ->
    IO.puts("Yes")
  1 + 1 == 2 ->
    IO.puts("no")
  true ->
    IO.puts("no")
end

IO.puts("# if/unless")

if true do
  IO.puts("Yes")
end

unless true  do
  IO.puts("No")
end
