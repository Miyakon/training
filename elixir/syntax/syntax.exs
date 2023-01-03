case {1, 2, 3} do
  {4, 5, 6} ->
    IO.puts("this clause won't match")
  {1, 2, 3} ->
    IO.puts("This clause will match and bind x to 2 in this clause")
  _ ->
    IO.puts("this clause would match any value")
end
