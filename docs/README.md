# ArcherDAO API

| Method                                                           | HTTP request                       | Description                                 |
| ---------------------------------------------------------------- | ---------------------------------- | ------------------------------------------- |
| [**constructor_post**](DefaultApi.md#constructor_post)           | **POST** /v1/constructor           | constructor()                               |
| [**get_price_get**](DefaultApi.md#get_price_get)                 | **GET** /v1/{address}/getPrice     | getPrice(address,bytes) [read only]         |
| [**get_price_post**](DefaultApi.md#get_price_post)               | **POST** /v1/{address}/getPrice       | getPrice(address,bytes) [read only]         |
| [**query_all_prices_get**](DefaultApi.md#query_all_prices_get)   | **GET** /v1/{address}/queryAllPrices  | queryAllPrices(bytes,uint256[]) [read only] |
| [**query_all_prices_post**](DefaultApi.md#query_all_prices_post) | **POST** /v1/{address}/queryAllPrices | queryAllPrices(bytes,uint256[]) [read only] |
| [**query_get**](DefaultApi.md#query_get)                         | **GET** /v1/{address}/query           | query(bytes,uint256[]) [read only]          |
| [**query_post**](DefaultApi.md#query_post)                       | **POST** /v1/{address}/query          | query(bytes,uint256[]) [read only]          |

# **constructor_post**

> ConstructorOutputs constructor_post(body, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber, fold_register=fold_register)

constructor()

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
body = swagger_client.ConstructorInputs() # ConstructorInputs |
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)
fold_sync = true # bool | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync) (optional) (default to true)
fold_call = true # bool | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call) (optional)
fold_privatefrom = 'fold_privatefrom_example' # str | Private transaction sender (header: x-manifold-privatefrom) (optional)
fold_privatefor = 'fold_privatefor_example' # str | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor) (optional)
fold_blocknumber = 'fold_blocknumber_example' # str | The target block number for eth_call requests. One of 'earliest/latest/pending', a number or a hex string (header: x-manifold-blocknumber) (optional)
fold_register = 'fold_register_example' # str | Register the installed contract on a friendly path (overwrites existing) (header: x-manifold-register) (optional)

try:
    # constructor()
    api_response = api_instance.constructor_post(body, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber, fold_register=fold_register)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->constructor_post: %s\n" % e)
```

### Parameters

| Name                 | Type                                          | Description                                                                                                                                          | Notes                                                             |
| -------------------- | --------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **body**             | [**ConstructorInputs**](ConstructorInputs.md) |                                                                                                                                                      |
| **fold_from**        | **str**                                       | The &#x27;from&#x27; address (header: x-manifold-from)                                                                                               | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**    | **int**                                       | Ether value to send with the transaction (header: x-manifold-ethvalue)                                                                               | [optional]                                                        |
| **fold_gas**         | **int**                                       | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas)                                                               | [optional]                                                        |
| **fold_gasprice**    | **int**                                       | Gas Price offered (header: x-manifold-gasprice)                                                                                                      | [optional]                                                        |
| **fold_sync**        | **bool**                                      | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync)                                                  | [optional][default to true]                                       |
| **fold_call**        | **bool**                                      | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call)                          | [optional]                                                        |
| **fold_privatefrom** | **str**                                       | Private transaction sender (header: x-manifold-privatefrom)                                                                                          | [optional]                                                        |
| **fold_privatefor**  | **str**                                       | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor)                                                  | [optional]                                                        |
| **fold_blocknumber** | **str**                                       | The target block number for eth_call requests. One of &#x27;earliest/latest/pending&#x27;, a number or a hex string (header: x-manifold-blocknumber) | [optional]                                                        |
| **fold_register**    | **str**                                       | Register the installed contract on a friendly path (overwrites existing) (header: x-manifold-register)                                               | [optional]                                                        |

### Return type

[**ConstructorOutputs**](ConstructorOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_price_get**

> GetPriceOutputs get_price_get(address, contract_address, data, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)

getPrice(address,bytes) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
address = 'address_example' # str | The contract address
contract_address = 'contract_address_example' # str | address: contract to query
data = 'data_example' # str | bytes: the bytecode for the contract call
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)

try:
    # getPrice(address,bytes) [read only]
    api_response = api_instance.get_price_get(address, contract_address, data, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->get_price_get: %s\n" % e)
```

### Parameters

| Name                 | Type    | Description                                                                            | Notes                                                             |
| -------------------- | ------- | -------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **address**          | **str** | The contract address                                                                   |
| **contract_address** | **str** | address: contract to query                                                             |
| **data**             | **str** | bytes: the bytecode for the contract call                                              |
| **fold_from**        | **str** | The &#x27;from&#x27; address (header: x-manifold-from)                                 | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**    | **int** | Ether value to send with the transaction (header: x-manifold-ethvalue)                 | [optional]                                                        |
| **fold_gas**         | **int** | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) | [optional]                                                        |
| **fold_gasprice**    | **int** | Gas Price offered (header: x-manifold-gasprice)                                        | [optional]                                                        |

### Return type

[**GetPriceOutputs**](GetPriceOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_price_post**

> GetPriceOutputs get_price_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)

getPrice(address,bytes) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
body = swagger_client.GetPriceInputs() # GetPriceInputs |
address = 'address_example' # str | The contract address
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)
fold_sync = true # bool | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync) (optional) (default to true)
fold_call = true # bool | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call) (optional)
fold_privatefrom = 'fold_privatefrom_example' # str | Private transaction sender (header: x-manifold-privatefrom) (optional)
fold_privatefor = 'fold_privatefor_example' # str | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor) (optional)
fold_blocknumber = 'fold_blocknumber_example' # str | The target block number for eth_call requests. One of 'earliest/latest/pending', a number or a hex string (header: x-manifold-blocknumber) (optional)

try:
    # getPrice(address,bytes) [read only]
    api_response = api_instance.get_price_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->get_price_post: %s\n" % e)
```

### Parameters

| Name                 | Type                                    | Description                                                                                                                                          | Notes                                                             |
| -------------------- | --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **body**             | [**GetPriceInputs**](GetPriceInputs.md) |                                                                                                                                                      |
| **address**          | **str**                                 | The contract address                                                                                                                                 |
| **fold_from**        | **str**                                 | The &#x27;from&#x27; address (header: x-manifold-from)                                                                                               | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**    | **int**                                 | Ether value to send with the transaction (header: x-manifold-ethvalue)                                                                               | [optional]                                                        |
| **fold_gas**         | **int**                                 | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas)                                                               | [optional]                                                        |
| **fold_gasprice**    | **int**                                 | Gas Price offered (header: x-manifold-gasprice)                                                                                                      | [optional]                                                        |
| **fold_sync**        | **bool**                                | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync)                                                  | [optional][default to true]                                       |
| **fold_call**        | **bool**                                | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call)                          | [optional]                                                        |
| **fold_privatefrom** | **str**                                 | Private transaction sender (header: x-manifold-privatefrom)                                                                                          | [optional]                                                        |
| **fold_privatefor**  | **str**                                 | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor)                                                  | [optional]                                                        |
| **fold_blocknumber** | **str**                                 | The target block number for eth_call requests. One of &#x27;earliest/latest/pending&#x27;, a number or a hex string (header: x-manifold-blocknumber) | [optional]                                                        |

### Return type

[**GetPriceOutputs**](GetPriceOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **query_all_prices_get**

> QueryAllPricesOutputs query_all_prices_get(address, script, input_locations, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)

queryAllPrices(bytes,uint256[]) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
address = 'address_example' # str | The contract address
script = 'script_example' # str | bytes: the compiled bytecode for the series of function calls to get the final price
input_locations = 'input_locations_example' # str | uint256[]: index locations within the script to insert input amounts dynamically
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)

try:
    # queryAllPrices(bytes,uint256[]) [read only]
    api_response = api_instance.query_all_prices_get(address, script, input_locations, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->query_all_prices_get: %s\n" % e)
```

### Parameters

| Name                | Type    | Description                                                                            | Notes                                                             |
| ------------------- | ------- | -------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **address**         | **str** | The contract address                                                                   |
| **script**          | **str** | bytes: the compiled bytecode for the series of function calls to get the final price   |
| **input_locations** | **str** | uint256[]: index locations within the script to insert input amounts dynamically       |
| **fold_from**       | **str** | The &#x27;from&#x27; address (header: x-manifold-from)                                 | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**   | **int** | Ether value to send with the transaction (header: x-manifold-ethvalue)                 | [optional]                                                        |
| **fold_gas**        | **int** | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) | [optional]                                                        |
| **fold_gasprice**   | **int** | Gas Price offered (header: x-manifold-gasprice)                                        | [optional]                                                        |

### Return type

[**QueryAllPricesOutputs**](QueryAllPricesOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **query_all_prices_post**

> QueryAllPricesOutputs query_all_prices_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)

queryAllPrices(bytes,uint256[]) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
body = swagger_client.QueryAllPricesInputs() # QueryAllPricesInputs |
address = 'address_example' # str | The contract address
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)
fold_sync = true # bool | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync) (optional) (default to true)
fold_call = true # bool | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call) (optional)
fold_privatefrom = 'fold_privatefrom_example' # str | Private transaction sender (header: x-manifold-privatefrom) (optional)
fold_privatefor = 'fold_privatefor_example' # str | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor) (optional)
fold_blocknumber = 'fold_blocknumber_example' # str | The target block number for eth_call requests. One of 'earliest/latest/pending', a number or a hex string (header: x-manifold-blocknumber) (optional)

try:
    # queryAllPrices(bytes,uint256[]) [read only]
    api_response = api_instance.query_all_prices_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->query_all_prices_post: %s\n" % e)
```

### Parameters

| Name                 | Type                                                | Description                                                                                                                                          | Notes                                                             |
| -------------------- | --------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **body**             | [**QueryAllPricesInputs**](QueryAllPricesInputs.md) |                                                                                                                                                      |
| **address**          | **str**                                             | The contract address                                                                                                                                 |
| **fold_from**        | **str**                                             | The &#x27;from&#x27; address (header: x-manifold-from)                                                                                               | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**    | **int**                                             | Ether value to send with the transaction (header: x-manifold-ethvalue)                                                                               | [optional]                                                        |
| **fold_gas**         | **int**                                             | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas)                                                               | [optional]                                                        |
| **fold_gasprice**    | **int**                                             | Gas Price offered (header: x-manifold-gasprice)                                                                                                      | [optional]                                                        |
| **fold_sync**        | **bool**                                            | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync)                                                  | [optional][default to true]                                       |
| **fold_call**        | **bool**                                            | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call)                          | [optional]                                                        |
| **fold_privatefrom** | **str**                                             | Private transaction sender (header: x-manifold-privatefrom)                                                                                          | [optional]                                                        |
| **fold_privatefor**  | **str**                                             | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor)                                                  | [optional]                                                        |
| **fold_blocknumber** | **str**                                             | The target block number for eth_call requests. One of &#x27;earliest/latest/pending&#x27;, a number or a hex string (header: x-manifold-blocknumber) | [optional]                                                        |

### Return type

[**QueryAllPricesOutputs**](QueryAllPricesOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **query_get**

> QueryOutputs query_get(address, script, input_locations, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)

query(bytes,uint256[]) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
address = 'address_example' # str | The contract address
script = 'script_example' # str | bytes: the compiled bytecode for the series of function calls to get the final price
input_locations = 'input_locations_example' # str | uint256[]: index locations within the script to insert input amounts dynamically
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)

try:
    # query(bytes,uint256[]) [read only]
    api_response = api_instance.query_get(address, script, input_locations, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->query_get: %s\n" % e)
```

### Parameters

| Name                | Type    | Description                                                                            | Notes                                                             |
| ------------------- | ------- | -------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **address**         | **str** | The contract address                                                                   |
| **script**          | **str** | bytes: the compiled bytecode for the series of function calls to get the final price   |
| **input_locations** | **str** | uint256[]: index locations within the script to insert input amounts dynamically       |
| **fold_from**       | **str** | The &#x27;from&#x27; address (header: x-manifold-from)                                 | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**   | **int** | Ether value to send with the transaction (header: x-manifold-ethvalue)                 | [optional]                                                        |
| **fold_gas**        | **int** | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) | [optional]                                                        |
| **fold_gasprice**   | **int** | Gas Price offered (header: x-manifold-gasprice)                                        | [optional]                                                        |

### Return type

[**QueryOutputs**](QueryOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **query_post**

> QueryOutputs query_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)

query(bytes,uint256[]) [read only]

### Example

```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
body = swagger_client.QueryInputs() # QueryInputs |
address = 'address_example' # str | The contract address
fold_from = '0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147' # str | The 'from' address (header: x-manifold-from) (optional) (default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147)
fold_ethvalue = 56 # int | Ether value to send with the transaction (header: x-manifold-ethvalue) (optional)
fold_gas = 56 # int | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas) (optional)
fold_gasprice = 56 # int | Gas Price offered (header: x-manifold-gasprice) (optional)
fold_sync = true # bool | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync) (optional) (default to true)
fold_call = true # bool | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call) (optional)
fold_privatefrom = 'fold_privatefrom_example' # str | Private transaction sender (header: x-manifold-privatefrom) (optional)
fold_privatefor = 'fold_privatefor_example' # str | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor) (optional)
fold_blocknumber = 'fold_blocknumber_example' # str | The target block number for eth_call requests. One of 'earliest/latest/pending', a number or a hex string (header: x-manifold-blocknumber) (optional)

try:
    # query(bytes,uint256[]) [read only]
    api_response = api_instance.query_post(body, address, fold_from=fold_from, fold_ethvalue=fold_ethvalue, fold_gas=fold_gas, fold_gasprice=fold_gasprice, fold_sync=fold_sync, fold_call=fold_call, fold_privatefrom=fold_privatefrom, fold_privatefor=fold_privatefor, fold_blocknumber=fold_blocknumber)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->query_post: %s\n" % e)
```

### Parameters

| Name                 | Type                              | Description                                                                                                                                          | Notes                                                             |
| -------------------- | --------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| **body**             | [**QueryInputs**](QueryInputs.md) |                                                                                                                                                      |
| **address**          | **str**                           | The contract address                                                                                                                                 |
| **fold_from**        | **str**                           | The &#x27;from&#x27; address (header: x-manifold-from)                                                                                               | [optional][default to 0xf62fbd52d83fded97386fb9dfcb0bfbb3c164147] |
| **fold_ethvalue**    | **int**                           | Ether value to send with the transaction (header: x-manifold-ethvalue)                                                                               | [optional]                                                        |
| **fold_gas**         | **int**                           | Gas to send with the transaction (auto-calculated if not set) (header: x-manifold-gas)                                                               | [optional]                                                        |
| **fold_gasprice**    | **int**                           | Gas Price offered (header: x-manifold-gasprice)                                                                                                      | [optional]                                                        |
| **fold_sync**        | **bool**                          | Block the HTTP request until the tx is mined (does not store the receipt) (header: x-manifold-sync)                                                  | [optional][default to true]                                       |
| **fold_call**        | **bool**                          | Perform a read-only call with the same parameters that would be used to invoke, and return result (header: x-manifold-call)                          | [optional]                                                        |
| **fold_privatefrom** | **str**                           | Private transaction sender (header: x-manifold-privatefrom)                                                                                          | [optional]                                                        |
| **fold_privatefor**  | **str**                           | Private transaction recipients (comma separated or multiple params) (header: x-manifold-privatefor)                                                  | [optional]                                                        |
| **fold_blocknumber** | **str**                           | The target block number for eth_call requests. One of &#x27;earliest/latest/pending&#x27;, a number or a hex string (header: x-manifold-blocknumber) | [optional]                                                        |

### Return type

[**QueryOutputs**](QueryOutputs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
