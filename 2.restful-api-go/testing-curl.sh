
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
#############################################

############### TESTING VALIDATION
# GET data null
curl -X GET http://fawzi.linggo.id:3000/api/categories/100 -H 'Content-Type: application/json' | jq

# Update by Categories id null
curl -X PUT http://fawzi.linggo.id:3000/api/categories/90 -H 'Content-Type: application/json' -d '{"name":"Kuda"}' | jq  
#############################################

############### AUTH
# GET categories AUTH
curl -X GET http://fawzi.linggo.id:3000/api/categories -H 'X-API-Key:  3<vnqdO+vQ5fr'  | jq

# Add data AUTH
curl -X POST http://fawzi.linggo.id:3000/api/categories -H 'Content-Type: application/json' -H 'X-API-Key:  3<vnqdO+vQ5fr' -d '{"name":"computer"}' | jq

# Find data By categories id AUTH
curl -X GET http://fawzi.linggo.id:3000/api/categories/1 -H 'Content-Type: application/json' -H 'X-API-Key:  3<vnqdO+vQ5fr' | jq

# Update by Categories id AUTH
curl -X PUT http://fawzi.linggo.id:3000/api/categories/3 -H 'Content-Type: application/json' -H 'X-API-Key:  3<vnqdO+vQ5fr' -d '{"name":"Kuda"}' | jq  

# Delete by Categories id AUTH
curl -X DELETE http://fawzi.linggo.id:3000/api/categories/3 -H 'X-API-Key:  3<vnqdO+vQ5fr' | jq