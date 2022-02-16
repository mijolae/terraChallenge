## Terra Code Challenge

### by Mijolae Wright

<br>
Hello!   
 
This is my github repo of the coding challenge given to for the Dev Relations role at Terra. The challenge was to build a liquidation monitor for people who deposit collateral to specific pools. In addition, the core requirements were:
<br>
> 1. There is a ‘alpha’ service Observer Docs (Terra_Observer_Integration_Guide.pdf) that you can connect to and receive feeds (with an example site here - https://observer.starport.services/ .. you want the ‘new-block’ feed)
>
> 2. Identifies key events (for example the ‘borrow_stable’ message )
Creates some form of totals that can be accessed via a web browser via /totals
In its simplest form it should output some json something similar to 
{ prices: 
[ 
{price: NNN, volume: NNN}, 
{price: NNN, volume: NNN} 
] 
}   
	This should be requestable from a regular browser like GET /prices

### My Current Solution

So far, I have a program that connects to the websocket and places the response into a self-defined structure. I can access every part of the response. When running the program, you will see the output containing the actions that have occurred in the past five seconds, as well as the current supply of each coin in the Terra ecosystem (UST, USD, etc). In about six hours of work, I've completed the first core requirement and part of the second.

### How to Run

I prefer to run from the binary and given that there is a Makefile, it is fairly simple to run. Run the following in your terminal to start the program:

```
make install
terraTest
```

### Improvements to make

1. Adding a webpage for the JSON output would be my first improvement. I can use the `html/template` Golang library to pass the slice of Supply and the slice of PastActions to a webpage. Golang objects can be passed into a webpage using `{{.}}` syntax.
2. To get liquidation prices, one question I need answered is how to we lock to one specific address? With my current solution, we would be getting the response of the websocket for different addresses. Once we can peg the collateral to one specific address, we'll be able to calculate the liquidation price. I would do that by first looking for the `deposit_stable` action and obtain the collateral/denom put up by the account owner. Then we can grab the supply of that particaular denom from the Supply slice within my program. With those two numbers we can calculate the liquidation price.
