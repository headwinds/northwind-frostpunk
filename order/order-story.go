package order

import (
	"fmt"
)

// Glossary
// 3PL "Third Party Logistics"

type Status int

const (
	IN_PROGRESS Status = iota
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
    description string
	id Status
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
	order := OrderStatus{ description: "The order is in progress", id: IN_PROGRESS }

	message := fmt.Sprintf("Order status: %s", order.description)
	return message
}

/*

Description

FRAUD_CHECKED	At the time of placing the Order → During payment Authorization with CyberSource 

SUBMITTED/CANCELLED

This is when an order is placed.  This is the first stage of an order. Customer can cancel the order within 1 hour buffer time till the order is in SUBMITTED status and if the customer cancels the order then the status would become CANCELLED and no further action would be taken on it.

READY

If an order is not canceled by the customer after 1 hr of buffer time the status would be changed to ready.  Does not stay in this status.  If it does then that means the order is stuck.  

SEND_TO_3PL / SENT_TO_3PL_FAILED

This status means that the order information is being sent to the warehouse.  For BB the warehouse is called "Think Logistics". If more than 15 mins then the order Is stuck.

If any failure in sending the Order status/ invoice feed the Order is moved to SENT_TO_3PL_FAILED.

ACK_RECEIVED_FROM_3PL / ACK_NOT_RECEIVED

When the warehouse has received the order information then an acknowledgement is received, otherwise that status would be ACK_NOT_RECEIVED.

IN_PROGRESS / CANCELLED

After acknowledgement, status feeds will be obtained and there can be multiple status feeds. If warehouse cancel the order then the status will be marked as CANCELLED. 

SHIPPED

This status means that order is shipped from warehouse. An order number will be provided.   

PAYMENT_CAPTURED / PAYMENT_NOT_CAPTURED

Payment is received successfully if order is in PAYMENT_CAPTURED status otherwise there was a error when obtaining payment. This should be only there for a few seconds.

SENT_TO_POSLOG

Log of the order is sent. 

RETURNED

If customer returns the order.

SENT_TO_POSLOG	

Log of the return request is sent. 

*/