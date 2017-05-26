#!/bin/bash

function get_id()
{
  local version=$1
  echo $version | jq -r '.version.ref // "1"'

}

function get_ids()
{
  local json=$1
  echo $json | jq -r '.results[].id'
}

function get_last_date()
{
  local json_data=$1

  echo $json_data | jq -r '.results[0].created_at // ""'
}

function transform()
{
  local result=""
  local list=$1
  local ids=$(echo $list | tr " " "\n")

  for id in $ids
  do
    result="{ \"ref\": \"$id\" }, $result"
  done
  result=${result%, }
  echo "[ $result ]"
}

function find_by_id()
{
  subdomain=$1
  user_name=$2
  password=$3
  id=$4

  curl "https://$subdomain.zendesk.com/api/v2/search.json" \
  -G --data-urlencode "query=$id" \
  -v -u $user_name:$password
}

function find_after_timestamp()
{
  subdomain=$1
  user_name=$2
  password=$3
  date=$4

  curl "https://$subdomain.zendesk.com/api/v2/search.json" \
  -G --data-urlencode "query=status:new created>$date" \
  -v -u $user_name:$password
}
