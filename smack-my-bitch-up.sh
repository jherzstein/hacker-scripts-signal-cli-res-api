#!/bin/sh -e

# Exit early if no sessions with my username are found
if ! who | grep -wq $USER; then
  exit
fi

# Phone numbers
MY_NUMBER='"+xxx"'
HER_NUMBER='"+xxx"'
echo $MY_NUMBER 
echo $HER_NUMBER

REASONS=(
  'Working hard'
  'Gotta ship this feature'
  'Someone fucked the system again'
)
rand=$[ $RANDOM % ${#REASONS[@]} ]

RANDOM_REASON=${REASONS[$rand]}
MESSAGE='"Late at work. '"$RANDOM_REASON"'"'
echo $MESSAGE

# Send a text message
RESPONSE=`curl -X POST -H "Content-Type: application/json" 'localhost:8080/v2/send' -d '{"message": '"$MESSAGE"', "number": '"$MY_NUMBER"', "recipients": [ '"$HER_NUMBER"' ] }'`

# Log errors
if [ $? -gt 0 ]; then
  echo "Failed to send SMS: $RESPONSE"
  exit 1
fi
