#!/bin/sh -e

# Exit early if any session with my username is found
if who | grep -wq $USER; then
  exit
fi

# Phone numbers
MY_NUMBER='"+xxx"'
NUMBER_OF_BOSS='"+xxx"'

EXCUSES=(
  'Locked out'
  'Pipes broke'
  'Food poisoning'
  'Not feeling well'
)
rand=$[ $RANDOM % ${#EXCUSES[@]} ]

RANDOM_EXCUSE=${EXCUSES[$rand]}
MESSAGE='"Gonna work from home. '"$RANDOM_EXCUSE"'"'

# Send a text message
RESPONSE=`curl -X POST -H "Content-Type: application/json" 'localhost:8080/v2/send' -d '{"message": '"$MESSAGE"', "number": '"$MY_NUMBER"', "recipients": [ '"$NUMBER_OF_BOSS"' ] }'`

# Log errors
if [ $? -gt 0 ]; then
  echo "Failed to send SMS: $RESPONSE"
  exit 1
fi
