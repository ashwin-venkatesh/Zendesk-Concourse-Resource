#!/bin/bash

set -e

source $(dirname $0)/common.sh

function assert_equal()
{
  local expected=$1
  local actual=$2

  [ "$(echo $expected)" = "$(echo $actual)" ] \
  && echo pass \
  || echo expected [$expected] but was [$actual]
}

#=====================================
echo "get_id"
json=$(cat <<EOF
{"version":
  {"ref": "1"}
}
EOF
)

actual=$(get_id "$json")
assert_equal 1 $actual

#=====================================
echo "get_id #can seed a value"
json=$(cat <<EOF
{
}
EOF
)

actual=$(get_id "$json")
assert_equal 1 $actual

#=====================================
echo "get_last_date"
expected_date="2017-05-25T19:57:38Z"
json=$(cat <<EOF
{
  "results": [
    {
      "created_at": "$expected_date"
    }
  ],
  "count": 1
}
EOF
)
actual=$(get_last_date "$json")

assert_equal $expected_date "$actual"

#=====================================
echo "get_ids # returns nothing when there are no result"
json=$(cat <<EOF
{
  "results": [],
  "count": 0
}
EOF
)

actual=$(get_ids "$json")
assert_equal "" $actual

#=====================================
echo "get_ids # returns ids"
json=$(cat <<EOF
{
	"results": [
		{
	    	  "id": 1
		},
		{
		  "id": 2
		}
	],
	"count": 2
}
EOF
)

actual=$(get_ids "$json")
expected="1 2"
assert_equal "$expected" "$actual"

#=====================================
echo "transform"
actual=$(transform "1 2")
expected="[ { \"ref\": \"2\" }, { \"ref\": \"1\" } ]"
assert_equal "$expected" "$actual"
