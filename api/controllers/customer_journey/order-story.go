package customer_journey

import (
	"fmt"
)

// Glossary
// 3PL "Third Party Logistics"

type Status int

const (
	IN_PROGRESS Status = iota
	FRAUD_CHECKED
	SUBMITTED
	CANCELLED
	READY
	SEND_TO_3PL
	SENT_TO_3PL_FAILED
	ACK_RECEIVED_FROM_3PL
	ACK_NOT_RECEIVED
	SHIPPED
	PAYMENT_CAPTURED
	PAYMENT_NOT_CAPTURED
	SENT_TO_POSLOG
	RETURNED
)

type OrderStatus struct {
    Description string
	Id Status
	Step int
}

/*
I want to tell a story about the order and show how 
the order's status can change over time. 
*/


func Order() string { 
	message := story()
	return message;
}

func story() string {
	//order := OrderStatus{ Description: "The order is in progress", Id: IN_PROGRESS }

	description := GetStatusDescriptions(FRAUD_CHECKED)

	message := fmt.Sprintf("Order status: %s", description)
	return message
}
 

func GetStatusDescriptions( Id Status) string {
	var order_status_list []OrderStatus

	// 1
	submitted_description := "This is when an order is placed. Customer can cancel the order within 1 hour buffer time."
	submitted := OrderStatus{ Description: submitted_description, Id: SUBMITTED, Step: 1 }
	order_status_list = append(order_status_list, submitted)
	// 2
	fraud_description := "At the time of placing the Order, we talk to a third party payments service which will conduct the payment authorization and detect fraud."
	fraud := OrderStatus{ Description: fraud_description, Id: FRAUD_CHECKED, Step: 2 }
	order_status_list = append(order_status_list, fraud)
	// 3
	ready_description := "Once an order has passed the time buffer, it should briefly move to the ready status where it should be transition to another status."
	ready := OrderStatus{ Description: ready_description, Id: READY, Step: 3 }
	order_status_list = append(order_status_list, ready)
	// 4 - success
	send_to_3pl_description := "The order information is being sent to the warehouse. This process should take no more than 15 mins."
	send_to_3pl := OrderStatus{ Description: send_to_3pl_description, Id: SEND_TO_3PL, Step: 4 }
	order_status_list = append(order_status_list, send_to_3pl)
	// 4 - fail
	send_to_3pl_description_fail := "The warehouse has rejected the order."
	send_to_3pl_fail := OrderStatus{ Description: send_to_3pl_description_fail, Id: SENT_TO_3PL_FAILED, Step: 4 }
	order_status_list = append(order_status_list, send_to_3pl_fail)
	// 5 - success 
	ack_description := "The warehouse has received the order and is processing it."
	ack := OrderStatus{ Description: ack_description, Id: ACK_RECEIVED_FROM_3PL, Step: 5 }
	order_status_list = append(order_status_list, ack)
	// 5 - fail
	ack_description_fail := "The warehouse has not received the order."
	ack_fail := OrderStatus{ Description: ack_description_fail, Id: ACK_NOT_RECEIVED, Step: 5 }
	order_status_list = append(order_status_list, ack_fail)
	// 6 - success
	in_progress_description := "The order is being processed by the warehouse."
	in_progress := OrderStatus{ Description: in_progress_description, Id: IN_PROGRESS, Step: 6 }
	order_status_list = append(order_status_list, in_progress)
	// 6 - fail
	cancelled_description := "The order has been cancelled by the warehouse."
	cancelled := OrderStatus{ Description: cancelled_description, Id: CANCELLED, Step: 6 }
	order_status_list = append(order_status_list, cancelled)
	// 7
	shipped_description := "The order has been shipped."
	shipped := OrderStatus{ Description: shipped_description, Id: SHIPPED, Step: 7 }
	order_status_list = append(order_status_list, shipped)
	// 8 - success
	payment_capture_description := "The payment has been captured."
	payment_capture := OrderStatus{ Description: payment_capture_description, Id: PAYMENT_CAPTURED, Step: 8 }
	order_status_list = append(order_status_list, payment_capture)
	// 8 - fail
	payment_capture_description_fail := "The payment has not been captured."
	payment_capture_fail := OrderStatus{ Description: payment_capture_description_fail, Id: PAYMENT_NOT_CAPTURED, Step: 8 }
	order_status_list = append(order_status_list, payment_capture_fail)
	// 9
	sent_to_poslog_description := "The order has been sent to the POSLOG system."
	sent_to_poslog := OrderStatus{ Description: sent_to_poslog_description, Id: SENT_TO_POSLOG, Step: 9 }
	order_status_list = append(order_status_list, sent_to_poslog)
	// 10
	returned_description := "The order has been returned."			
	returned := OrderStatus{ Description: returned_description, Id: RETURNED, Step: 10 }
	order_status_list = append(order_status_list, returned)

	/*
	str, ok := lo.Find([]string{"foobar"}, func(i string) bool {
    return i == "b"
})
	}
	*/

	// how do I pass in the argument to this function?

	/*
	orderStatus, ok := lo.Find(order_status_list []OrderStatus, func(orderStatus OrderStatus) bool {
		return orderStatus.Id == FRAUD_CHECKED
	})
	*/
	/*
	orderStatus, ok := lo.Find(order_status_list, hello, func(orderStatus OrderStatus) bool {
		return orderStatus.Id == FRAUD_CHECKED
	})*/
	// before I use the lo library, I want to see if I can do this with a for loop
	var orderStatus OrderStatus
	var searchFor = Id
	for _, value := range order_status_list {
		//fmt.Println(i, s)
		if (value.Id == searchFor) {
			orderStatus = value
		}
	}

	return orderStatus.Description 
}

