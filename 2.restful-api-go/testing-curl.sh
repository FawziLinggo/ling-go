
# GET categories
curl -X GET http://fawzi.linggo.id:3000/api/categories | jq

# Add data
curl -X POST http://fawzi.linggo.id:3000/api/categories -H 'Content-Type: application/json' -d '{"name":"computer"}' | jq

# Find data By categories id
curl -X GET http://fawzi.linggo.id:3000/api/categories/3 -H 'Content-Type: application/json' | jq

# Update by Categories id
curl -X PUT http://fawzi.linggo.id:3000/api/categories/3 -H 'Content-Type: application/json' -d '{"name":"Kuda"}' | jq  

# Delete by Categories id
curl -X DELETE http://fawzi.linggo.id:3000/api/categories/3 | jq  



############### TESTING VALIDATION
# Add data null
curl -X POST http://fawzi.linggo.id:3000/api/categories -H 'Content-Type: application/json' -d '{"name":""}' | jq
