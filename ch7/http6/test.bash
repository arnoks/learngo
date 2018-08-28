
curl localhost:8001/list
curl localhost:8001/price?ITEM=shoes
curl localhost:8001/price?ITEM=shirt
curl -d Item="shirt" -d Price="15.0" localhost:8001/add
curl -d Item="pant" -d Price="11.0" localhost:8001/add
curl -d Item="pullover" -d Price="13.0" localhost:8001/add
curl -d Item="skirt" -d Price="15.0" localhost:8001/add
curl -d Item="sweater" -d Price="11.0" localhost:8001/add
curl -d Item="sneaker" -d Price="13.0" localhost:8001/add
curl -X "DELETE" localhost:8001/stock
