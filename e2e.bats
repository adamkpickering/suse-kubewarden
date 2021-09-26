#!/usr/bin/env bats

@test "accept because there are no palindrome labels" {
  run kwctl run policy.wasm -r test_data/ingress.json --settings-json '{}'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}

@test "deny because there is a palindrome" {
  run kwctl run policy.wasm -r test_data/palindrome.json --settings-json '{}'
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
}
