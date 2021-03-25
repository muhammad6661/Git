package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampPaymentSources(){
	card:=[]types.Card{
		{
		Balance: 10_000_00,
		Active: true,
	   },
	     {
		   Balance: 10_000_00,
          Active: false,
		},
		{
		   Balance: 0,
          Active: true,
		},
		{
		   Balance: 20_000_00,
          Active: true,
		},
	
}
payment:=PaymentSources(card)
  for _,operation:=range payment{
	  fmt.Println(operation.Number)
  }
  //Output: 
  //1000000
  //2000000
}
