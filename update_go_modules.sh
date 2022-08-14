echo "Updating GO modules..."

echo "Updating payment service..."
go get -v -u github.com/Tiger-Coders/tigerlily-payment@main

echo "Updating inventory service..."
go get -v -u github.com/Tiger-Coders/tigerlily-inventories@main

echo "Updating cache library..."
go get -u -v github.com/ZAF07/tigerlily-e-bakery-cache/redis-cache-manager@main

echo "DONE UPDATING GO MODULES"