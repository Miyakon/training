string = "hello"
IO.puts is_binary(string)

string = "hełło"
IO.puts byte_size(string)
IO.puts String.length(string)

IO.puts ?a
IO.puts ?ł

IO.puts String.codepoints("hello")
