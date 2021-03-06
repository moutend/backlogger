#!/usr/bin/env ruby

# This code generates methods.go.
# 
# Go's AST package is too much for this task, so I decided to write simple Ruby script.

functions = []

ARGF.each do |line|
  # Extract function statements.
  next if !line.start_with? "func (c *Client)"

  line = line.sub("func (c *Client) ", "")

  # Skip private methods
  next if /[a-z]/.match line[0]

  p_begin = line.index "("
  p_end = line.index ")"

  f_name = line[0 ... p_begin]
  f_args = line[(p_begin  + 1) .. (p_end - 1)]
  f_ret = line[(p_end + 2) ... -3]
  f_params = []

  f_args.split(",").each do |arg|
    f_params.push(arg.split(" "))
  end

  next if f_name == "SetHTTPClient"

  functions.push({
    :name => f_name,
    :args => f_args,
    :ret => f_ret,
    :params => f_params,
  })
end

puts <<END
// Code generated by "generate.rb"; DO NOT EDIT.
package backlog

import (
  . "github.com/moutend/go-backlog/pkg/types"
)
END

functions.each do |f|
  params = f[:params].map do |v|
    next v[0] if v.size == 1
    v[1].start_with?("...") ? "#{v[0]}..." : v[0]
  end.join(", ")

  puts ""
  puts "func #{f[:name]}(#{f[:args]}) #{f[:ret]} {"
  puts "  return backlogClient.#{f[:name]}(#{params})"
  puts "}"
end
