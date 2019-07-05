#!/usr/bin/env python3

import requests
import re
from requests.exceptions import HTTPError

for url in ['https://orders.olkb.com']:
    try:
        response = requests.get(url)

        # If the response was successful, no Exception will be raised
        response.raise_for_status()
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')
    except Exception as err:
        print(f'Other error occurred: {err}')
    else:
        #print('Success!')
        pass

# Need to grab extra 4 chars at begining so it does not count orders
# that were combined with others
p = re.compile('<li>10000\d{4}')
temp = p.findall(response.text)
# Removes the extra first 4 chars of each entry in the list
orders = [s[4:] for s in temp]
#print(orders)
order_number='100007000' // put your own order number
if order_number in orders:
	order_position = orders.index(order_number) + 1
	print(f'olkb position: {order_position}')

print('---')

order_count = len(orders)
print(f'Total orders: {order_count}')
print(f'Order#: {order_number}')




        